package donations

import (
	"github.com/P47H4N/shafar_foundation_api/internals/middleware"
	"github.com/gin-gonic/gin"
)

func DonationRoutes(router *gin.Engine, dc *DonationControllers) {
	router.POST("/donations", dc.CreateDonation)
	projectGroup := router.Group("/donations")
	projectGroup.Use(middleware.UserMiddleware(), middleware.AdminMiddleware())
	{
		projectGroup.GET("/", dc.GetDonations)
		projectGroup.PUT("/:status/:id", dc.UpdateDonation)
		projectGroup.DELETE("/:id", dc.DeleteDonation)
	}
}