package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			code := http.StatusInternalServerError

			if c.Writer.Status() == http.StatusOK {
				code = http.StatusInternalServerError
			}

			c.JSON(code, ErrorResponse{
				Status:  "error",
				Message: err.Error(),
				Code:    code,
			})

			c.Abort()
		}
	}
}
