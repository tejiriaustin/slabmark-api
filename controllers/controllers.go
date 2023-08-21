package controllers

import "context"

type Controller struct {
	AccountsController       *AccountsController
	QualityControlController *QualityControlController
	FractionationController  *FractionationController
	RefineryController       *RefineryController
}

func BuildNewController(ctx context.Context) *Controller {
	return &Controller{
		AccountsController:       NewAccountController(),
		QualityControlController: NewQualityControlController(),
		FractionationController:  NewFractionationController(),
		RefineryController:       NewRefineryController(),
	}
}
