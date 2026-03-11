package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
	"github.com/gin-gonic/gin"
)

func InitUserCotroller(userService *UserServices) *UserControllers {
	return &UserControllers{service: userService}
}

func (uc *UserControllers) GetUsers(c *gin.Context) {
	users, err := uc.service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status:  "failed",
			Message: err.Error(),
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
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status:  "failed",
			Message: "Invalid User ID.",
		})
		return
	}
	user, err := uc.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Data:   user,
	})
}

