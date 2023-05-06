package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"os"
)

var DB *gorm.DB

func Connect() {
	//host := os.Getenv("PGHOST")
	//user := os.Getenv("PGUSER")
	//password := os.Getenv("PGPASSWORD")
	//dbname := os.Getenv("PGDATABASE")
	//port := os.Getenv("PGPORT")

	dsn := os.Getenv("DSN")

	database, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		},
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
