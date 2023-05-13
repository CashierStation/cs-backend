package models

import (
	"gorm.io/gorm"
)

func GetOrCreateRental(tx *gorm.DB, id string, email string) (Rental, error) {
	var rental Rental
	result := tx.FirstOrCreate(&rental, Rental{
		ID:      id,
		Address: "",
		Email:   email,
	})
	return rental, result.Error
}

func GetRentalById(tx *gorm.DB, id string) (Rental, error) {
	var rental Rental
	result := tx.Where("id = ?", id).First(&rental)
	return rental, result.Error
}
