package models

import (
	"csbackend/enum"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

func CreateUnitSession(tx *gorm.DB, unitID uint) (UnitSession, error) {
	unitSession := &UnitSession{
		UnitID:            unitID,
		StartTime:         sql.NullTime{Time: time.Now(), Valid: true},
		FinishTime:        sql.NullTime{},
		SnackTransactions: []SnackTransaction{},
		Tarif:             0,
	}
	result := tx.Create(unitSession)

	// get unit session with unit
	tx.Joins("Unit").First(&unitSession, unitSession.ID)

	return *unitSession, result.Error
}

func GetUnitSessions(tx *gorm.DB, rentalID string, unitID uint, offset uint, limit uint, order string, sortBy string, latest bool) ([]UnitSession, error) {
	var unitSessions []UnitSession
	query := tx

	if unitID != 0 {
		query = tx.Where("unit_id = ?", unitID)
	}

	if order == "" {
		order = "desc"
	}

	if sortBy == "" {
		sortBy = "start_time"
	}

	query = query.Order(sortBy + " " + order)

	if latest {
		subQuery := tx.Model(&UnitSession{}).Select("unit_id, max(start_time) as start_time").Group("unit_id")
		query = query.Where("(unit_id, start_time) in (?)", subQuery)
	}

	query = query.Joins("left join units on unit_sessions.unit_id = units.id").Where("units.rental_id = ?", rentalID)

	result := query.Preload("SnackTransactions.Snack").Limit(int(limit)).Find(&unitSessions)
	return unitSessions, result.Error
}

func GetLastUnitSession(tx *gorm.DB, unitID uint) (UnitSession, error) {
	var unitSession UnitSession
	result := tx.Order("start_time desc").Joins("Unit").Where("unit_id = ?", unitID).First(&unitSession)
	return unitSession, result.Error
}

type unitStatus struct {
	UnitID     uint            `json:"unit_id"`
	Status     enum.UnitStatus `json:"status"`
	Booked     bool            `json:"booked"`
	StartTime  *time.Time      `json:"latest_start_time"`
	FinishTime *time.Time      `json:"latest_finish_time"`
	Tarif      int             `json:"tarif"`
}

func GetLastUnitStatuses(tx *gorm.DB, unitIDs []uint) ([]unitStatus, error) {
	var unitSessions []unitStatus = []unitStatus{}

	rows, err := tx.Raw(`
		with latest_time as (
			select 
				rank() over (partition by unit_id order by start_time desc) as max,
				us.*
			from unit_sessions us
			where unit_id in ?
		), booking_status as (
			select
				b.unit_id,
				count(b.status) > 0 as booked 
			from bookings b 
			where status = 'waiting' and unit_id in ?
			group by 1
		) select 
			u.id, 
			case when start_time is not null and finish_time is null then 'in_use' else 'idle' end as status,
			coalesce(bs.booked, false) as booked,
			coalesce(tarif, 0) as tarif,
			start_time, 
			finish_time
		from units u
		left join latest_time lt on u.id = lt.unit_id
		left join booking_status bs on u.id = bs.unit_id
		where (lt.max is null or lt.max = 1) and u.id in ?
		order by u.id 
	`, unitIDs, unitIDs, unitIDs).Rows()

	if err != nil {
		return unitSessions, err
	}

	defer rows.Close()

	for rows.Next() {
		var unitSession unitStatus
		err := rows.Scan(&unitSession.UnitID, &unitSession.Status, &unitSession.Booked, &unitSession.Tarif, &unitSession.StartTime, &unitSession.FinishTime)
		if err != nil {
			return unitSessions, err
		}

		if unitSession.FinishTime != nil && unitSession.FinishTime.IsZero() {
			unitSession.FinishTime = nil
		}

		unitSessions = append(unitSessions, unitSession)
	}

	return unitSessions, err
}

func StopUnitSession(tx *gorm.DB, unitSessionID uint) (UnitSession, error) {
	var unitSession UnitSession

	result := tx.Joins("Unit").Where("unit_sessions.id = ?", unitSessionID).First(&unitSession)
	if result.Error != nil {
		return unitSession, result.Error
	}

	unitSession.FinishTime = sql.NullTime{Time: time.Now(), Valid: true}
	tx.Save(&unitSession)
	return unitSession, result.Error
}

func GetOngoingUnitSessions(tx *gorm.DB) ([]UnitSession, error) {
	var unitSessions []UnitSession
	result := tx.Joins("Unit").Where("finish_time is null").Find(&unitSessions)
	return unitSessions, result.Error
}
