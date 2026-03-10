package database

import (
	"fmt"

	"github.com/P47H4N/shafar_foundation_api/cmd"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cnf *cmd.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Dhaka",
		cnf.DBHost, cnf.DBUser, cnf.DBPassword, cnf.DBName, cnf.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}