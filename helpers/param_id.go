package helpers

import (
	"github.com/gin-gonic/gin"
)

func ParamID(c *gin.Context) (id uint, err error) {
	return ParamUint(c, "id")
}
