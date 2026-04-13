package service

import (
	"errors"

	"finora-wealthlab/internal/model"
	"finora-wealthlab/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: &repository.UserRepository{},
	}
}

func (s *UserService) Register(email, password string) error {
	// check email
	existing, _ := s.repo.FindByEmail(email)
	if existing != nil {
		return errors.New("email already exists")
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashed),
	}

	return s.repo.Create(user)
}
