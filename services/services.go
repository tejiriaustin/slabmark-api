package services

import (
	"github.com/tejiriaustin/slabmark-api/env"
)

type Service struct {
	AccountsService AccountsServiceInterface
	StoreService    StoreServiceInterface
	LabService      LabServiceInterface
}

func NewService(conf *env.Environment) *Service {
	return &Service{
		AccountsService: NewAccountsService(conf),
		StoreService:    NewStoreService(conf),
		LabService:      NewLabService(conf),
	}
}
