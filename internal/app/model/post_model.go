package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Content        string    `json:"content,omitempty"`
	AuthorUsername string    `json:"author_username,omitempty"`
	Author         *User     `json:"author,omitempty" gorm:"foreignKey:AuthorUsername;references:Username"`
	CreatedAt      time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

type PostCreateRequest struct {
	Title          string `validate:"required,min=3,max=100" json:"title,omitempty"`
	Content        string `validate:"required" json:"content,omitempty"`
	AuthorUsername string `validate:"required" json:"author_username,omitempty"`
}

type PostUpdateRequest struct {
	Title   string `validate:"omitempty,min=3,max=100" json:"title,omitempty"`
	Content string `validate:"omitempty" json:"content,omitempty"`
}
