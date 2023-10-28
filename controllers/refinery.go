package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/requests"
	"github.com/tejiriaustin/slabmark-api/services"
	"net/http"
)

type RefineryController struct {
	conf *env.Environment
}

func NewRefineryController(conf *env.Environment) *RefineryController {
	return &RefineryController{
		conf: conf,
	}
}

func (c *RefineryController) CreateRefineryRecord(
	refineryService services.RefineryServiceInterface,
	refineryRepo *repository.Repository[models.RefineryReport],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accountInfo, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		var requestBody requests.CreateRefineryRecordRequest

		err = ctx.BindJSON(&requestBody)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		input := services.CreateRefineryInput{
			PlantSituation: requestBody.PlantSituation,
			HourlyReport:   requestBody.HourlyReports,
			AccountInfo:    *accountInfo,
		}

		record, err := refineryService.CreateRefineryRecord(ctx, input, refineryRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", record)
		return
	}
}

func (c *RefineryController) EditRefineryRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *RefineryController) GetRefineryRecord(
	refineryService services.RefineryServiceInterface,
	refineryRepo *repository.Repository[models.RefineryReport],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		input := services.GetRefineryRecordInput{
			ID: ctx.Param("id"),
		}

		record, err := refineryService.GetRefineryRecord(ctx, input, refineryRepo)
		if err != nil {
			response.FormatResponse(ctx, http.StatusBadRequest, err.Error(), nil)
			return
		}

		response.FormatResponse(ctx, http.StatusOK, "successful", record)
		return

	}
}

func (c *RefineryController) ListRefineryRecords(
	refineryService services.RefineryServiceInterface,
	refineryRepo *repository.Repository[models.RefineryReport],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		input := services.ListRefineryReportsInput{
			Pager: services.Pager{
				Page:    services.GetPageNumberFromContext(ctx),
				PerPage: services.GetPerPageLimitFromContext(ctx),
			},
			Filters: services.RefineryListFilters{
				Query: ctx.Query("query"),
			},
		}

		records, paginator, err := refineryService.ListRefineryRecords(ctx, input, refineryRepo)
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
