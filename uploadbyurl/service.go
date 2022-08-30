package uploadbyurl

import (
	"path/filepath"
	"projects/adiza-exchange-server/env"
	"projects/adiza-exchange-server/helpers"
)

func Service(url string) (filename string, err error) {
	filename = helpers.MakeName(url)
	err = helpers.DownloadImage(url, filepath.Join(env.PhotosDIr, filename))
	return
}
