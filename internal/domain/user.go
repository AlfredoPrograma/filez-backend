package domain

import "gorm.io/gorm"

type userHiddenFields struct {
	Password string
}

type userPublicFields struct {
	FirstName string
	LastName  string
	Email     string
}

type User struct {
	gorm.Model
	userPublicFields
	userHiddenFields
}

type CreateUserDTO struct {
	userPublicFields
	userHiddenFields
}

type PublicUser struct {
	userPublicFields
}
