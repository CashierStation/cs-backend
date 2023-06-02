package models

import "gorm.io/gorm"

func GetEmployeeInRental(tx *gorm.DB, rentalID string, username string) (Employee, error) {
	var employee Employee
	result := tx.Joins("Rental").Joins("Role").Where("username = ? AND rental_id = ?", username, rentalID).First(&employee)
	return employee, result.Error
}

func GetAnyEmployeeByName(tx *gorm.DB, username string) (Employee, error) {
	var employee Employee
	result := tx.Joins("Rental").Joins("Role").Where("username = ?", username).First(&employee)
	return employee, result.Error
}

func GetAllEmployeeInRental(tx *gorm.DB, rentalID string) ([]Employee, error) {
	var employees []Employee
	result := tx.Joins("Rental").Joins("Role").Where("rental_id = ?", rentalID).Find(&employees)
	return employees, result.Error
}

func RentalHasEmployee(tx *gorm.DB, rentalID string) (bool, error) {
	var employee Employee
	result := tx.Joins("Rental").Joins("Role").Where("rental_id = ?", rentalID).First(&employee)

	if result.Error == nil {
		return true, nil
	}

	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	}

	return false, result.Error
}
