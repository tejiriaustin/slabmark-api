package services

import "context"

type Service struct {
}

type IServiceInterface interface{}

func NewService(ctx context.Context) IServiceInterface {
	return &Service{}
}
