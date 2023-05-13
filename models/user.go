package models

import (
	"gorm.io/gorm"
)

func CreateEmployee(tx *gorm.DB, id string, username string, passwordHash string, roleID uint, rentalID string) (Employee, error) {
	employee := &Employee{
		ID:           id,
		Username:     username,
		PasswordHash: passwordHash,
		RoleID:       roleID,
		RentalID:     rentalID,
	}
	result := tx.Create(employee)
	return *employee, result.Error
}
