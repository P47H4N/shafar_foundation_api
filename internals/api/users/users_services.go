package users

import "gorm.io/gorm"

func InitUserServices(database *gorm.DB) *UserServices {
	database.AutoMigrate(&Users{})
	return &UserServices{
		db: database,
	}
}

