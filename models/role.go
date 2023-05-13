package models

import (
	"gorm.io/gorm"
)

func GetRoleByName(tx *gorm.DB, name string) (Role, error) {
	var role Role
	result := tx.Where("name = ?", name).First(&role)
	return role, result.Error
}
