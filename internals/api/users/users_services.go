package users

import (
	"errors"

	"github.com/P47H4N/shafar_foundation_api/internals/models"
	"gorm.io/gorm"
)

func InitUserServices(database *gorm.DB) *UserServices {
	database.AutoMigrate(&models.User{})
	return &UserServices{
		db: database,
	}
}

func (us *UserServices) GetUsers() ([]models.User, error) {
	var users []models.User
	query := us.db.Find(&users).Error
	if query != nil {
		return nil, errors.New("Internal error.")
	}
	if len(users)==0 {
		return nil, errors.New("No user found.")
	}
	return users, nil
}

func (us *UserServices) GetUserById(id int) (*models.User, error) {
	var user models.User
	query := us.db.First(&user, id).Error
	if query != nil {
		if query == gorm.ErrRecordNotFound {
			return nil, errors.New("User id not found.")
		}
		return nil, errors.New("Internal error.")
	}
	return &user, nil
}

