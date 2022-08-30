package helpers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryID(c *gin.Context, key string) (id int, err error) {
	idStr := c.Query(key)
	if len(idStr) == 0 {
		err = fmt.Errorf("error, no %s provided", key)
		return
	}
	id, err = strconv.Atoi(idStr)
	if err != nil {
		err = fmt.Errorf("error parsing %s to int: %v", key, err)
		return
	}
	return
}
