package uploadbyurl

import (
	"projects/adiza-exchange-server/say"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	url := c.Query("url")

	filename, err := Service(url)
	if err != nil {
		say.Unproccessable(c, err)
		return
	}

	var (
		msg  = "Uploaded Successfully"
		data = map[string]interface{}{"filename": filename}
	)

	say.OK(c, msg, data)
}
