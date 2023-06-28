package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

func GetAllUnits(tx *gorm.DB, rentalID string) ([]Unit, error) {
	var units []Unit
	result := tx.Where("rental_id = ?", rentalID).Find(&units)
	return units, result.Error
}

func CreateUnit(tx *gorm.DB, name string, hourlyPrice int, category string, rentalID string) (Unit, error) {
	unit := &Unit{
		Name:        name,
		HourlyPrice: hourlyPrice,
		Category:    category,
		RentalID:    rentalID,
	}
	result := tx.Create(unit)
	return *unit, result.Error
}

func GetUnit(tx *gorm.DB, unitID uint, rentalID string) (Unit, error) {
	var unit Unit
	result := tx.Where("id = ? AND rental_id = ?", unitID, rentalID).First(&unit)
	return unit, result.Error
}

type HistoricalRevenueValue struct {
	Time        string  `json:"time"`
	Value       float64 `json:"revenue"`
	Improvement float64 `json:"improvement_pct"`
}

type HistoricalRevenue struct {
	AvgValue float64                  `json:"avg"`
	MaxValue float64                  `json:"max"`
	MinValue float64                  `json:"min"`
	History  []HistoricalRevenueValue `json:"history"`
}

func GetRevenue(tx *gorm.DB, rentalID string, aggregation string, startTime time.Time, endTime time.Time) (HistoricalRevenue, error) {
	var qryResult string
	qry := `
		with raw_profit as (
			select 
				time_bucket(?, finish_time) as time,
				coalesce(sum(tarif)) as value
			from unit_sessions us 
			join units u on u.id = us.unit_id 
			where u.rental_id = ? and finish_time > ? and finish_time <= ?
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
