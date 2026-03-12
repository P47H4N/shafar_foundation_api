package projects

import (
	"github.com/P47H4N/shafar_foundation_api/internals/middleware"
	"github.com/gin-gonic/gin"
)

func ProjectRoutes(router *gin.Engine, pc *ProjectControllers) {
	router.GET("/projects", pc.GetProjects)
	router.GET("/project/:id", pc.GetProjectsById)
	projectGroup := router.Group("/projects")
	projectGroup.Use(middleware.UserMiddleware(), middleware.AdminMiddleware())
	{
		projectGroup.POST("/", pc.CreateProjects)
		projectGroup.PUT("/:id", pc.UpdateProject)
		projectGroup.DELETE("/:id", pc.DeleteProject)
	}
}