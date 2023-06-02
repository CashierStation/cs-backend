package jobs

import (
	"csbackend/global"
	"csbackend/lib"
	"csbackend/models"
	"time"

	"gorm.io/gorm"
)

type UpdateTarif struct {
	db *gorm.DB
}

func CreateUpdateTarif() UpdateTarif {
	return UpdateTarif{db: global.DB}
}

func (e UpdateTarif) Run() {
	tx := e.db.Begin()

	ongoingSessions, err := models.GetOngoingUnitSessions(tx)
	if err != nil {
		return
	}

	// log.Printf("Updating %d ongoing sessions...\n", len(ongoingSessions))

	if len(ongoingSessions) == 0 {
		return
	}

	for _, sess := range ongoingSessions {
		hourlyPrice := sess.Unit.HourlyPrice
		tarif := lib.CalculateTarif(sess.StartTime.Time, time.Now(), hourlyPrice)
		sess.Tarif = tarif
		tx.Save(&sess)
	}

	tx.Commit()
}
