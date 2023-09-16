package services

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"time"
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
		Id         string
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

func (s *AccountsService) CreateUser(ctx context.Context,
	input AddAccountInput,
	accountsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	filters := []map[string]interface{}{
		{"email": input.Email},
		{"phone": input.Phone},
	}

	qf := repository.NewQueryFilter().AddFilter("$or", filters)
	matchedUser, err := accountsRepo.FindOne(ctx, qf, nil, nil)
	if err != nil && err != repository.NoDocumentsFound {
		return nil, err
	}
	if matchedUser.Email == input.Email {
		return nil, errors.New("user with this email already exists")
	}
	if matchedUser.Phone == input.Phone {
		return nil, errors.New("user with this phone number already exists")
	}

	account := models.Account{
		Shared: models.Shared{
			ID: primitive.NewObjectID(),
		},
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Phone:     input.Phone,
		Email:     input.Email,
		Status:    "ACTIVE",
		Password:  input.Password,
	}
	account.FullName = account.GetFullName()
	account.Username = account.GetUsername()
	_, err = accountsRepo.Create(ctx, account)
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

	filters := []map[string]interface{}{
		{"email": input.Email},
		{"phone": input.Phone},
	}

	qf := repository.NewQueryFilter().AddFilter("$or", filters)
	matchedUser, err := accountsRepo.FindOne(ctx, qf, nil, nil)
	if err != nil && err != repository.NoDocumentsFound {
		return nil, err
	}
	if matchedUser.Email == input.Email {
		return nil, errors.New("user with this email already exists")
	}
	if matchedUser.Phone == input.Phone {
		return nil, errors.New("user with this phone number already exists")
	}

	randPassword := passwordGen()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(randPassword), 8)
	if err != nil {
		return nil, errors.New("couldn't generate password")
	}

	account := models.Account{
		Username:   input.FirstName,
		FirstName:  input.FirstName,
		LastName:   input.LastName,
		Phone:      input.Phone,
		Email:      input.Email,
		Department: input.Department,
		Status:     models.ActiveStatus,
		Password:   string(passwordHash),
	}

	account.FullName = account.GetFullName()
	account.Username = account.GetUsername()

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
		fields[models.FieldAccountFirstName] = input.FirstName
	}
	if input.LastName != "" {
		fields[models.FieldAccountLastName] = input.LastName
	}
	if input.Email != "" {
		fields[models.FieldAccountEmail] = input.Email
	}
	if input.Phone != "" {
		fields[models.FieldAccountPhone] = input.Phone
	}
	if input.Department != "" {
		fields[models.FieldAccountDepartment] = input.Department
	}
	updates := map[string]interface{}{
		"$set": fields,
	}

	filter := repository.NewQueryFilter().AddFilter(models.FieldId, input.Id)
	err := accountsRepo.UpdateMany(ctx, filter, updates)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AccountsService) LoginUser(ctx context.Context,
	input LoginUserInput,
	accountsRepo *repository.Repository[models.Account],
) (*models.Account, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return nil, err
	}

	filter := repository.NewQueryFilter()

	filter.
		AddFilter(models.FieldAccountUsername, input.Username).
		AddFilter(models.FieldAccountPassword, hashedPassword)

	account, err := accountsRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("incorrect login credentials")
		}
		return nil, errors.New("an error occurred")
	}

	token, err := s.generateAuthToken(ctx, models.AccountInfo{
		Id:         account.ID.Hex(),
		FirstName:  account.FirstName,
		LastName:   account.LastName,
		FullName:   account.FullName,
		Email:      account.Email,
		Department: account.Department,
	})
	if err != nil {
		return nil, errors.New("an error occurred")
	}

	account.Token = token
	return &account, nil
}

func (s *AccountsService) generateAuthToken(ctx context.Context, account models.AccountInfo) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(3600 * time.Minute)
	claims["authorized"] = true
	claims["account_info"] = account

	tokenString, err := token.SignedString(s.conf.GetAsString(env.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
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
