package repository

import (
	"time"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/pkg"
	"gorm.io/gorm"
)

type OTPRepository struct {
	db *gorm.DB
}

func NewOTPRepository(db *gorm.DB) *OTPRepository {
	return &OTPRepository{
		db: db,
	}
}

func (r *OTPRepository) Create(username string) (*model.OTP, error) {
	code := pkg.RandomString(4)

	otp := &model.OTP{
		Username:  username,
		Code:      code,
		ExpiredAt: time.Now().Add(10 * time.Minute),
	}

	err := r.db.Save(&otp).Error
	if err != nil {
		return nil, err
	}

	return otp, nil
}

func (r *OTPRepository) Find(username string) (*model.OTP, error) {
	var otp *model.OTP
	err := r.db.Take(&otp, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return otp, nil
}
