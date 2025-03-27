package domain

type CommonUserFields struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
type User struct {
	Model
	CommonUserFields
	Password string
}

type CreateUserDTO struct {
	CommonUserFields
	Password string
}

type PublicUser struct {
	Model
	CommonUserFields
}

func ToPublicUser(user User) PublicUser {
	return PublicUser{
		Model:            user.Model,
		CommonUserFields: user.CommonUserFields,
	}
}

func FromCreateUserDTO(dto CreateUserDTO) User {
	return User{
		CommonUserFields: dto.CommonUserFields,
		Password:         dto.Password,
	}
}
