package services

import "golang.org/x/crypto/bcrypt"

type EncryptService interface {
	Hash(plain string) (string, error)
	Verify(hashed string, plain string) error
}

type encryptService struct{}

func NewEncryptService() EncryptService {
	return encryptService{}
}

func (s encryptService) Hash(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), 13)
	return string(hashed), err
}

func (s encryptService) Verify(hashed string, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}
