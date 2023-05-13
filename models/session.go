package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	result := tx.Joins("Session").Where("token = ?", token).First(&employee)
	return employee, result.Error
}

func GetSessionRole(tx *gorm.DB, token string) (string, error) {
	var role string
	result := tx.Joins("Session").Where("token = ?", token).Select("role").First(&role)
	return role, result.Error
}
