package projects

import (
	"errors"

	"github.com/P47H4N/shafar_foundation_api/internals/models"
	"gorm.io/gorm"
)

func InitProjectServices(database *gorm.DB) *ProjectServices {
	database.AutoMigrate(&models.Project{})
	return &ProjectServices{
		db: database,
	}
}

func (ps *ProjectServices) GetProjects() ([]models.Project, error) {
	var projects []models.Project
	if err := ps.db.Find(&projects).Error; err != nil {
		return nil, errors.New("Unable to find projects.")
	}
	if len(projects) == 0 {
		return nil, errors.New("No projects found.")
	}
	return projects, nil
}

func (ps *ProjectServices) GetProjectsById(id uint) (*models.Project, error) {
	var project models.Project
	if err := ps.db.Where("id = ?", id).First(&project).Error; err != nil {
		return nil, errors.New("No project found.")
	}
	return &project, nil
}

func (ps *ProjectServices) CreateProjects(pb *CreateProjectBody) error {
	if err := ps.db.Model(&models.Project{}).Create(&pb).Error; err != nil {
		return errors.New("Project not created")
	}
	return nil
}

func (ps *ProjectServices) UpdateProject(pb *CreateProjectBody, id uint) error {
	if err := ps.db.Model(&models.Project{}).First(id).Error; err != nil {
		return errors.New("Project not found")
	}
	if err := ps.db.Model(&models.Project{}).Where("id = ?", id).Updates(pb); err != nil {
		return errors.New("Update Failed.")
	}
	return nil
}

func (ps *ProjectServices) DeleteProject(id uint) error {
	if err := ps.db.Model(&models.Project{}).First(id).Error; err != nil {
		return errors.New("Project not found")
	}
	if err := ps.db.Delete(&models.Project{}, id).Error; err != nil {
		return errors.New("Project not deleted.")
	}
	return nil
}
