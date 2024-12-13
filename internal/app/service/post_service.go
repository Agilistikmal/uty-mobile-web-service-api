package service

import (
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/repository"
	"github.com/go-playground/validator/v10"
)

type PostService struct {
	postRepository *repository.PostRepository
	validate       *validator.Validate
}

func NewPostService(postRepository *repository.PostRepository, validate *validator.Validate) *PostService {
	return &PostService{
		postRepository: postRepository,
		validate:       validate,
	}
}

func (s *PostService) Create(request *model.PostCreateRequest) (*model.Post, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return nil, err
	}

	post := &model.Post{
		Title:          request.Title,
		Content:        request.Content,
		AuthorUsername: request.AuthorUsername,
	}

	post, err = s.postRepository.Create(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) Update(id string, request *model.PostUpdateRequest) (*model.Post, error) {
	err := s.validate.Struct(request)
	if err != nil {
		return nil, err
	}

	post := &model.Post{
		Title:   request.Title,
		Content: request.Content,
	}

	post, err = s.postRepository.Update(id, post)
	if err != nil {
		return nil, err
	}

	post, err = s.postRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) Delete(id string) (*model.Post, error) {
	post, err := s.postRepository.Delete(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) FindByID(id string) (*model.Post, error) {
	post, err := s.postRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *PostService) FindMany() []*model.Post {
	posts := s.postRepository.FindMany()

	return posts
}
