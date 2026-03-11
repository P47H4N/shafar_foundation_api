package users

import (
	"github.com/P47H4N/shafar_foundation_api/internals/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController *UserControllers) {
	userGroup := router.Group("/users")
	userGroup.Use(middleware.UserMiddleware())
	{
		userGroup.GET("/", middleware.AdminMiddleware(), userController.GetUsers)
		userGroup.GET("/:id", userController.GetUserById)
		userGroup.PUT("/:id", userController.UpdateUser)
		userGroup.PUT("/change-password", userController.ChangePassword)
		userGroup.DELETE("/:id", userController.DeleteUser)
	}
}
