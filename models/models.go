package models

import "gorm.io/gorm"

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
	gorm.Model
	RentalID     string
	RoleID       uint
	Username     string
	PasswordHash string
	Rental       Rental `gorm:"foreignKey:RentalID"`
	Role         Role   `gorm:"foreignKey:RoleID"`
}

type Unit struct {
	gorm.Model
	RentalID    string
	Name        string
	HourlyPrice int
	Rental      Rental `gorm:"foreignKey:RentalID"`
}

type Booking struct {
	gorm.Model
	UnitID uint
	Time   string
	Status int
	Unit   Unit `gorm:"foreignKey:UnitID"`
}

type Transaction struct {
	gorm.Model
	UnitID     uint
	StartTime  string
	FinishTime string
	Unit       Unit    `gorm:"foreignKey:UnitID"`
	Snacks     []Snack `gorm:"many2many:snack_transactions;"`
}

type Snack struct {
	gorm.Model
	RentalID string
	Name     string
	Price    int
	Rental   Rental `gorm:"foreignKey:RentalID"`
}

type SnackRestock struct {
	gorm.Model
	SnackID  uint
	Quantity int
	Total    int
	Snack    Snack `gorm:"foreignKey:SnackID"`
}
