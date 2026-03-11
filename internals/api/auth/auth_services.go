package auth

import (
	"errors"
	"time"

	"github.com/P47H4N/shafar_foundation_api/internals/helpers"
	"github.com/P47H4N/shafar_foundation_api/internals/models"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func InitAuthServices(database *gorm.DB) *AuthServices {
	database.AutoMigrate(&models.User{})
	return &AuthServices{
		db: database,
	}
}

func (as *AuthServices) RegisterUser(rb *RegisterBody) error {
	hashPassword, err := helpers.HashPassword(rb.Password)
	if err != nil {
		return errors.New("Password can't be secured. Try another.")
	}
	var count int64
	as.db.Model(&models.User{}).Where("mobile = ?", rb.Mobile).Count(&count)
	if count > 0 {
		return errors.New("Mobile Number Already Exists.")
	}
	newUser := &models.User{
		Name:         rb.Name,
		Email:        rb.Email,
		Mobile:       rb.Mobile,
		Password:     hashPassword,
		BloodGroup:   rb.BloodGroup,
		Address:      rb.Address,
		Role:         rb.Role,
		ProfileImage: rb.ProfileImage,
	}
	err = as.db.Create(&newUser).Error
	if err != nil {
		return errors.New("User Creation Failed")
	}
	return nil
}

func (as *AuthServices) LoginUser(lb *LoginBody) (string, *models.User, error) {
	var user models.User
	err := as.db.Where("mobile = ?", lb.Mobile).First(&user).Error
	if err != nil {
		return "", nil, errors.New("User not found with this mobile number.")
	}
	if !helpers.CheckPasswordHash(lb.Password, user.Password) {
		return "", nil, errors.New("Invalid credentials.")
	}
	token, err := helpers.GenerateToken(&models.TokenBody{
		Id:     user.Id,
		Mobile: user.Mobile,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	if err != nil {
		return "", nil, errors.New("Token Generate Failed")
	}
	return token, &user, nil
}
