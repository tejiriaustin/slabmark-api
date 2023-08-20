package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
)

type LabService struct {
	conf *env.Environment
}

func NewLabService(conf *env.Environment) *LabService {
	return &LabService{
		conf: conf,
	}
}

type AddDailyReadingInput struct {
	ItemName string
	Qty      int
}

func (s *LabService) AddDailyReading(ctx context.Context, input AddDailyReadingInput) (*models.LabReading, error) {
	return nil, nil
}
