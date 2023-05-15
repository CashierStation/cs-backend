package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func UpsertSession(tx *gorm.DB, token string, employeeID string) (Session, error) {
	session := Session{
		Token:      token,
		EmployeeID: employeeID,
	}

	query := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "employee_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"token"}),
	}).Create(&session)

	return session, query.Error
}

func GetSessionUser(tx *gorm.DB, token string) (Employee, error) {
	var employee Employee
	tx.Logger.LogMode(logger.Info)

	result := tx.Joins("Role").Joins("Rental").Joins("Session").Where("token = ?", token).First(&employee)

	return employee, result.Error
}

// untested
func GetSessionRole(tx *gorm.DB, token string) (string, error) {
	var role string
	result := tx.Model(&Session{}).Select("roles.name").Joins("Employee.Role").Joins("Employee.Session").Where("token = ?", token).Scan(&role)
	return role, result.Error
}
