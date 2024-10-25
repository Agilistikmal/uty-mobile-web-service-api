package repository

import (
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"gorm.io/gorm"
)

// Struct User Repository
type UserRepository struct {
	// Field db untuk mengakses database
	db *gorm.DB
}

// Constructor untuk membuat user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Untuk membuat user baru
func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	// Argumen user akan dibuat ke database
	// lalu akan diperbarui datanya saat selesai dibuat
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Untuk mencari user berdasarkan username
func (r *UserRepository) Find(username string) (*model.User, error) {
	// Membuat variable untuk menyimpan data user
	var user *model.User
	err := r.db.Take(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Untuk mengupdate user
func (r *UserRepository) Update(username string, user *model.User) (*model.User, error) {
	err := r.db.Where("username = ?", username).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
