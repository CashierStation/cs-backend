package models

import (
	"gorm.io/gorm"
)

func GetAllUnits(tx *gorm.DB, rentalID string) ([]Unit, error) {
	var units []Unit
	result := tx.Where("rental_id = ?", rentalID).Find(&units)
	return units, result.Error
}

func CreateUnit(tx *gorm.DB, name string, hourlyPrice int, rentalID string) (Unit, error) {
	unit := &Unit{
		Name:        name,
		HourlyPrice: hourlyPrice,
		RentalID:    rentalID,
	}
	result := tx.Create(unit)
	return *unit, result.Error
}

func GetUnit(tx *gorm.DB, unitID uint, rentalID string) (Unit, error) {
	var unit Unit
	result := tx.Where("id = ? AND rental_id = ?", unitID, rentalID).First(&unit)
	return unit, result.Error
}
