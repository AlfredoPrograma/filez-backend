package repositories

import (
	"github.com/alfredoprograma/filez-server/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (domain.User, error)
	Create(user domain.User) error
}

type userRepository struct {
	db *gorm.DB
}

func (repo userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User

	query := repo.db.Where("email = ?", email).First(&user)

	if query.Error != nil {
		return user, query.Error
	}

	return user, nil
}

func (repo userRepository) Create(user domain.User) error {
	query := repo.db.Create(&user)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
