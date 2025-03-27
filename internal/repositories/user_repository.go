package repositories

import (
	"github.com/alfredoprograma/filez-server/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*domain.PublicUser, error)
	Create(user domain.CreateUserDTO) error
}

type userRepository struct {
	db *gorm.DB
}

func (repo userRepository) GetByEmail(email string) (*domain.PublicUser, error) {
	var user domain.PublicUser

	query := repo.db.Where("email = ?", email).First(&user)

	if query.Error != nil {
		return nil, query.Error
	}

	return &user, nil
}

func (repo userRepository) Create(user domain.CreateUserDTO) error {
	query := repo.db.Create(&user)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
