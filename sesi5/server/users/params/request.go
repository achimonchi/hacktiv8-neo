package params

import "sesi5/server/users/models"

type UserCreate struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func (u *UserCreate) ParseToModel() *models.User {
	return &models.User{
		Name:  u.Name,
		Email: u.Email,
	}
}
