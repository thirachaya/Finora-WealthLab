package service

import (
	"finora-wealthlab/internal/model"
	"finora-wealthlab/internal/repository"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Register(email, password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	user := &model.User{
		ID:       uuid.New(),
		Email:    email,
		Password: string(hashed),
	}

	return s.UserRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (bool, error) {
	user, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
