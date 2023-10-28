package controllers

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/env"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/services"
	"github.com/tejiriaustin/slabmark-api/utils"
)

func AddRoutes(
	ctx context.Context,
	routerEngine *gin.Engine,
	sc *services.Container,
	repos *repository.Container,
	conf *env.Environment,
) {

	controllers := BuildNewController(ctx, conf)

	passwordGenerator := utils.RandomStringGenerator()

	r := routerEngine.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		response.FormatResponse(c, http.StatusOK, "OK", nil)
	})

	user := r.Group("/user")
	{
		user.POST("", controllers.AccountsController.AddAccount(passwordGenerator, sc.DeptService, sc.AccountsService, repos.AccountsRepo))
		user.PUT("", controllers.AccountsController.EditAccount(sc.DeptService, sc.AccountsService, repos.AccountsRepo))
		user.POST("/login", controllers.AccountsController.Login(sc.AccountsService, repos.AccountsRepo))
		user.POST("/forgot-password", controllers.AccountsController.ForgotPassword(sc.AccountsService, repos.AccountsRepo))
		user.POST("/reset-password", controllers.AccountsController.ResetPassword(sc.AccountsService, repos.AccountsRepo))
		user.GET("/")
		user.GET("/roles", controllers.AccountsController.GetRoles(sc.DeptService))
	}

	fractionation := r.Group("/fractionation")
	{
		fractionation.POST("", controllers.FractionationController.CreateFractionationRecord(sc.FractionationService, repos.FractionationRepo))
		fractionation.PUT("", controllers.FractionationController.UpdateFractionationRecord(sc.FractionationService, repos.FractionationRepo))
		fractionation.GET("/:id", controllers.FractionationController.GetFractionationRecord(sc.FractionationService, repos.FractionationRepo))
		fractionation.GET("/list", controllers.FractionationController.ListFractionationRecords(sc.FractionationService, repos.FractionationRepo))
	}

	refinery := r.Group("/refinery")
	{
		refinery.POST("", controllers.RefineryController.CreateRefineryRecord(sc.RefineryService, repos.RefineryRepo))
		refinery.PUT("", controllers.RefineryController.EditRefineryRecords())
		refinery.GET("/:id", controllers.RefineryController.GetRefineryRecord(sc.RefineryService, repos.RefineryRepo))
		refinery.GET("/list", controllers.RefineryController.ListRefineryRecords(sc.RefineryService, repos.RefineryRepo))
	}

	qualityControl := r.Group("/quality-control")
	{
		qualityControl.POST("", controllers.QualityControlController.CreateQualityControlRecord(sc.QualityControlService, repos.QualityRepo))
		qualityControl.PUT("", controllers.QualityControlController.EditQualityRecords())
		qualityControl.GET("/:id", controllers.QualityControlController.GetQualityRecord(sc.QualityControlService, repos.QualityRepo))
		qualityControl.GET("/list", controllers.QualityControlController.ListQualityRecords(sc.QualityControlService, repos.QualityRepo))
	}

}
