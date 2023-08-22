package services

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
)

type QualityControlService struct {
}

func NewQualityControlService() *QualityControlService {
	return &QualityControlService{}
}

type (
	CreateQualityRecordInput struct {
		HourlyReadings []models.HourlyQualityReadings
		DailyReadings  models.DailyQualityReadings
	}

	UpdateQualityRecordInput struct {
		ID              string                 `json:"id"`
		ResumptionStock models.ResumptionStock `json:"resumption_stock"`
		ClosingStock    models.ClosingStock    `Json:"closing_stock"`
		Filtration      models.Filtration      `json:"filtration"`
		Loading         models.Loading         `json:"loading"`
	}

	GetQualityRecordInput struct {
		ID string `json:"id"`
	}

	QualityListFilters struct {
		Query       string // for partial free hand lookups
		AccountInfo models.AccountInfo
	}

	ListQualityReportsInput struct {
		Pager
		Projection *repository.QueryProjection
		Sort       *repository.QuerySort
		Filters    QualityListFilters
	}
)

func (s *FractionationService) CreateQualityRecord(
	ctx context.Context,
	input CreateQualityRecordInput,
	hourlyQualityRepo *repository.Repository[models.HourlyQualityReadings],
	dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
) (*models.DailyQualityReadings, error) {

	recordIds, err := hourlyQualityRepo.CreateMany(ctx, input.HourlyReadings)
	if err != nil {
		return nil, err
	}

	input.DailyReadings.IdsOfReadings = recordIds

	dailyQualityRecord, err := dailyQualityRepo.Create(ctx, input.DailyReadings)
	if err != nil {
		return nil, err
	}

	return &dailyQualityRecord, nil
}

//func (s *FractionationService) UpdateQualityRecord(
//	ctx context.Context,
//	input UpdateFractionationRecordInput,
//	fractionationRepo *repository.Repository[models.FractionationReport],
//) (*models.FractionationReport, error) {
//
//	recordId, err := primitive.ObjectIDFromHex(input.ID)
//	if err != nil {
//		return nil, err
//	}
//	report := models.FractionationReport{
//		Shared: models.Shared{
//			ID: recordId,
//		},
//		ResumptionStock: input.ResumptionStock,
//		ClosingStock:    input.ClosingStock,
//		Filtration:      input.Filtration,
//		Loading:         input.Loading,
//	}
//
//	report, err = fractionationRepo.Update(ctx, report)
//	if err != nil {
//		return nil, err
//	}
//
//	return &report, nil
//}

func (s *FractionationService) GetDailyQualityRecord(
	ctx context.Context,
	input GetFractionationRecordInput,
	fractionationRepo *repository.Repository[models.DailyQualityReadings],
) (*models.DailyQualityReadings, error) {

	filter := repository.
		NewQueryFilter().
		AddFilter(models.FieldAccountId, input.ID)

	report, err := fractionationRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *FractionationService) GetHourlyQualityRecord(
	ctx context.Context,
	input GetFractionationRecordInput,
	hourlyQualityRepo *repository.Repository[models.HourlyQualityReadings],
) (*models.HourlyQualityReadings, error) {

	filter := repository.
		NewQueryFilter().
		AddFilter(models.FieldAccountId, input.ID)

	report, err := hourlyQualityRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *FractionationService) ListQualityRecords(
	ctx context.Context,
	input ListQualityReportsInput,
	fractionationRepo *repository.Repository[models.DailyQualityReadings],
) ([]models.DailyQualityReadings, *repository.Paginator, error) {

	filter := repository.NewQueryFilter()

	if input.Filters.Query != "" {
		freeHandFilters := []map[string]interface{}{
			{"product_code": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"overall_remark": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"account_info.first_name": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"account_info.last_name": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"account_info.email": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
			{"account_info.phone": map[string]interface{}{"$regex": input.Filters.Query, "$options": "i"}},
		}
		filter.AddFilter("$or", freeHandFilters)
	}

	report, _, err := fractionationRepo.Paginate(ctx, filter, input.PerPage, input.Page, input.Projection, input.Sort)
	if err != nil {
		return nil, nil, err
	}

	return report, nil, nil
}
