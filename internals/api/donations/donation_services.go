package donations

import (
	"errors"

	"github.com/P47H4N/shafar_foundation_api/internals/models"
	"gorm.io/gorm"
)

func InitDonationService(database *gorm.DB) *DonationServices {
	database.AutoMigrate(&models.Donation{})
	return &DonationServices{
		db: database,
	}
}

func (ds *DonationServices) GetDonations() ([]models.Donation, error) {
	var donations []models.Donation
	if err := ds.db.Find(&models.Donation{}).Error; err != nil {
		return nil, errors.New("Unable to find donations.")
	}
	if len(donations) == 0 {
		return nil, errors.New("No donation found.")
	}
	return donations, nil
}

func (ds *DonationServices) CreateDonation(db *DonationBody) error {
	if err := ds.db.Model(&models.Donation{}).Create(&db).Error; err != nil {
		return errors.New("Unable to create donation.")
	}
	return nil
}

func (ds *DonationServices) UpdateDonation(status string, id uint) error {
	if err := ds.db.Model(&models.Donation{}).First(id).Error; err != nil {
		return errors.New("Donation not found")
	}
	statusMap := map[string]string{
		"approve": "approved",
		"reject":  "rejected",
	}
	if err := ds.db.Model(&models.Donation{}).Where("id = ?", id).Update("status", statusMap[status]); err != nil {
		return errors.New("Update Failed.")
	}
	return nil
}

func (ds *DonationServices) DeleteDonation(id uint) error {
	if err := ds.db.Model(&models.Donation{}).First(id).Error; err != nil {
		return errors.New("Donation not found")
	}
	if err := ds.db.Delete(&models.Donation{}, id).Error; err != nil {
		return errors.New("Donation not deleted.")
	}
	return nil
}
