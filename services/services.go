package services

import (
	"github.com/tejiriaustin/slabmark-api/env"
	"log"
)

type Service struct {
	AccountsService AccountsServiceInterface
	StoreService    StoreServiceInterface
	LabService      LabServiceInterface
}

func NewService(conf *env.Environment) *Service {
	log.Println("Creating Service...")
	return &Service{
		AccountsService: NewAccountsService(conf),
		StoreService:    NewStoreService(conf),
		LabService:      NewLabService(conf),
	}
}
