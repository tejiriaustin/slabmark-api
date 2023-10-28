package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/requests"
	"github.com/tejiriaustin/slabmark-api/response"
	"github.com/tejiriaustin/slabmark-api/services"
	"github.com/tejiriaustin/slabmark-api/utils"
)

type AccountsController struct {
	conf *env.Environment
}

func NewAccountController(conf *env.Environment) *AccountsController {
	return &AccountsController{
		conf: conf,
	}
}

func (c *AccountsController) SignUp(
	passwordGen utils.StrGenFunc,
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.CreateUserRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request", nil)
			return
		}

		input := services.AddAccountInput{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
			Password:  req.Phone,
		}

		user, err := acctService.CreateUser(ctx, input, passwordGen, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(user))
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
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		input := services.LoginUserInput{
			Username: req.Username,
			Password: req.Password,
		}

		user, err := acctService.LoginUser(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(user))
	}
}

func (c *AccountsController) LogOut() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SetCookie("auth", "", -1, "/", c.conf.GetAsString(env.FrontendUrl), false, true)
		ctx.String(http.StatusOK, "Cookie has been deleted")
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
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		input := services.ForgotPasswordInput{
			Email: req.EmailAddress,
		}

		user, err := acctService.ForgotPassword(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(user))
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
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		input := services.ResetPasswordInput{
			NewPassword: req.NewPassword,
			ResetCode:   req.ResetCode,
		}

		user, err := acctService.ResetPassword(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(user))
	}
}

func (c *AccountsController) GetRoles(
	rolesService services.DepartmentsServiceInterface,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		roles, err := rolesService.GetRoles()
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
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

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		var req requests.AddAccountRequest

		err = ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
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

		user, err := acctService.CreateUser(ctx, input, passwordGen, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(user))
	}
}

func (c *AccountsController) EditAccount(
	roleService services.DepartmentsServiceInterface,
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		var req requests.EditAccountRequest

		err = ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
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
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", response.SingleAccountResponse(user))
	}
}
func (c *AccountsController) ListAccounts(
	acctService services.AccountsServiceInterface,
	accountsRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		var req requests.ListAccountsRequest

		err = ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		input := services.ListAccountReportsInput{
			Pager: services.Pager{
				Page:    services.GetPageNumberFromContext(ctx),
				PerPage: services.GetPerPageLimitFromContext(ctx),
			},
			Filters: services.AccountListFilters{
				Query: ctx.Param("query"),
			},
		}
		accounts, paginator, err := acctService.ListAccounts(ctx, input, accountsRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		payload := map[string]interface{}{
			"records": response.MultipleAccountResponse(accounts),
			"meta":    paginator,
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", payload)
	}
}
