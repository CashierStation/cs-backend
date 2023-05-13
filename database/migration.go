package db

import (
	"csbackend/models"

	"gorm.io/gorm"
)

func postMigration(db *gorm.DB) error {
	// Insert data to table roles
	ownerRole := models.Role{Name: "owner"}
	db.FirstOrCreate(&ownerRole, ownerRole)

	employeeRole := models.Role{Name: "karyawan"}
	db.FirstOrCreate(&employeeRole, employeeRole)

	return nil
}
