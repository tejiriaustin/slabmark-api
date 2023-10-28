package services

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
)

type FractionationService struct {
	conf *env.Environment
}

func NewFractionationService(conf *env.Environment) FractionationServiceInterface {
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
		AccountInfo     *models.AccountInfo    `json:"account_info"`
	}

	UpdateFractionationRecordInput struct {
		ID              string                 `json:"id"`
		ResumptionStock models.ResumptionStock `json:"resumption_stock"`
		ClosingStock    models.ClosingStock    `Json:"closing_stock"`
		Filtration      models.Filtration      `json:"filtration" `
		Loading         models.Loading         `json:"loading"`
	}

	GetFractionationRecordInput struct {
		ID string `json:"id"`
	}

	FractionationListFilters struct {
		Query string // for partial free hand lookups
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

	now := time.Now()

	report := models.FractionationReport{
		Shared: models.Shared{
			ID:        primitive.NewObjectID(),
			CreatedAt: &now,
		},
		ResumptionStock: input.ResumptionStock,
		ClosingStock:    input.ClosingStock,
		Filtration:      input.Filtration,
		Loading:         input.Loading,
		AccountInfo:     input.AccountInfo,
	}

	report, err := fractionationRepo.Create(ctx, report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *FractionationService) UpdateFractionationRecord(
	ctx context.Context,
	input UpdateFractionationRecordInput,
	fractionationRepo *repository.Repository[models.FractionationReport],
) (*models.FractionationReport, error) {

	recordId, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, err
	}
	report := models.FractionationReport{
		Shared: models.Shared{
			ID: recordId,
		},
		ResumptionStock: input.ResumptionStock,
		ClosingStock:    input.ClosingStock,
		Filtration:      input.Filtration,
		Loading:         input.Loading,
	}

	report, err = fractionationRepo.Update(ctx, report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *FractionationService) GetFractionationRecord(
	ctx context.Context,
	input GetFractionationRecordInput,
	fractionationRepo *repository.Repository[models.FractionationReport],
) (*models.FractionationReport, error) {

	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, errors.New("invalid identifier")
	}
	filter := repository.
		NewQueryFilter().
		AddFilter(models.FieldId, id)

	report, err := fractionationRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *FractionationService) ListFractionationRecords(
	ctx context.Context,
	input ListFractionationReportsInput,
	fractionationRepo *repository.Repository[models.FractionationReport],
) ([]models.FractionationReport, *repository.Paginator, error) {

	filter := repository.NewQueryFilter()

	if input.Filters.Query != "" {
		freeHandFilters := []map[string]interface{}{
			{"status": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"cr_batch_number": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"reference": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"reference": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
		}
		filter.AddFilter("$or", freeHandFilters)
	}

	report, _, err := fractionationRepo.Paginate(ctx, filter, input.PerPage, input.Page, input.Projection, input.Sort)
	if err != nil {
		return nil, nil, err
	}

	return report, nil, nil
}
