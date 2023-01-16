package usecase

type UserRepo interface {
	GetAll() ([]User, error)
}
