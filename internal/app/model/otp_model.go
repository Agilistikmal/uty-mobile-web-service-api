package model

import "time"

type OTP struct {
	Username  string    `json:"username,omitempty" gorm:"primaryKey"`
	Code      string    `json:"code,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
