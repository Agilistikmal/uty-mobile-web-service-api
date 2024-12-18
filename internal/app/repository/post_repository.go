package repository

import (
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) Create(post *model.Post) (*model.Post, error) {
	err := r.db.Save(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) Update(id string, post *model.Post) (*model.Post, error) {
	err := r.db.Where("id = ?", id).Updates(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) Delete(id string) (*model.Post, error) {
	var post *model.Post
	err := r.db.Take(&post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Delete(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *PostRepository) FindByID(id string) (*model.Post, error) {
	var post *model.Post
	err := r.db.Preload("Author").Take(&post, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	post.Author.Password = ""

	return post, nil
}

func (r *PostRepository) FindMany() []*model.Post {
	var posts []*model.Post
	r.db.Preload("Author").Order("updated_at DESC").Find(&posts)

	for _, post := range posts {
		post.Author.Password = ""
	}

	return posts
}
