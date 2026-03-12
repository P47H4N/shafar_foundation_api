package projects

import (
	"time"

	"gorm.io/gorm"
)

type ProjectServices struct {
	db *gorm.DB
}

type ProjectControllers struct {
	service *ProjectServices
}

type CreateProjectBody struct {
	ProjectCode  string     `json:"project_code" binding:"required"`
	Title        string     `json:"title" binding:"required"`
	Description  string     `json:"description"`
	Status       string     `json:"status"`
	ProjectDate  *time.Time `json:"project_date"`
	ProjectPlace string     `json:"project_place"`
	ProjectPost  string     `json:"project_post"`
	ImageURL     string     `json:"image_url"`
	DetailsURL   string     `json:"details_url"`
	GoalAmount   float64    `json:"goal_amount"`
	ProjectCost  float64    `json:"project_cost"`
	RaisedAmount float64    `json:"raised_amount"`
}


