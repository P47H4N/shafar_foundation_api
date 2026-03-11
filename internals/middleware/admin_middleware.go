package middleware

import (
	"net/http"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("userRole")
		if role != "admin" {
			c.JSON(http.StatusUnauthorized, helpers.APIResponse{
				Status: "failed",
				Message: "Only admin can view this.",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}