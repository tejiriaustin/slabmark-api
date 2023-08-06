package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
)

type AccountsServiceInterface interface {
	SignInUser(ctx context.Context, input SignInInput, repo repository.Container) (*models.Account, error)
}

type StoreServiceInterface interface {
	AddItem(ctx context.Context, input AddItemInput) (*models.StoreItem, error)
}
type LabServiceInterface interface {
	AddDailyReading(ctx context.Context, input AddDailyReadingInput) (*models.LabReading, error)
}
