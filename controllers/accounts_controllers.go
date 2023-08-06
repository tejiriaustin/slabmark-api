package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/services"
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

	}
}
