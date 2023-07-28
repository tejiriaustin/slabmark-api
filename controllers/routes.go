package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/response"
)

func BuildRoutes(
	ctx context.Context,
	routerEngine *gin.Engine,
) {

	//controllers := BuildNewController(ctx)

	r := routerEngine.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		response.FormatResponse(c, http.StatusOK, "OK", nil)
	})
}
