package models

import "gorm.io/gorm"

func GetEmployeeByUsername(tx *gorm.DB, username string) (Employee, error) {
	var employee Employee
	result := tx.Where("username = ?", username).First(&employee)
	return employee, result.Error
}
