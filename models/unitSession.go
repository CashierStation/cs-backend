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
	return *unitSession, result.Error
}

func GetUnitSessions(tx *gorm.DB, unitID uint, offset uint, limit uint, order string, sortBy string) ([]UnitSession, error) {
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

	result := query.Preload("SnackTransactions").Limit(int(limit)).Find(&unitSessions)
	return unitSessions, result.Error
}

func GetLastUnitSession(tx *gorm.DB, unitID uint) (UnitSession, error) {
	var unitSession UnitSession
	result := tx.Order("start_time desc").Where("unit_id = ?", unitID).First(&unitSession)
	return unitSession, result.Error
}

type unitStatus struct {
	UnitID     uint            `json:"unit_id"`
	Status     enum.UnitStatus `json:"status"`
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
		) select 
			u.id, 
			case when start_time is not null and finish_time is null then 'in_use' else 'idle' end as status, -- TODO: adjust for booking
			coalesce(tarif, 0),
			start_time, 
			finish_time
		from units u
		left join latest_time lt on u.id = lt.unit_id
		where (lt.max is null or lt.max = 1) and u.id in ?
		order by unit_id 
	`, unitIDs, unitIDs).Rows()

	if err != nil {
		return unitSessions, err
	}

	defer rows.Close()

	for rows.Next() {
		var unitSession unitStatus
		err := rows.Scan(&unitSession.UnitID, &unitSession.Status, &unitSession.Tarif, &unitSession.StartTime, &unitSession.FinishTime)
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

	result := tx.Where("id = ?", unitSessionID).First(&unitSession)
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
