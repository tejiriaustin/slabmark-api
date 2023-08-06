package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
)

type AccountsService struct {
	conf *env.Environment
}

func NewAccountsService(conf *env.Environment) *AccountsService {
	return &AccountsService{
		conf: conf,
	}
}

type SignInInput struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
}

func (s *AccountsService) SignInUser(ctx context.Context, input SignInInput, repo *repository.Container) (*models.Account, error) {
	return nil, nil
}
