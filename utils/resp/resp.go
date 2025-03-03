package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(
		http.StatusOK,
		Response{
			Code:    200,
			Message: "success",
			Data:    data,
		},
	)
}

func Fail(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(
		http.StatusOK,
		Response{
			Code:    code,
			Message: message,
			Data:    nil,
		},
	)
}
