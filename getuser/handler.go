package getuser

import (
	"projects/adiza-exchange-server/helpers"
	"projects/adiza-exchange-server/say"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {

	id, err := helpers.ParamID(c)
	if err != nil {
		say.BadRequest(c, err)
		return
	}

	user, err := service(id)
	if err != nil {
		say.Unproccessable(c, err)
		return
	}

	say.OK(c, "got user successfully", user)
}
