package login

import (
	"errors"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
	"strings"
)

func service(email, password string) (*models.User, error) {

	email = strings.TrimSpace(email)
	if len(email) == 0 {
		return nil, errors.New("invlaid email or password") // generalised
	}

	user := &models.User{}

	err := repo.DB.First(user).Error
	if err != nil {
		return nil, err
	}

	if user.Email != email || user.Password != password {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
