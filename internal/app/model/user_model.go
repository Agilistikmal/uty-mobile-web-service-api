package model

import "time"

type User struct {
	Username      string    `json:"username,omitempty" gorm:"primaryKey" validate:"required,min=3,max=20"`
	FullName      string    `json:"full_name,omitempty" validate:"required,min=3,max=50"`
	Phone         string    `json:"phone,omitempty" validate:"required,e164"`
	Password      string    `json:"password,omitempty" validate:"required,min=8"`
	PasswordRetry *int      `json:"password_retry,omitempty" gorm:"default:0"`
	Verified      bool      `json:"verified,omitempty" gorm:"default:false"`
	LockedAt      time.Time `json:"locked_at,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

type UserResponse struct {
	Username      string    `json:"username,omitempty"`
	FullName      string    `json:"full_name,omitempty"`
	Phone         string    `json:"phone,omitempty"`
	PasswordRetry *int      `json:"password_retry,omitempty"`
	Verified      bool      `json:"verified,omitempty"`
	LockedAt      time.Time `json:"locked_at,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type UserCreateRequest struct {
	Username string `json:"username,omitempty" gorm:"primaryKey" validate:"required,min=3,max=20"`
	FullName string `json:"full_name,omitempty" validate:"required,min=3,max=50"`
	Phone    string `json:"phone,omitempty" validate:"required,e164"`
	Password string `json:"password,omitempty" validate:"required,min=8"`
}

type UserUpdateRequest struct {
	FullName      string    `json:"full_name,omitempty" validate:"omitempty,min=3,max=50"`
	Phone         string    `json:"phone,omitempty" validate:"omitempty,e164"`
	Password      string    `json:"password,omitempty" validate:"omitempty,min=8"`
	PasswordRetry *int      `json:"password_retry,omitempty" validate:"omitempty,number,min=0"`
	LockedAt      time.Time `json:"locked_at,omitempty" validate:"omitempty,datetime"`
}
