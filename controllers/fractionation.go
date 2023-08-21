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

type FractionationController struct {
}

func NewFractionationController() *FractionationController {
	return &FractionationController{}
}

func (c *FractionationController) CreateFractionationRecord(
	fractionationService services.FractionationServiceInterface,
	fractionationRepo *repository.Repository[models.FractionationReport],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var requestBody requests.CreateFractionationReportRequest

		err := ctx.BindJSON(&requestBody)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		input := services.CreateFractionationRecordInput{
			ResumptionStock: requestBody.ResumptionStock,
			ClosingStock:    requestBody.ClosingStock,
			Filtration:      requestBody.Filtration,
			Loading:         requestBody.Loading,
		}

		record, err := fractionationService.CreateFractionationRecord(ctx, input, fractionationRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", record)
		return
	}
}

func (c *FractionationController) UpdateFractionationRecord(
	fractionationService services.FractionationServiceInterface,
	fractionationRepo *repository.Repository[models.FractionationReport],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestBody requests.UpdateFractionationRecordRequest

		err := ctx.BindJSON(&requestBody)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		input := services.UpdateFractionationRecordInput{
			ResumptionStock: requestBody.ResumptionStock,
			ClosingStock:    requestBody.ClosingStock,
			Filtration:      requestBody.Filtration,
			Loading:         requestBody.Loading,
		}

		record, err := fractionationService.UpdateFractionationRecord(ctx, input, fractionationRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", record)
		return
	}
}

func (c *FractionationController) ListFractionationRecords(
	fractionationService services.FractionationServiceInterface,
	fractionationRepo *repository.Repository[models.FractionationReport],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := services.ListFractionationReportsInput{
			Pager: services.Pager{
				Page:    services.GetPageNumberFromContext(ctx),
				PerPage: services.GetPerPageLimitFromContext(ctx),
			},
			Filters: services.FractionationListFilters{
				Query: ctx.Param("query"),
			},
		}

		records, paginator, err := fractionationService.ListFractionationRecords(ctx, input, fractionationRepo)
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

func (c *FractionationController) GetFractionationRecord(
	fractionationService services.FractionationServiceInterface,
	fractionationRepo *repository.Repository[models.FractionationReport],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		input := services.GetFractionationRecordInput{
			ID: ctx.Param("id"),
		}

		record, err := fractionationService.GetFractionationRecord(ctx, input, fractionationRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, "Bad Request 1", nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", record)
		return
	}
}
