package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/services"
)

type ActivityController struct {
	conf *env.Environment
}

func NewActivityController(conf *env.Environment) *ActivityController {
	return &ActivityController{conf: conf}
}

func (c *ActivityController) GetActivity(
	activityService services.ActivityServiceInterface,
	activityRepo *repository.Repository[models.Activity],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		_, err := GetAccountInfo(ctx, c.conf.GetAsBytes(env.JwtSecret))
		if err != nil {
			response.FormatResponse(ctx, http.StatusUnauthorized, "Unauthorized access", nil)
			return
		}

		input := services.ListActivityInput{
			Pager: services.Pager{
				Page:    services.GetPageNumberFromContext(ctx),
				PerPage: services.GetPerPageLimitFromContext(ctx),
			},
			Filters: services.ActivityListFilters{
				Query: ctx.Query("query"),
			},
		}

		records, paginator, err := activityService.ListActivities(ctx, input, activityRepo)
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
