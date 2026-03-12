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
	paramId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid user id.",
		})
		return
	}
	getUserId, _ := c.Get("userId")
	getUserRole, _ := c.Get("userRole")
	userId := getUserId.(uint)
	userRole := getUserRole.(string)
	if userId != uint(paramId) && userRole != "admin" {
		c.JSON(http.StatusUnauthorized, helpers.APIResponse{
			Status:  "failed",
			Message: "Unauthorized.",
		})
		return
	}
	user, err := uc.service.GetUserById(uint(userId))
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

func (uc *UserControllers) UpdateUser(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	getUserRole, _ := c.Get("userRole")
	userId := getUserId.(string)
	userRole := getUserRole.(string)
	paramId := c.Param("id")
	uintID, _ := strconv.ParseUint(userId, 10, 32)
	var user UserUpdateBody
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Data.",
		})
		return
	}
	if paramId != userId && userRole != "admin" {
		c.JSON(http.StatusUnauthorized, helpers.APIResponse{
			Status: "failed",
			Message: "Unauthorized.",
		})
		return
	}
	if userRole != user.Role && userRole != "admin" {
		c.JSON(http.StatusUnauthorized, helpers.APIResponse{
			Status: "failed",
			Message: "Role can't be updated",
		})
		user.Role = userRole
	}
	if err := uc.service.UpdateUser(&user, uint(uintID)); err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "User Updated.",
		Data: user,
	})
}

func (uc *UserControllers) DeleteUser(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	getUserRole, _ := c.Get("userRole")
	userId := getUserId.(string)
	userRole := getUserRole.(string)
	paramId := c.Param("id")
	uintID, _ := strconv.ParseUint(userId, 10, 32)
	if paramId != userId && userRole != "admin" {
		c.JSON(http.StatusUnauthorized, helpers.APIResponse{
			Status: "failed",
			Message: "Unauthorized.",
		})
		return
	}
	err := uc.service.DeleteUser(uint(uintID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "failed",
		Message: "User Deleted.",
	})
}

func (uc *UserControllers) ChangePassword(c *gin.Context) {
	getUserId, _ := c.Get("userId")
	userId := getUserId.(string)
	uintID, _ := strconv.ParseUint(userId, 10, 32)
	var passwordBody ChangePasswordBody
	if err := c.ShouldBindBodyWithJSON(&passwordBody); err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Data.",
		})
		return
	}
	if len(passwordBody.NewPassword) < 8 || len(passwordBody.NewPassword) > 32 {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Password must be between 8 and 32 characters.",
		})
		return
	}
	if err := uc.service.ChangePassword(&passwordBody, uint(uintID)); err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "Password Changed Successfully.",
	})
}
