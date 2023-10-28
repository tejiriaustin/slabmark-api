package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/constants"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/utils"
	"log"
)

type (
	Container struct {
		AccountsService       AccountsServiceInterface
		DeptService           DepartmentsServiceInterface
		RefineryService       RefineryServiceInterface
		FractionationService  FractionationServiceInterface
		QualityControlService QualityControlServiceInterface
		PasswordGenerator     utils.StrGenFunc
	}

	Pager struct {
		Page    int64
		PerPage int64
	}
)

func NewService(conf *env.Environment) *Container {
	log.Println("Creating Container...")
	return &Container{
		AccountsService:       NewAccountsService(conf),
		FractionationService:  NewFractionationService(conf),
		QualityControlService: NewQualityControlService(),
		RefineryService:       NewRefineryService(conf),
		DeptService:           NewDeptService(),
	}
}

func GetPageNumberFromContext(ctx context.Context) int64 {
	n, ok := ctx.Value(constants.ContextKeyPageNumber).(int64)
	if !ok {
		return 0
	}
	return n
}

func GetPerPageLimitFromContext(ctx context.Context) int64 {
	l, ok := ctx.Value(constants.ContextKeyPerPageLimit).(int64)
	if !ok {
		return 0
	}
	return l
}
