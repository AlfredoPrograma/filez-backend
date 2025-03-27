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
	usersRepo repositories.UserRepository
}

func NewUserService(usersRepo repositories.UserRepository) UserService {
	return &userService{usersRepo}
}

func (s userService) GetByEmail(email string) (*domain.PublicUser, error) {
	return s.usersRepo.GetByEmail(email)
}

func (s userService) Create(user domain.CreateUserDTO) error {
	return s.usersRepo.Create(user)
}
