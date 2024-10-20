package model

import "time"

type User struct {
	Username  string    `json:"username,omitempty" gorm:"primaryKey" validate:"required,min=3,max=20"`
	FullName  string    `json:"full_name,omitempty" validate:"required,min=3,max=50"`
	Phone     string    `json:"phone,omitempty" validate:"required,e164"`
	Password  string    `json:"password,omitempty" validate:"required,min=8"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
