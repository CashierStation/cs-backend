package models

import "gorm.io/gorm"

func GetEmployeeInRental(tx *gorm.DB, rentalID string, username string) (Employee, error) {
	var employee Employee
	result := tx.Where("username = ? AND rental_id = ?", username, rentalID).First(&employee)
	return employee, result.Error
}

func GetAnyEmployeeByName(tx *gorm.DB, username string) (Employee, error) {
	var employee Employee
	result := tx.Where("username = ?", username).First(&employee)
	return employee, result.Error
}

func RentalHasEmployee(tx *gorm.DB, rentalID string) (bool, error) {
	var employee Employee
	result := tx.Where("rental_id = ?", rentalID).First(&employee)

	if result.Error == nil {
		return true, nil
	}

	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	return false, result.Error
}
