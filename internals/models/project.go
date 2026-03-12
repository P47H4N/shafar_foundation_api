package models

import (
	"time"
)

type Project struct {
	ID           uint       `json:"-" gorm:"primaryKey;autoIncrement"`
	ProjectCode  string     `json:"project_code" gorm:"type:varchar(50)"`
	Title        string     `json:"title" gorm:"type:varchar(255)"`
	Description  string     `json:"description" gorm:"type:text"`
	Status       string     `json:"status" gorm:"type:varchar(20);default:upcoming;check:status IN ('running','upcoming','past')"`
	ProjectDate  *time.Time `json:"project_date" gorm:"type:date"`
	ProjectPlace string     `json:"project_place" gorm:"type:varchar(255)"`
	ProjectPost  string     `json:"project_post" gorm:"type:varchar(255)"`
	ImageURL     string     `json:"image_url" gorm:"type:varchar(255)"`
	DetailsURL   string     `json:"details_url" gorm:"type:varchar(255)"`
	GoalAmount   float64    `json:"goal_amount" gorm:"type:numeric(15,2)"`
	ProjectCost  float64    `json:"project_cost" gorm:"type:numeric(15,2);default:0.00"`
	RaisedAmount float64    `json:"raised_amount" gorm:"type:numeric(15,2);default:0.00"`
	CreatedAt    time.Time  `json:"-" gorm:"autoCreateTime"`
}
