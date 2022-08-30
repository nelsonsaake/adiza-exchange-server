package smpids

import (
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
)

func GetIDs() (ids []uint, err error) {
	ids = make([]uint, 0)
	err = repo.DB.
		Model(&models.SocialMediaPlatform{}).
		Distinct().Pluck("id", &ids).
		Error
	return
}
