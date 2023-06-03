package models

import (
	"csbackend/enum"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func CreateBooking(tx *gorm.DB, customerName string, unitID uint, time time.Time) (Booking, error) {
	booking := Booking{
		CustomerName: customerName,
		UnitID:       unitID,
		Time:         time,
		Status:       enum.Waiting.String(),
	}

	err := tx.Create(&booking).Error
	if err != nil {
		return Booking{}, err
	}

	return booking, nil
}

func GetBookingList(tx *gorm.DB, rentalID string, customerName string, unitID uint, status string, unitInUse *bool, offset int, limit int) ([]Booking, error) {
	var bookings []Booking

	query := tx.Model(&Booking{}).Where("rental_id = ?", rentalID)

	if customerName != "" {
		query = query.Where("customer_name LIKE ?", "%"+customerName+"%")
	}

	if unitID != 0 {
		query = query.Where("unit_id = ?", unitID)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if unitInUse != nil {
		ongoingUnitSessions, err := GetOngoingUnitSessions(tx)
		if err != nil {
			return []Booking{}, err
		}

		ongoingUnitSessionUnitIDs := []uint{}
		for _, ongoingUnitSession := range ongoingUnitSessions {
			ongoingUnitSessionUnitIDs = append(ongoingUnitSessionUnitIDs, ongoingUnitSession.UnitID)
		}

		if len(ongoingUnitSessionUnitIDs) == 0 {
			query = query.Where(strconv.FormatBool(!*unitInUse))
		} else if *unitInUse {
			query = query.Where("unit_id IN (?)", ongoingUnitSessionUnitIDs)
		} else if !*unitInUse {
			query = query.Where("unit_id NOT IN (?)", ongoingUnitSessionUnitIDs)
		}
	}

	err := query.Joins("Unit").Offset(offset).Limit(limit).Find(&bookings).Order("time DESC").Error
	if err != nil {
		return []Booking{}, err
	}

	return bookings, nil
}

func GetBookingWithRentalID(tx *gorm.DB, rentalID string, bookingID uint) (Booking, error) {
	var booking Booking
	err := tx.Joins("JOIN units ON bookings.unit_id = units.id").Where("bookings.id = ? AND units.rental_id = ?", bookingID, rentalID).First(&booking).Error
	if err != nil {
		return Booking{}, err
	}

	return booking, nil
}

func GetBooking(tx *gorm.DB, bookingID uint) (Booking, error) {
	var booking Booking
	err := tx.First(&booking, bookingID).Error
	if err != nil {
		return Booking{}, err
	}

	return booking, nil
}

func UpdateBooking(tx *gorm.DB, bookingID uint, unitID uint, customerName string, status string, time *time.Time) (Booking, error) {
	booking, err := GetBooking(tx, bookingID)
	if err != nil {
		return Booking{}, err
	}

	if unitID != 0 {
		booking.UnitID = unitID
	}

	if customerName != "" {
		booking.CustomerName = customerName
	}

	if status != "" {
		booking.Status = status
	}

	if time != nil {
		booking.Time = *time
	}

	err = tx.Save(&booking).Error
	if err != nil {
		return Booking{}, err
	}

	return booking, nil
}
