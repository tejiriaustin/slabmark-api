package controllers

import "context"

type Controller struct {
}

func BuildNewController(ctx context.Context) *Controller {
	return &Controller{}
}
