package users

import "gorm.io/gorm"

type UserServices struct {
	db *gorm.DB
}

type UserControllers struct {
	service *UserServices
}
