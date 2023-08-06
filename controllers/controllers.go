package controllers

import "context"

type Controller struct {
	AccountsController *AccountsController
}

func BuildNewController(ctx context.Context) *Controller {
	return &Controller{
		AccountsController: NewAccountController(),
	}
}
