package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"os"
)

var DB *gorm.DB

func Connect() {
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
		panic("Failed to connect database")
	}

	DB = database
}

func Migrate() error {
	var models = []interface{}{
		&Owner{},
		&Rental{},
		&Role{},
		&Access{},
		&Employee{},
		&Unit{},
		&Booking{},
		&Transaction{},
		&Snack{},
		&SnackRestock{},
	}

	for _, model := range models {
		err := DB.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}
