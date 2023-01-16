package usecase

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type repoMock struct{}

// GetAll implements UserRepo
func (repoMock) GetAll() ([]User, error) {
	return getAll()
}

var (
	repo UserRepo = repoMock{}

	svc    = NewUserUsecase(repo)
	getAll func() ([]User, error)
)

func TestGetAll_Success(t *testing.T) {
	getAll = func() ([]User, error) {
		return []User{
			{
				Name: "Reyhan",
			},
		}, nil
	}
	users, err := svc.GetAll()
	require.Nil(t, err)

	fmt.Println(users)
}
func TestGetAll_NoData(t *testing.T) {
	getAll = func() ([]User, error) {
		return []User{}, sql.ErrNoRows
	}
	_, err := svc.GetAll()
	require.Equal(t, "no data", err.Error())
}
