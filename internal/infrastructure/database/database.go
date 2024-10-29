package database

import (
	"log"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("postgres.dsn")))
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.Payment{}, &model.User{}, &model.OTP{})

	return db
}
