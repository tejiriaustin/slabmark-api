package services

import (
	"context"
	"errors"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
)

type RefineryService struct {
	conf *env.Environment
}

func NewRefineryService(conf *env.Environment) *RefineryService {
	return &RefineryService{
		conf: conf,
	}
}

type (
	CreateRefineryInput struct {
		PlantSituation string
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

	record := models.RefineryReport{
		HourlyReports:  input.HourlyReport,
		PlantSituation: input.PlantSituation,
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
	dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
) (*models.DailyQualityReadings, error) {

	filter := repository.
		NewQueryFilter().
		AddFilter(models.FieldId, input.ID)

	report, err := dailyQualityRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *RefineryService) ListRefineryRecords(
	ctx context.Context,
	input ListRefineryReportsInput,
	dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
) ([]models.DailyQualityReadings, *repository.Paginator, error) {

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

	report, _, err := dailyQualityRepo.Paginate(ctx, filter, input.PerPage, input.Page, input.Projection, input.Sort)
	if err != nil {
		return nil, nil, err
	}

	return report, nil, nil
}
