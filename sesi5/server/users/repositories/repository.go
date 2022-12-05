package repositories

import "sesi5/server/users/models"

type UserRepo struct {
	users *[]models.User
}

func NewUserRepo(users *[]models.User) *UserRepo {
	return &UserRepo{
		users: users,
	}
}

func (u *UserRepo) AddUser(user models.User) error {
	users := *u.users
	users = append(users, user)
	u.users = &users
	return nil
}

func (u *UserRepo) GetUsers() *[]models.User {
	return u.users
}
