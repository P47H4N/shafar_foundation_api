package users

import (
	"time"
	"gorm.io/gorm"
)

type Users struct {
	Id           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"type:varchar(255);not null"`
	Email        string    `json:"email" gorm:"type:varchar(255)"`
	Mobile       string    `json:"mobile" gorm:"type:varchar(20);unique;not null"`
	Password     string    `json:"-" gorm:"type:varchar(255);not null"`
	BloodGroup   string    `json:"blood_group" gorm:"type:varchar(10)"`
	Address      string    `json:"address" gorm:"type:text"`
	Role         string    `json:"role" gorm:"type:varchar(20);default:'user'"`
	ProfileImage string    `json:"profile_image" gorm:"type:varchar(255)"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type UserServices struct {
	db *gorm.DB
}

type UserControllers struct {
	service *UserServices
}