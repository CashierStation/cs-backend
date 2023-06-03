package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Rental struct {
	Address string
	ID      string `gorm:"primaryKey"`
	Email   string `gorm:"unique"`
}

type Role struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Accesses []Access `gorm:"many2many:role_accesses;"`
}

type Access struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Employee struct {
	ID           string `gorm:"primaryKey"`
	RentalID     string
	RoleID       uint
	Username     string
	PasswordHash string
	Rental       Rental `gorm:"foreignKey:RentalID"`
	Role         Role   `gorm:"foreignKey:RoleID"`
	Session      Session
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Unit struct {
	ID          uint `gorm:"primaryKey"`
	RentalID    string
	Name        string
	Category    string
	HourlyPrice int
	Rental      Rental `gorm:"foreignKey:RentalID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Booking struct {
	ID           uint `gorm:"primaryKey"`
	CustomerName string
	UnitID       uint
	Time         time.Time
	Status       string
	Unit         Unit `gorm:"foreignKey:UnitID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type UnitSession struct {
	ID                uint `gorm:"primaryKey"`
	UnitID            uint
	StartTime         sql.NullTime
	FinishTime        sql.NullTime
	Tarif             int
	Unit              Unit `gorm:"foreignKey:UnitID"`
	SnackTransactions []SnackTransaction
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type SnackTransaction struct {
	ID            uint `gorm:"primaryKey"`
	UnitSessionID uint
	SnackID       uint
	Quantity      int
	Total         int   // keep track of total price of this snack as of this transaction
	Snack         Snack `gorm:"foreignKey:SnackID"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Snack struct {
	ID        uint `gorm:"primaryKey"`
	RentalID  string
	Name      string
	Category  string
	Price     int
	Stock     int
	Rental    Rental `gorm:"foreignKey:RentalID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type SnackRestock struct {
	ID        uint `gorm:"primaryKey"`
	SnackID   uint
	Quantity  int
	Total     int
	Snack     Snack `gorm:"foreignKey:SnackID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Session struct {
	Token      string `gorm:"primaryKey"`
	EmployeeID string `gorm:"unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
