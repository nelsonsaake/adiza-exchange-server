package server

import "github.com/gin-gonic/gin"

func New() *gin.Engine {

	engine := gin.Default()

	engine.Use(cors)

	routing(engine)

	return engine
}
