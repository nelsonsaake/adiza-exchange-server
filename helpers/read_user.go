package helpers

import (
	"projects/adiza-exchange-server/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ReadUser(c *gin.Context) (models.User, error) {

	user := models.User{}
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		return models.User{}, err
	}

	// this is neccessary, because by default, user json don't write out or read password
	// to reduce the chance of it being mistakenly displayed
	emailAndPassword := struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gte=8"`
	}{}

	err := c.ShouldBindBodyWith(&emailAndPassword, binding.JSON)
	if err != nil {
		return user, err
	}

	user.Email = emailAndPassword.Email
	user.Password = emailAndPassword.Password

	return user, nil
}
