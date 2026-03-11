package auth

import (
	"net/http"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
	"github.com/gin-gonic/gin"
)

func InitAuthController(service *AuthServices) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (ac *AuthController) RegisterUser(c *gin.Context) {
	var registerBody RegisterBody
	err := c.ShouldBindBodyWithJSON(&registerBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Data.",
		})
		return
	}
	if len(registerBody.Password) < 8 || len(registerBody.Password) > 32 {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Password must be between 8 and 32 characters.",
		})
		return
	}
	if len(registerBody.Mobile) < 10 {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Mobile Number.",
		})
		return
	}
	err = ac.service.RegisterUser(&registerBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "User Registered Successfully.",
	})
}

func (ac *AuthController) LoginUser(c *gin.Context) {
	var loginBody LoginBody
	err := c.ShouldBindBodyWithJSON(&loginBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Data",
		})
		return
	}
	if len(loginBody.Mobile) < 10 {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Mobile Number",
		})
		return
	}
	if len(loginBody.Password) < 8 || len(loginBody.Password) > 32 {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: "Invalid Password",
		})
		return
	}
	token, user, err := ac.service.LoginUser(&loginBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.APIResponse{
			Status: "failed",
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, helpers.APIResponse{
		Status: "success",
		Message: "Login successful.",
		Data: gin.H{
			"token": token,
			"user": user,
		},
	})
}
