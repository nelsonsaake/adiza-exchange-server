package repo

import (
	"projects/adiza-exchange-server/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newSqliteDB(url string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(url), GormConfig)
	if err != nil {
		return
	}

	err = db.AutoMigrate(models.Migration...)
	if err != nil {
		return
	}

	return
}
