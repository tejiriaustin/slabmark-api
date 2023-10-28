package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/requests"
	"github.com/tejiriaustin/slabmark-api/response"
	"github.com/tejiriaustin/slabmark-api/services"
)

type QualityControlController struct {
	conf *env.Environment
}

func NewQualityControlController(conf *env.Environment) *QualityControlController {
	return &QualityControlController{
		conf: conf,
	}
}

func (c *QualityControlController) CreateQualityControlRecord(
	qcService services.QualityControlServiceInterface,
	qcDailyRepo *repository.Repository[models.DailyQualityReadings],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accountInfo, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		var requestBody requests.CreateQualityRecordRequest

		err = ctx.BindJSON(&requestBody)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		input := services.CreateQualityRecordInput{
			ProductCode:    requestBody.ProductCode,
			OverallRemark:  requestBody.OverallRemark,
			AccountInfo:    accountInfo,
			HourlyReadings: requestBody.HourlyReadings,
		}
		record, err := qcService.CreateQualityRecord(ctx, input, qcDailyRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
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
	qcDailyRepo *repository.Repository[models.DailyQualityReadings],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		input := services.GetQualityRecordInput{
			ID: ctx.Param("id"),
		}

		record, err := qcService.GetDailyQualityRecord(ctx, input, qcDailyRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
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

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		input := services.ListQualityReportsInput{
			Pager: services.Pager{
				Page:    services.GetPageNumberFromContext(ctx),
				PerPage: services.GetPerPageLimitFromContext(ctx),
			},
			Filters: services.QualityListFilters{
				Query: ctx.Query("query"),
			},
		}

		records, paginator, err := qcService.ListQualityRecords(ctx, input, qcDailyRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
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
