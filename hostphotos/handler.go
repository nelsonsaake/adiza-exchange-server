package hostphotos

import (
	"projects/adiza-exchange-server/env"

	"github.com/gin-gonic/gin"
)

func AddRoute(r gin.IRouter) {
	r.Static(env.OnlinePhotosDIr, env.PhotosDIr)
}
