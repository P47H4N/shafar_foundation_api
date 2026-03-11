package auth

import "github.com/gin-gonic/gin"

func AuthRoutes(router *gin.Engine, authController *AuthController) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.RegisterUser)
		authGroup.POST("/login", authController.LoginUser)
		// authGroup.POST("/forgot")
		// authGroup.POST("/verify-otp")
	}
}
