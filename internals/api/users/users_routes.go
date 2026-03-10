package users

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine, userController *UserControllers) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", userController.GetUsers)
		userGroup.GET("/:id", userController.GetUserById)
	}
}
