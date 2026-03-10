package users

import (
	"fmt"
	"net/http"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserCotroller(userService *UserServices) *UserControllers {
	return &UserControllers{service: userService}
}

func (uc *UserControllers) GetUsers(c *gin.Context) {
	var users []Users
	query := uc.service.db.Find(&users)
	if query.Error != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status:  "failed",
			Message: "Internal Error.",
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status:  "success",
		Message: fmt.Sprintf("%d users found.", len(users)),
		Data:    users,
	})
}

func (uc *UserControllers) GetUserById(c *gin.Context) {
	id := c.Param("id")
	var user Users
	query := uc.service.db.First(&user, id)
	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, helpers.APIResponse{
				Status:  "failed",
				Message: "Invalid ID.",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status:  "failed",
			Message: "Internal Error.",
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Data:   user,
	})
}
