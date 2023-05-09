package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"csbackend/models"

	"os"
)

func New() (*gorm.DB, error) {
	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")
	port := os.Getenv("PGPORT")

	//dsn := os.Getenv("DSN")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)

	database, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
	)

	if err != nil {
		return nil, err
	}

	return database, nil
}

func Migrate(db *gorm.DB) error {
	var models = []interface{}{
		&models.Owner{},
		&models.Rental{},
		&models.Role{},
		&models.Access{},
		&models.Employee{},
		&models.Unit{},
		&models.Booking{},
		&models.Transaction{},
		&models.Snack{},
		&models.SnackRestock{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}
