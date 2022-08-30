package getuser

import (
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
)

func service(id uint) (user models.User, err error) {
	err = repo.DB.First(&user, id).Error
	return
}
