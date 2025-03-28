package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/alfredoprograma/filez-server/internal/domain"
	"github.com/alfredoprograma/filez-server/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type fakeUserRepo struct {
	mock.Mock
}

// Create implements repositories.UserRepository.
func (m *fakeUserRepo) Create(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// GetByEmail implements repositories.UserRepository.
func (m *fakeUserRepo) GetByEmail(email string) (domain.User, error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}

type fakeEncryptService struct {
	mock.Mock
}

// Hash implements services.EncryptService.
func (m *fakeEncryptService) Hash(plain string) (string, error) {
	args := m.Called(plain)
	return args.String(0), args.Error(1)
}

// Verify implements services.EncryptService.
func (m *fakeEncryptService) Verify(hashed string, plain string) error {
	args := m.Called(plain)
	return args.Error(0)
}

func TestUserService_Create(t *testing.T) {
	tests := []struct {
		name               string
		createUserDto      domain.CreateUserDTO
		mockUserRepo       func(m *fakeUserRepo)
		mockEncryptService func(m *fakeEncryptService)
		expectedErr        error
	}{
		{
			name: "creates a new user",
			createUserDto: domain.CreateUserDTO{
				CommonUserFields: domain.CommonUserFields{
					FirstName: "Alfredo",
					LastName:  "Arvelaez",
					Email:     "example@mail.com",
				},
				Password: "secr3tP4ssw0rd",
			},
			mockUserRepo: func(m *fakeUserRepo) {
				m.On("Create", mock.AnythingOfType("domain.User")).Return(nil)
			},
			mockEncryptService: func(m *fakeEncryptService) {
				m.On("Hash", mock.AnythingOfType("string")).Return("HASHED_PASSWORD", nil)
			},
			expectedErr: nil,
		},
		{
			name: "user repository returns error",
			createUserDto: domain.CreateUserDTO{
				CommonUserFields: domain.CommonUserFields{
					FirstName: "Alfredo",
					LastName:  "Arvelaez",
					Email:     "example@mail.com",
				},
				Password: "secr3tP4ssw0rd",
			},
			mockUserRepo: func(m *fakeUserRepo) {
				m.On("Create", mock.AnythingOfType("domain.User")).Return(errors.New("user repository error"))
			},
			mockEncryptService: func(m *fakeEncryptService) {
				m.On("Hash", mock.AnythingOfType("string")).Return("HASHED_PASSWORD", nil)
			},
			expectedErr: errors.New("user repository error"),
		},
		{
			name: "encrypt service returns error",
			createUserDto: domain.CreateUserDTO{
				CommonUserFields: domain.CommonUserFields{
					FirstName: "Alfredo",
					LastName:  "Arvelaez",
					Email:     "example@mail.com",
				},
				Password: "secr3tP4ssw0rd",
			},
			mockEncryptService: func(m *fakeEncryptService) {
				m.On("Hash", mock.AnythingOfType("string")).Return("", errors.New("encrypt service error"))
			},
			expectedErr: errors.New("encrypt service error"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			userRepo := new(fakeUserRepo)
			if tc.mockUserRepo != nil {
				tc.mockUserRepo(userRepo)
			}

			encryptService := new(fakeEncryptService)
			if tc.mockEncryptService != nil {
				tc.mockEncryptService(encryptService)
			}

			userService := services.NewUserService(userRepo, encryptService)

			err := userService.Create(tc.createUserDto)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}

			userRepo.AssertExpectations(t)
			encryptService.AssertExpectations(t)
		})
	}
}

func TestUserService_GetByEmail(t *testing.T) {
	date := time.Now()
	tests := []struct {
		name               string
		email              string
		mockUserRepo       func(m *fakeUserRepo)
		expectedPublicUser domain.PublicUser
		expectedErr        error
	}{
		{
			name:  "get user by email",
			email: "example@mail.com",
			mockUserRepo: func(m *fakeUserRepo) {
				m.On("GetByEmail", mock.AnythingOfType("string")).Return(domain.User{
					Model: domain.Model{
						ID:        1,
						CreatedAt: date,
						UpdatedAt: date,
					},
					CommonUserFields: domain.CommonUserFields{
						FirstName: "Alfredo",
						LastName:  "Arvelaez",
						Email:     "example@mail.com",
					},
					Password: "HASHED_PASSWORD",
				}, nil)
			},
			expectedPublicUser: domain.PublicUser{
				Model: domain.Model{
					ID:        1,
					CreatedAt: date,
					UpdatedAt: date,
				},
				CommonUserFields: domain.CommonUserFields{
					FirstName: "Alfredo",
					LastName:  "Arvelaez",
					Email:     "example@mail.com",
				},
			},
			expectedErr: nil,
		},
		{
			name:  "user repository returns error",
			email: "example@mail.com",
			mockUserRepo: func(m *fakeUserRepo) {
				m.On("GetByEmail", mock.AnythingOfType("string")).Return(domain.User{}, errors.New("user repository error"))
			},
			expectedPublicUser: domain.PublicUser{},
			expectedErr:        errors.New("user repository error"),
		},
	}

	for _, tc := range tests {
		userRepo := new(fakeUserRepo)
		if tc.mockUserRepo != nil {
			tc.mockUserRepo(userRepo)
		}
		encryptService := new(fakeEncryptService)

		userService := services.NewUserService(userRepo, encryptService)
		user, err := userService.GetByEmail(tc.email)

		if err != nil {
			assert.Error(t, err)
			assert.Equal(t, tc.expectedErr, err)
		} else {
			assert.Equal(t, tc.expectedPublicUser, user)
		}

		userRepo.AssertExpectations(t)
		encryptService.AssertExpectations(t)
	}
}
