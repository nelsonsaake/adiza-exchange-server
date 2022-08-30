package helpers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamUint(c *gin.Context, name string) (id uint, err error) {
	idStr := c.Param(name)
	if len(idStr) == 0 {
		err = fmt.Errorf("error, no %s provided", name)
		return
	}
	_id, err := strconv.Atoi(idStr)
	if err != nil {
		err = fmt.Errorf("error parsing %s to int: %v", name, err)
		return
	}
	return uint(_id), err
}
