package register

import (
	"encoding/json"
	"errors"
	"fmt"
	"projects/adiza-exchange-server/models"
	"projects/adiza-exchange-server/say"
	"projects/adiza-exchange-server/upload"

	"github.com/gin-gonic/gin"
)

// mulitpart form
// expect user photo
// expect user json
func Handler(c *gin.Context) {

	error404 := "404"
	userJson := c.DefaultPostForm("user", error404)
	if userJson == error404 {
		say.BadRequest(c, errors.New("user json not found"))
	}

	// single file
	_, err := c.FormFile("file")
	if err != nil {
		say.BadRequest(c, err)
		return
	}

	// convert string to json
	var (
		user = models.User{}
		data = []byte(userJson)
	)
	err = json.Unmarshal(data, &user)
	if err != nil {
		say.BadRequest(c, err)
		return
	}

	// save photo
	photo, err := upload.Service(c)
	if err != nil {
		say.Unproccessable(c, err)
		return
	}

	user.Photo = photo

	if err := service(&user); err != nil {
		say.Unproccessable(c, err)
		return
	}

	location := fmt.Sprintf("/users/%d", user.ID)

	say.Created(c, "created user successfully", location, user)
}
