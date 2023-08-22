package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/requests"
	"github.com/tejiriaustin/slabmark-api/response"
	"github.com/tejiriaustin/slabmark-api/services"
	"net/http"
)

type QualityControlController struct {
}

func NewQualityControlController() *QualityControlController {
	return &QualityControlController{}
}

func (c *QualityControlController) CreateQualityControlRecord(
	qcService services.QualityControlServiceInterface,
	qcHourlyRepo *repository.Repository[models.HourlyQualityReadings],
	qcDailyRepo *repository.Repository[models.DailyQualityReadings],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody requests.CreateQualityRecordRequest

		err := ctx.BindJSON(&requestBody)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		input := services.CreateQualityRecordInput{
			ProductCode:    requestBody.ProductCode,
			OverallRemark:  requestBody.OverallRemark,
			AccountInfo:    requestBody.AccountInfo,
			HourlyReadings: requestBody.HourlyReadings,
		}
		record, err := qcService.CreateQualityRecord(ctx, input, qcHourlyRepo, qcDailyRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", record)
		return
	}
}

func (c *QualityControlController) EditQualityRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *QualityControlController) GetQualityRecord(
	qcService services.QualityControlServiceInterface,
	qcHourlyRepo *repository.Repository[models.HourlyQualityReadings],
	qcDailyRepo *repository.Repository[models.DailyQualityReadings],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := services.GetQualityRecordInput{
			ID: ctx.Param("id"),
		}

		record, err := qcService.GetDailyQualityRecord(ctx, input, qcDailyRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", record)
		return
	}
}

func (c *QualityControlController) ListQualityRecords(
	qcService services.QualityControlServiceInterface,
	qcDailyRepo *repository.Repository[models.DailyQualityReadings],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := services.ListQualityReportsInput{
			Pager: services.Pager{
				Page:    services.GetPageNumberFromContext(ctx),
				PerPage: services.GetPerPageLimitFromContext(ctx),
			},
			Filters: services.QualityListFilters{
				Query: ctx.Param("query"),
			},
		}

		records, paginator, err := qcService.ListQualityRecords(ctx, input, qcDailyRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		payload := map[string]interface{}{
			"records": records,
			"meta":    paginator,
		}
		response.FormatResponse(ctx, http.StatusOK, "successful", payload)
		return

	}
}
