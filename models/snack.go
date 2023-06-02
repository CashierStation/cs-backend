package models

import (
	"gorm.io/gorm"
)

func GetAllSnacks(tx *gorm.DB, rentalID string) ([]Snack, error) {
	var snacks []Snack
	result := tx.Where("rental_id = ?", rentalID).Find(&snacks)
	return snacks, result.Error
}

func CreateSnack(tx *gorm.DB, rentalID string, name string, category string, price int) (Snack, error) {
	snack := &Snack{
		Name:     name,
		RentalID: rentalID,
		Category: category,
		Price:    price,
	}
	result := tx.Create(snack)
	return *snack, result.Error
}

func GetSnack(tx *gorm.DB, rentalID string, snackID uint) (Snack, error) {
	var snack Snack
	result := tx.Where("id = ? AND rental_id = ?", snackID, rentalID).First(&snack)
	return snack, result.Error
}
