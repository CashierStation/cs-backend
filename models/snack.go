package models

import (
	"errors"

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
		Stock:    0,
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

func CreateSnackTransaction(tx *gorm.DB, unitSessionID uint, snackID uint, quantity int) (SnackTransaction, error) {
	var snack Snack
	result := tx.Where("id = ?", snackID).First(&snack)

	if result.Error != nil {
		return SnackTransaction{}, result.Error
	}

	if snack.Stock < quantity {
		return SnackTransaction{}, errors.New("not enough stock")
	}

	snack.Stock -= quantity
	tx.Save(&snack)

	totalPrice := snack.Price * quantity

	snackTransaction := &SnackTransaction{
		UnitSessionID: unitSessionID,
		SnackID:       snackID,
		Quantity:      quantity,
		Total:         totalPrice,
	}

	result = tx.Create(snackTransaction)
	return *snackTransaction, result.Error
}

func CreateSnackRestock(tx *gorm.DB, rentalID string, snackID uint, quantity int, totalPrice int) (SnackRestock, error) {
	var snack Snack
	result := tx.Where("id = ?", snackID).First(&snack)

	if result.Error != nil {
		return SnackRestock{}, result.Error
	}

	snack.Stock += quantity
	tx.Save(&snack)

	snackRestock := &SnackRestock{
		SnackID:  snackID,
		Quantity: quantity,
		Total:    totalPrice,
	}

	result = tx.Create(snackRestock)
	return *snackRestock, result.Error
}
