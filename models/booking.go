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

func GetBookingList(tx *gorm.DB, customerName string, unitID uint, status string, unitInUse *bool, offset int, limit int) ([]Booking, error) {
	var bookings []Booking

	query := tx.Model(&Booking{})

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

	err := query.Offset(offset).Limit(limit).Find(&bookings).Order("time DESC").Error
	if err != nil {
		return []Booking{}, err
	}

	return bookings, nil
}
