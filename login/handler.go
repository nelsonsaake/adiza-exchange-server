package login

import (
	"projects/adiza-exchange-server/say"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {

	req := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBind(&req); err != nil {
		say.BadRequest(c, err)
		return
	}

	user, err := service(req.Email, req.Password)
	if err != nil {
		say.Unproccessable(c, err)
		return
	}

	say.OK(c, "logged in successfully", user)
}
