package say

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Created(c *gin.Context, message, location string, data interface{}) {

	c.Header("Location", location)

	c.JSON(
		http.StatusCreated,
		Response{
			Message: message,
			Data:    data,
		},
	)
}

func OK(c *gin.Context, message string, data interface{}) {
	c.JSON(
		http.StatusOK,
		Response{
			Message: message,
			Data:    data,
		},
	)
}

func BadRequest(c *gin.Context, err interface{}) {
	c.JSON(
		http.StatusBadRequest,
		Response{
			Message: fmt.Sprint(err),
			Data:    nil,
		},
	)
}

func Unproccessable(c *gin.Context, err interface{}) {
	c.JSON(
		http.StatusUnprocessableEntity,
		Response{
			Message: fmt.Sprint(err),
			Data:    nil,
		},
	)
}

func Unauthorized(c *gin.Context, err interface{}) {
	c.JSON(
		http.StatusUnauthorized,
		Response{
			Message: fmt.Sprint(err),
			Data:    nil,
		},
	)
}

func Forbidden(c *gin.Context, err interface{}) {
	c.JSON(
		http.StatusForbidden,
		Response{
			Message: fmt.Sprint(err),
			Data:    nil,
		},
	)
}
