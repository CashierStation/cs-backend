package models

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

func GetAllSnacks(tx *gorm.DB, rentalID string) ([]Snack, error) {
	var snacks []Snack
	result := tx.Where("rental_id = ?", rentalID).Find(&snacks)
	return snacks, result.Error
}

func CreateSnack(tx *gorm.DB, rentalID string, name string, category string, price int, stock int) (Snack, error) {
	snack := &Snack{
		Name:     name,
		RentalID: rentalID,
		Category: category,
		Stock:    stock,
		Price:    price,
	}
	result := tx.Create(snack)
	return *snack, result.Error
}

func GetSnack(tx *gorm.DB, rentalID string, snackID uint) (Snack, error) {
	var snack Snack
	result := tx.Where("id = ? AND rental_id = ?", snackID, rentalID).First(&snack)
	return snack, result.Error
}

func CreateSnackTransaction(tx *gorm.DB, unitSessionID uint, snackID uint, quantity int) (SnackTransaction, error) {
	var snack Snack
	result := tx.Where("id = ?", snackID).First(&snack)

	if result.Error != nil {
		return SnackTransaction{}, result.Error
	}

	if snack.Stock < quantity {
		return SnackTransaction{}, errors.New("not enough stock")
	}

	snack.Stock -= quantity
	tx.Save(&snack)

	totalPrice := snack.Price * quantity

	snackTransaction := &SnackTransaction{
		UnitSessionID: unitSessionID,
		SnackID:       snackID,
		Quantity:      quantity,
		Total:         totalPrice,
	}

	result = tx.Create(snackTransaction)
	return *snackTransaction, result.Error
}

func CreateSnackRestock(tx *gorm.DB, rentalID string, snackID uint, quantity int, totalPrice int) (SnackRestock, error) {
	var snack Snack
	result := tx.Where("id = ?", snackID).First(&snack)

	if result.Error != nil {
		return SnackRestock{}, result.Error
	}

	snack.Stock += quantity
	tx.Save(&snack)

	snackRestock := &SnackRestock{
		SnackID:  snackID,
		Quantity: quantity,
		Total:    totalPrice,
	}

	result = tx.Create(snackRestock)
	return *snackRestock, result.Error
}

func GetSnackRevenue(tx *gorm.DB, rentalID string, aggregation string, startTime time.Time, endTime time.Time) (HistoricalRevenue, error) {
	var qryResult string
	qry := `
		with raw_profit as (
			select 
				time_bucket(?, st.created_at) as time,
				coalesce(sum(st.total)) as value
			from snack_transactions st 
			join snacks s on s.id = st.snack_id 
			where s.rental_id = ? and st.created_at > ? and st.created_at <= ? 
			group by 1 
			order by 1
		), avg_profit as (
			select
				case 
					when avg(value) != 0 then avg(value)
					else null
				end as avg_val
			from raw_profit		
		), summary as (
			select
				avg(value),
				max(value),
				min(value),
				json_agg(
					json_build_object(
						'time', time,
						'revenue', value,
						'improvement_pct', (value - avg_val) / avg_val * 100
					) 
				) as history
			from raw_profit, avg_profit
		) select to_json(summary) from summary
	`
	aggregation = "1 " + aggregation
	result := tx.Raw(qry, aggregation, rentalID, startTime, endTime).Scan(&qryResult)
	if result.Error != nil {
		return HistoricalRevenue{}, result.Error
	}

	var revenue HistoricalRevenue
	err := json.Unmarshal([]byte(qryResult), &revenue)
	if err != nil {
		return HistoricalRevenue{}, err
	}

	return revenue, nil
}
