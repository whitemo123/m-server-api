package middlewares

import (
	"fmt"
	"m-server-api/utils/resp"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			stack := debug.Stack()
			fmt.Printf("panic: %v\n%s\n", err, stack)
			resp.Fail(c, http.StatusInternalServerError, fmt.Sprintf("%+v", err))
			c.Abort()
		}
	}()
	c.Next()
}
