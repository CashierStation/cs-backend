package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"csbackend/models"
	"csbackend/util"

	"os"
)

func New() (*gorm.DB, error) {
	//host := os.Getenv("PGHOST")
	//user := os.Getenv("PGUSER")
	//password := os.Getenv("PGPASSWORD")
	//dbname := os.Getenv("PGDATABASE")
	//port := os.Getenv("PGPORT")

	dsn := os.Getenv("DSN_DEV")
	if util.IsProduction() {
		dsn = os.Getenv("DSN_PROD")
	}
	//dsn := os.Getenv("DSN")
	println(dsn)
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbname, port)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel:                  logger.Error,           // Log level
			SlowThreshold:             time.Millisecond * 500, // Slow SQL threshold
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
		},
	)
	println("created logger")

	database, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			Logger: newLogger,
		},
	)

	if err != nil {
		println(err.Error())
		return nil, err
	}

	sqlDB, err := database.DB()

	if err != nil {
		println(err.Error())
		return nil, err
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(80)
	sqlDB.SetConnMaxLifetime(1000 * time.Second)
	sqlDB.SetConnMaxIdleTime(10 * time.Second)

	println("opened database")

	if err != nil {
		println(err.Error())
		return nil, err
	}

	return database, nil
}

func Migrate(db *gorm.DB) error {
	var models = []interface{}{
		&models.Rental{},
		&models.Role{},
		&models.Access{},
		&models.Employee{},
		&models.Unit{},
		&models.Booking{},
		&models.UnitSession{},
		&models.SnackTransaction{},
		&models.Snack{},
		&models.SnackRestock{},
		&models.Session{},
	}

	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	err := postMigration(db)
	if err != nil {
		return err
	}

	return nil
}
