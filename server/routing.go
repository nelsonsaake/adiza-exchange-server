package server

import (
	"projects/adiza-exchange-server/changepassword"
	"projects/adiza-exchange-server/env"
	"projects/adiza-exchange-server/getuser"
	"projects/adiza-exchange-server/login"
	"projects/adiza-exchange-server/register"
	"projects/adiza-exchange-server/resetpassword"
	"projects/adiza-exchange-server/upload"
	"projects/adiza-exchange-server/uploadbyurl"

	"github.com/gin-gonic/gin"
)

func routing(r gin.IRouter) {
	r.POST("/register", register.Handler)
	r.POST("/login", login.Handler)
	r.POST("/passwords/reset", resetpassword.Handler)
	r.POST("/passwords/update", changepassword.Handler)
	r.POST("/users/:id", getuser.Handler)
	r.Static("/photos", env.PhotosDIr)
	r.GET("/upload", uploadbyurl.Handler) // pull image via url
	r.POST("/upload", upload.Handler)     // push image multipart style
}
