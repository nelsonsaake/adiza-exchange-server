package upload

import (
	"projects/adiza-exchange-server/say"

	"github.com/gin-gonic/gin"
)

// Set a lower memory limit for multipart forms (default is 32 MiB)
// router.MaxMultipartMemory = 8 << 20  // 8 MiB
func Handler(c *gin.Context) {

	// single file
	_, err := c.FormFile("file")
	if err != nil {
		say.BadRequest(c, err)
		return
	}

	// Upload the file to specific dst
	photoURL, err := Service(c)
	if err != nil {
		say.Unproccessable(c, err)
		return
	}

	var (
		msg  = "Uploaded Successfully"
		data = map[string]interface{}{"location": photoURL}
	)

	say.OK(c, msg, data)
}
