package services

import (
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/utils"
	"log"
)

type Service struct {
	AccountsService   AccountsServiceInterface
	DeptService       DepartmentsServiceInterface
	StoreService      StoreServiceInterface
	LabService        LabServiceInterface
	PasswordGenerator utils.StrGenFunc
}

func NewService(conf *env.Environment) *Service {
	log.Println("Creating Service...")
	return &Service{
		AccountsService: NewAccountsService(conf),
		StoreService:    NewStoreService(conf),
		LabService:      NewLabService(conf),
		DeptService:     NewDeptService(),
	}
}
