package services

import (
	"context"
	"errors"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RefineryService struct {
	conf *env.Environment
}

func NewRefineryService(conf *env.Environment) *RefineryService {
	return &RefineryService{
		conf: conf,
	}
}

var _ RefineryServiceInterface = (*RefineryService)(nil)

type (
	CreateRefineryInput struct {
		PlantSituation string
		AccountInfo    models.AccountInfo
		HourlyReport   []models.HourlyReport
	}

	UpdateRefineryRecordInput struct {
	}

	GetRefineryRecordInput struct {
		ID string `json:"id"`
	}

	RefineryListFilters struct {
		Query string // for partial free hand lookups
	}

	ListRefineryReportsInput struct {
		Pager
		Projection *repository.QueryProjection
		Sort       *repository.QuerySort
		Filters    RefineryListFilters
	}
)

func (r *RefineryService) CreateRefineryRecord(
	ctx context.Context,
	input CreateRefineryInput,
	refineryRepo *repository.Repository[models.RefineryReport],
) (*models.RefineryReport, error) {

	if input.HourlyReport == nil {
		return nil, errors.New("at least one hourly report is required")
	}

	now := time.Now()

	record := models.RefineryReport{
		Shared: models.Shared{
			ID:        primitive.NewObjectID(),
			CreatedAt: &now,
		},
		HourlyReports:  input.HourlyReport,
		PlantSituation: input.PlantSituation,
		AccountInfo:    input.AccountInfo,
	}

	record, err := refineryRepo.Create(ctx, record)
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (r *RefineryService) UpdateRefineryRecord(
	ctx context.Context,
	input UpdateRefineryRecordInput,
	refineryRepo *repository.Repository[models.RefineryReport],
) (*models.RefineryReport, error) {

	//recordId, err := primitive.ObjectIDFromHex(input.ID)
	//if err != nil {
	//	return nil, err
	//}
	//report := models.FractionationReport{
	//	Shared: models.Shared{
	//		ID: recordId,
	//	},
	//	ResumptionStock: input.ResumptionStock,
	//	ClosingStock:    input.ClosingStock,
	//	Filtration:      input.Filtration,
	//	Loading:         input.Loading,
	//}
	//
	//report, err = refineryRepo.Update(ctx, report)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return &report, nil

	return nil, nil
}

func (r *RefineryService) GetRefineryRecord(
	ctx context.Context,
	input GetRefineryRecordInput,
	refineryRepo *repository.Repository[models.RefineryReport],
) (*models.RefineryReport, error) {

	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, errors.New("invalid identifier")
	}

	filter := repository.
		NewQueryFilter().
		AddFilter(models.FieldId, id)

	report, err := refineryRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *RefineryService) ListRefineryRecords(
	ctx context.Context,
	input ListRefineryReportsInput,
	refineryRepo *repository.Repository[models.RefineryReport],
) ([]models.RefineryReport, *repository.Paginator, error) {

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

	report, _, err := refineryRepo.Paginate(ctx, filter, input.PerPage, input.Page, input.Projection, input.Sort)
	if err != nil {
		return nil, nil, err
	}

	return report, nil, nil
}
