package repository

type Repository struct{}

type IRepositoryInterface interface {
}

func NewRepository(name string) IRepositoryInterface {
	return &Repository{}
}
