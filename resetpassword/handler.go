package resetpassword

import (
	"projects/adiza-exchange-server/say"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {

	req := struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}{}

	err := c.ShouldBind(&req)
	if err != nil {
		say.BadRequest(c, err)
		return
	}

	user, err := service(req.Email, req.Password, req.ConfirmPassword)
	if err != nil {
		say.Unproccessable(c, err)
	}

	say.OK(c, "reset password successfully", user)
}
