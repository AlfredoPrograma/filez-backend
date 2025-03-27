package services

import (
	"github.com/alfredoprograma/filez-server/internal/domain"
	"github.com/alfredoprograma/filez-server/internal/repositories"
)

type UserService interface {
	GetByEmail(email string) (*domain.PublicUser, error)
	Create(user domain.CreateUserDTO) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s userService) GetByEmail(email string) (*domain.PublicUser, error) {
	return s.userRepo.GetByEmail(email)
}

func (s userService) Create(user domain.CreateUserDTO) error {
	return s.userRepo.Create(user)
}
