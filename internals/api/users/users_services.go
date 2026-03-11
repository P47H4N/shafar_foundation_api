package users

import (
	"errors"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
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

func (us *UserServices) GetUserById(id uint) (*models.User, error) {
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

func (us *UserServices) UpdateUser(user *UserUpdateBody, id uint) error {
	if err := us.db.First(&user, id).Error; err != nil {
		return errors.New("User not found.")
	}
	if err := us.db.Model(&models.User{}).Where("id = ?", id).Updates(user); err != nil {
		return errors.New("Update Failed.")
	}
	return nil
}

func (us *UserServices) DeleteUser(id uint) error {
	if err := us.db.Model(&models.User{}).Where("id = ?", id).Error; err != nil {
		return errors.New("User not found.")
	}
	if err := us.db.Delete(&models.User{}, id); err != nil {
		return errors.New("Unable to delete user.")
	}
	return nil
}

func (us *UserServices) ChangePassword(pb *ChangePasswordBody, id uint) error {
	user := &models.User{}
	if err := us.db.Model(&user).Where("id = ?", id).Error; err != nil {
		return errors.New("User not found.")
	}
	if !helpers.CheckPasswordHash(pb.OldPassword, user.Password) {
		return errors.New("Wrong Current Password.")
	}
	newHashedPassword, _ := helpers.HashPassword(pb.NewPassword)
	if err := us.db.Model(&user).Update("password", newHashedPassword).Error; err != nil {
        return errors.New("Failed to update password.")
    }
	return nil
}
