package users

import "gorm.io/gorm"

type UserServices struct {
	db *gorm.DB
}

type UserControllers struct {
	service *UserServices
}

type UserUpdateBody struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	BloodGroup   string `json:"blood_group"`
	Address      string `json:"address"`
	Role         string `json:"-"`
	ProfileImage string `json:"-"`
}

type ChangePasswordBody struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
