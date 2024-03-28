package product

import (
	"context"
)

// UseCase ...
type Service interface {
	Add(context context.Context, product Product) (Product, error)
	GetIdByName(name string) (Product, error)
}

type service struct {
	repo Repository
}

func NewUseCase(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (us *service) Add(context context.Context, newProduct Product) (res Product, err error) {
	res, err = us.repo.Insert(context, newProduct)
	return res, err
}

func (us *service) GetIdByName(name string) (res Product, err error) {
	res, err = us.repo.GetIdByName(name)
	return res, err
}
