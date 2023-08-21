package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
)

type StoreService struct {
	conf *env.Environment
}

func NewStoreService(conf *env.Environment) *StoreService {
	return &StoreService{
		conf: conf,
	}
}

type AddItemInput struct {
	ItemName string
	Qty      int
}

func (s *StoreService) AddItem(ctx context.Context, input AddItemInput) (*models.StoreItem, error) {
	return nil, nil
}
