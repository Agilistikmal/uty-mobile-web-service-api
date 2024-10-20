package service

import (
	"fmt"
	"time"

	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/model"
	"github.com/agilistikmal/uty-mobile-web-service-api/internal/app/repository"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
	validate       *validator.Validate
}

func NewUserService(userRepository *repository.UserRepository, validate *validator.Validate) *UserService {
	return &UserService{
		userRepository: userRepository,
		validate:       validate,
	}
}

func (s *UserService) Register(user *model.User) (*model.User, error) {
	err := s.validate.Struct(user)
	if err != nil {
		return nil, err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashPassword)
	user.Verified = false

	user, err = s.userRepository.Create(user)
	return user, err
}

func (s *UserService) Login(username string, password string) (*model.User, error) {
	user, err := s.userRepository.Find(username)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	difference := now.Sub(user.LockedAt)
	retrySeconds := 10

	if difference < time.Duration(retrySeconds)*time.Second {
		return nil, fmt.Errorf("account locked, please wait %ds", retrySeconds-int(difference.Seconds()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		maxRetry := 3
		if *user.PasswordRetry < maxRetry {
			*user.PasswordRetry += 1
		} else {
			*user.PasswordRetry = 0
			user.LockedAt = time.Now()
		}
		_, err := s.userRepository.Update(user.Username, user)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("invalid password (%d/3)", *user.PasswordRetry)
	}

	return user, nil
}
