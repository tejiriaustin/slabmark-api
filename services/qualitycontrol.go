package services

import (
	"context"
	"errors"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type QualityControlService struct {
}

func NewQualityControlService() *QualityControlService {
	return &QualityControlService{}
}

type (
	CreateQualityRecordInput struct {
		ProductCode    string                         `json:"product_code"`
		OverallRemark  string                         `json:"overall_remark"`
		AccountInfo    *models.AccountInfo            `json:"account_info"`
		HourlyReadings []models.HourlyQualityReadings `json:"hourly_readings"`
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

func (s *QualityControlService) CreateQualityRecord(
	ctx context.Context,
	input CreateQualityRecordInput,
	dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
) (*models.DailyQualityReadings, error) {

	now := time.Now()

	dailyReadings := models.DailyQualityReadings{
		Shared: models.Shared{
			ID:        primitive.NewObjectID(),
			CreatedAt: &now,
		},
		ProductCode:    input.ProductCode,
		OverallRemark:  input.OverallRemark,
		AccountInfo:    *input.AccountInfo,
		HourlyReadings: input.HourlyReadings,
	}

	dailyQualityRecord, err := dailyQualityRepo.Create(ctx, dailyReadings)
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

func (s *QualityControlService) GetDailyQualityRecord(
	ctx context.Context,
	input GetQualityRecordInput,
	dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
) (*models.DailyQualityReadings, error) {

	id, err := primitive.ObjectIDFromHex(input.ID)
	if err != nil {
		return nil, errors.New("invalid identifier")
	}

	filter := repository.
		NewQueryFilter().
		AddFilter(models.FieldId, id)

	report, err := dailyQualityRepo.FindOne(ctx, filter, nil, nil)
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (s *QualityControlService) ListQualityRecords(
	ctx context.Context,
	input ListQualityReportsInput,
	dailyQualityRepo *repository.Repository[models.DailyQualityReadings],
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

	report, _, err := dailyQualityRepo.Paginate(ctx, filter, input.PerPage, input.Page, input.Projection, input.Sort)
	if err != nil {
		return nil, nil, err
	}

	return report, nil, nil
}
