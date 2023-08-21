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

func (c *FractionationController) EditFractionationRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *FractionationController) ListFractionationRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *FractionationController) GetFractionationRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
