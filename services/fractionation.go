package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
)

type FractionationService struct {
	conf *env.Environment
}

func NewFractionationService(conf *env.Environment) *FractionationService {
	return &FractionationService{
		conf: conf,
	}
}

type (
	CreateFractionationRecordInput struct {
		ResumptionStock models.ResumptionStock `json:"resumption_stock"`
		ClosingStock    models.ClosingStock    `Json:"closing_stock"`
		Filtration      models.Filtration      `json:"filtration" `
		Loading         models.Loading         `json:"loading"`
	}
	FractionationListFilters struct {
		Query string // for partial or general lookups
	}
	ListFractionationReportsInput struct {
		Pager
		Projection *repository.QueryProjection
		Sort       *repository.QuerySort
		Filters    FractionationListFilters
	}
)

func (s *FractionationService) CreateFractionationRecord(
	ctx context.Context,
	input CreateFractionationRecordInput,
	fractionationRepo *repository.Repository[models.FractionationReport],
) (*models.FractionationReport, error) {

	report := models.FractionationReport{
		ResumptionStock: input.ResumptionStock,
		ClosingStock:    input.ClosingStock,
		Filtration:      input.Filtration,
		Loading:         input.Loading,
	}

	report, err := fractionationRepo.Create(ctx, report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
