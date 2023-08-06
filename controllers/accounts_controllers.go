package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/requests"
	"github.com/tejiriaustin/slabmark-api/response"
	"github.com/tejiriaustin/slabmark-api/services"
	"net/http"
)

type AccountsController struct {
}

func NewAccountController() *AccountsController {
	return &AccountsController{}
}

func (c *AccountsController) SignUp(
	acctService services.AccountsServiceInterface,
	acctRepo *repository.Repository[models.Account],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req requests.CreateUserRequest

		err := ctx.BindJSON(&req)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		input := services.CreateAccountInput{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Phone:     req.Phone,
		}

		user, err := acctService.SignInUser(ctx, input, acctRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 2", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", user)
	}
}
