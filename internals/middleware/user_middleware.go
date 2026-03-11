package middleware

import (
	"net/http"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := helpers.GetToken(c)
		if err != nil {
			c.JSON(http.StatusNetworkAuthenticationRequired, helpers.APIResponse{
				Status: "failed",
				Message: "Authorization required.",
			})
			c.Abort()
			return
		}
		claims, err := helpers.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, helpers.APIResponse{
				Status: "failed",
				Message: "Unauthorized.",
			})
			c.Abort()
			return
		}
		c.Set("userId", claims.Id)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}