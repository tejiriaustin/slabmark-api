package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/response"
	"github.com/tejiriaustin/slabmark-api/services"
)

func AddRoutes(
	ctx context.Context,
	routerEngine *gin.Engine,
	service *services.Service,
	repos *repository.Container,
) {

	controllers := BuildNewController(ctx)

	r := routerEngine.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		response.FormatResponse(c, http.StatusOK, "OK", nil)
	})

	r.POST("/signup", controllers.AccountsController.SignUp(service.AccountsService, repos.AccountsRepo))
}
