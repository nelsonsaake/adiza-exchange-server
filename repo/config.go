package repo

import (
	"gorm.io/gorm"
)

var (
	GormConfig = &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
		// Logger: Logger(),
	}
)
