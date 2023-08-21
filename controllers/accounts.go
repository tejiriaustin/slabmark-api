package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/requests"
	"github.com/tejiriaustin/slabmark-api/response"
	"github.com/tejiriaustin/slabmark-api/services"
	"github.com/tejiriaustin/slabmark-api/utils"
)

type AccountsController struct {
}

func NewAccountController() *AccountsController {
	return &AccountsController{}
}

func (c *AccountsController) SignUp(
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.CreateUserRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		input := services.AddAccountInput{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
		}

		user, err := acctService.SignInUser(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 2", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", user)
	}
}

func (c *AccountsController) Login(
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.LoginUserRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		input := services.LoginUserInput{
			Username: req.Username,
			Password: req.Password,
		}

		user, err := acctService.LoginUser(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 2", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", user)
	}
}

func (c *AccountsController) ForgotPassword(
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.ForgotPasswordRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		input := services.ForgotPasswordInput{
			Email: req.EmailAddress,
		}

		user, err := acctService.ForgotPassword(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 2", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", user)
	}
}

func (c *AccountsController) ResetPassword(
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.ResetPasswordRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		input := services.ResetPasswordInput{
			NewPassword: req.NewPassword,
			ResetCode:   req.ResetCode,
		}

		user, err := acctService.ResetPassword(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 2", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", user)
	}
}

func (c *AccountsController) GetRoles(
	rolesService services.DepartmentsServiceInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		roles, err := rolesService.GetRoles()
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 2", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", roles)
	}
}

func (c *AccountsController) AddAccount(
	passwordGen utils.StrGenFunc,
	roleService services.DepartmentsServiceInterface,
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.AddAccountRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		if !roleService.IsValidDepartment(req.Department) {
			response.FormatResponse(ctx, http.StatusBadRequest, "The role added is not supported yet. Please contact support", nil)
		}

		input := services.AddAccountInput{
			FirstName:  req.FirstName,
			LastName:   req.LastName,
			Email:      req.Email,
			Phone:      req.Phone,
			Department: req.Department,
		}

		user, err := acctService.AddAccount(ctx, input, passwordGen, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 2", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", user)
	}
}

func (c *AccountsController) EditAccount(
	roleService services.DepartmentsServiceInterface,
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.EditAccountRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		if !roleService.IsValidDepartment(req.Department) {
			response.FormatResponse(ctx, http.StatusBadRequest, "The role added is not supported yet. Please contact support", nil)
		}

		input := services.EditAccountInput{
			FirstName:  req.FirstName,
			LastName:   req.LastName,
			Email:      req.Email,
			Phone:      req.Phone,
			Department: req.Department,
		}

		user, err := acctService.EditAccount(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 2", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", user)
	}
}
