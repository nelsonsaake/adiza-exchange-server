package register

import (
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
	"projects/adiza-exchange-server/users"
)

func service(user *models.User) error {

	if err := users.Valid(*user); err != nil {
		return err
	}

	return repo.DB.Create(user).Error
}
