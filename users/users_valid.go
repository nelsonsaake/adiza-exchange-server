package users

import (
	"fmt"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/passwords"

	"github.com/go-playground/validator"
)

func Valid(user models.User) (err error) {

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		return err
	}

	err = passwords.ValidatePassword(user.Password)
	if err != nil {
		return fmt.Errorf("password: %s; invalid: %v", user.Password, err)
	}

	return err
}
