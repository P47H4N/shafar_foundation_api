package models

import "time"

type Donation struct {
	ID            uint      `json:"-" gorm:"primaryKey;autoIncrement"`
	UserID        *uint     `json:"user_id" gorm:"index"`
	ProjectID     *uint     `json:"project_id" gorm:"index"`
	DonationType  string    `json:"donation_type" gorm:"type:varchar(20);default:general"`
	Amount        float64   `json:"amount" gorm:"type:decimal(15,2);not null"`
	DonorName     string    `json:"donor_name" gorm:"type:varchar(255)"`
	DonorMobile   string    `json:"donor_mobile" gorm:"type:varchar(20)"`
	DonorEmail    string    `json:"donor_email" gorm:"type:varchar(255)"`
	PaymentMethod string    `json:"payment_method" gorm:"type:varchar(50)"`
	TransactionID string    `json:"transaction_id" gorm:"type:varchar(100)"`
	Status        string    `json:"status" gorm:"type:varchar(20);default:pending"`
	DonatedAt     time.Time `json:"donated_at" gorm:"default:current_timestamp"`
}
