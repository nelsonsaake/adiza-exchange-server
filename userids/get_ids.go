package userids

import (
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
)

func GetIDs() (ids []uint, err error) {
	ids = make([]uint, 0)
	err = repo.DB.
		Model(&models.User{}).
		Distinct().Pluck("id", &ids).
		Error

	return
}
