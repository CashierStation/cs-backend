package models

import "gorm.io/gorm"

func GetEmployeeInRental(tx *gorm.DB, rentalID string, username string) (Employee, error) {
	var employee Employee
	result := tx.Where("username = ? AND rental_id = ?", username, rentalID).First(&employee)
	return employee, result.Error
}
