package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/utils"
)

type AccountsService struct {
	conf *env.Environment
}

func NewAccountsService(conf *env.Environment) *AccountsService {
	return &AccountsService{
		conf: conf,
	}
}

type (
	AddAccountInput struct {
		FirstName  string
		LastName   string
		Email      string
		Phone      string
		Password   string
		Department string
	}
	EditAccountInput struct {
		FirstName  string
		LastName   string
		Email      string
		Phone      string
		Department string
	}
	LoginUserInput struct {
		Username string
		Password string
	}
	ForgotPasswordInput struct {
		Email string
	}
	ResetPasswordInput struct {
		ResetCode   string
		NewPassword string
	}
)

func (s *AccountsService) SignInUser(ctx context.Context,
	input AddAccountInput,
	accountsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	account := models.Account{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
		Email:     input.Email,
		Status:    "ACTIVE",
		Password:  "whimpy-boy",
	}

	_, err := accountsRepo.Create(ctx, account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (s *AccountsService) AddAccount(ctx context.Context,
	input AddAccountInput,
	passwordGen utils.StrGenFunc,
	accountsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	randPassword := passwordGen()

	account := models.Account{
		Username:   input.FirstName,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Phone:      input.Phone,
		Email:      input.Email,
		Department: input.Department,
		Status:     models.ActiveStatus,
		Password:   randPassword,
	}

	// TODO: send notification to email and whatsapp

	account.FullName = account.GetFullName()
	account.Password = passwordGen()

	acct, err := accountsRepo.Create(ctx, account)
	if err != nil {
		return nil, err
	}
	return &acct, nil
}

func (s *AccountsService) EditAccount(ctx context.Context,
	input EditAccountInput,
	accountsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	fields := map[string]interface{}{}

	if input.FirstName != "" {
		fields["first_name"] = input.FirstName
	}
	if input.LastName != "" {
		fields["last_name"] = input.LastName
	}
	if input.Email != "" {
		fields["email"] = input.Email
	}
	if input.Phone != "" {
		fields["phone"] = input.Phone
	}
	if input.Department != "" {
		fields["department"] = input.Department
	}

	//TODO: implement DB update

	return nil, nil
}

func (s *AccountsService) LoginUser(ctx context.Context,
	input LoginUserInput,
	accountsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	filter := repository.NewQueryFilter()

	filter.
		AddFilter(models.FieldAccountUsername, input.Username).
		AddFilter(models.FieldAccountPassword, input.Password)

	account, err := accountsRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *AccountsService) ForgotPassword(ctx context.Context,
	input ForgotPasswordInput,
	accountsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	filter := repository.NewQueryFilter()

	filter.AddFilter(models.FieldAccountEmail, input.Email)

	account, err := accountsRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *AccountsService) ResetPassword(ctx context.Context,
	input ResetPasswordInput,
	accountsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	filter := repository.NewQueryFilter()

	account, err := accountsRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
