package upload

import (
	"path"
	"path/filepath"
	"projects/adiza-exchange-server/env"
	"projects/adiza-exchange-server/helpers"

	"github.com/gin-gonic/gin"
)

func Service(c *gin.Context) (string, error) {

	// single file
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}

	// Upload the file to specific dst
	fpath := filepath.Join(env.PhotosDIr, helpers.MakeName(file.Filename))
	err = c.SaveUploadedFile(file, fpath)
	if err != nil {
		return "", err
	}

	photoURL := path.Join(env.OnlinePhotosDIr, file.Filename)

	return photoURL, nil
}
