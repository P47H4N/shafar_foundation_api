package cmd

import (
	"github.com/P47H4N/shafar_foundation_api/internals/api/auth"
	"github.com/P47H4N/shafar_foundation_api/internals/api/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Start(router *gin.Engine, db *gorm.DB) error {
	// Users
	userService := users.InitUserServices(db)
	userController := users.InitUserCotroller(userService)
	users.UserRoutes(router, userController)

	// Auth
	authService := auth.InitAuthServices(db)
	authController := auth.InitAuthController(authService)
	auth.AuthRoutes(router, authController)

	return nil
}