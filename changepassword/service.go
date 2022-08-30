package changepassword

import (
	"errors"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/repo"
	"strings"
)

func service(email, password, newPassword, confirmPassword string) (*models.User, error) {

	email = strings.TrimSpace(email)
	if len(email) == 0 {
		return nil, errors.New("invlaid email")
	}

	if newPassword != confirmPassword {
		return nil, errors.New("password doesn't match confirm password")
	}

	user := &models.User{Email: email}

	err := repo.DB.First(&user).Error
	if err != nil {
		return nil, err
	}

	if user.Password == password {
		return nil, errors.New("invalid email or password")
	}

	user.Password = newPassword

	err = repo.DB.Save(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
