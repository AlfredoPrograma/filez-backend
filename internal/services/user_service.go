package services

import (
	"github.com/alfredoprograma/filez-server/internal/domain"
	"github.com/alfredoprograma/filez-server/internal/repositories"
)

type UserService interface {
	GetByEmail(email string) (domain.PublicUser, error)
	Create(data domain.CreateUserDTO) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(usersRepo repositories.UserRepository) UserService {
	return &userService{usersRepo}
}

func (s userService) GetByEmail(email string) (domain.PublicUser, error) {
	user, err := s.userRepo.GetByEmail(email)

	if err != nil {
		return *new(domain.PublicUser), err
	}

	return domain.ToPublicUser(user), nil
}

func (s userService) Create(data domain.CreateUserDTO) error {
	user := domain.FromCreateUserDTO(data)
	return s.userRepo.Create(user)
}
