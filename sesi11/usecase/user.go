package usecase

import (
	"database/sql"
	"fmt"
)

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) UserUsecase {
	return UserUsecase{
		repo: repo,
	}
}

func (u UserUsecase) GetAll() ([]User, error) {
	users, err := u.repo.GetAll()
	if err != nil {
		if err == sql.ErrNoRows {
			return []User{}, fmt.Errorf("no data")
		}
	}

	return users, nil
}
