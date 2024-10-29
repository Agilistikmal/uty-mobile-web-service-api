package model

import "time"

type Payment struct {
	ID          string `json:"id,omitempty" gorm:"primaryKey"`
	ReferenceID string `json:"reference_id,omitempty" gorm:"unique"`
	Username    string `json:"username,omitempty"`
	Url         string `json:"url,omitempty"`
	Status      string `json:"status,omitempty"`
	Amount      int
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
