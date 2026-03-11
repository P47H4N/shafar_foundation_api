package auth

import (
	"gorm.io/gorm"
)

type AuthServices struct {
	db *gorm.DB
}

type AuthController struct {
	service *AuthServices
}

type RegisterBody struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile" binding:"required"`
	Password     string `json:"password" binding:"required"`
	BloodGroup   string `json:"blood_group"`
	Address      string `json:"address"`
	Role         string `json:"role"`
	ProfileImage string `json:"profile_image"`
}

type LoginBody struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

