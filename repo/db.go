package repo

import (
	"projects/adiza-exchange-server/env"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	DB, err = newSqliteDB(env.SQLiteDBURL)
	if err != nil {
		panic(err)
	}
}
