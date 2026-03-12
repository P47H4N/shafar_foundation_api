package donations

import "gorm.io/gorm"

type DonationServices struct {
	db *gorm.DB
}

type DonationControllers struct {
	service *DonationServices
}

type DonationBody struct {
	UserID        *uint     `json:"user_id"`
	ProjectID     *uint     `json:"project_id"`
	DonationType  string    `json:"donation_type"`
	Amount        float64   `json:"amount"`
	DonorName     string    `json:"donor_name"`
	DonorMobile   string    `json:"donor_mobile"`
	DonorEmail    string    `json:"donor_email"`
	PaymentMethod string    `json:"payment_method"`
	TransactionID string    `json:"transaction_id"`
}
