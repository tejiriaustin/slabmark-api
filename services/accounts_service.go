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

type CreateAccountInput struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Password  string
}

func (s *AccountsService) SignInUser(ctx context.Context,
	input CreateAccountInput,
	acctsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	account := models.Account{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
		Email:     input.Email,
		Kind:      "SLABMARK.ACCOUNT.KIND.ADMIN",
		Status:    "ACTIVE",
	}

	_, err := acctsRepo.Create(ctx, account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
