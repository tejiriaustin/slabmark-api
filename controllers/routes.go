package controllers

import (
	"context"
	"github.com/tejiriaustin/slabmark-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/response"
	"github.com/tejiriaustin/slabmark-api/services"
)

func AddRoutes(
	ctx context.Context,
	routerEngine *gin.Engine,
	sc *services.Service,
	repos *repository.Container,
) {

	controllers := BuildNewController(ctx)

	passwordGenerator := utils.RandomStringGenerator()

	r := routerEngine.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		response.FormatResponse(c, http.StatusOK, "OK", nil)
	})

	user := r.Group("/user")
	{
		// TODO: Remove this
		user.POST("/signup", controllers.AccountsController.SignUp(sc.AccountsService, repos.AccountsRepo))

		user.POST("", controllers.AccountsController.AddAccount(passwordGenerator, sc.DeptService, sc.AccountsService, repos.AccountsRepo))
		user.PUT("", controllers.AccountsController.EditAccount(sc.DeptService, sc.AccountsService, repos.AccountsRepo))
		user.POST("/login", controllers.AccountsController.Login(sc.AccountsService, repos.AccountsRepo))
		user.POST("/forgot-password", controllers.AccountsController.ForgotPassword(sc.AccountsService, repos.AccountsRepo))
		user.POST("/reset", controllers.AccountsController.ResetPassword(sc.AccountsService, repos.AccountsRepo))
		user.GET("/roles", controllers.AccountsController.GetRoles(sc.DeptService))
	}

	fractionation := r.Group("/fractionation")
	{
		fractionation.POST("", controllers.FractionationController.CreateFractionationRecord(sc.FractionationService, repos.FractionationRepo))
		fractionation.PUT("", controllers.FractionationController.EditFractionationRecord())
		fractionation.GET("", controllers.FractionationController.GetFractionationRecord())
		fractionation.GET("/list", controllers.FractionationController.GetFractionationRecord())
	}

	refinery := r.Group("refinery")
	{
		refinery.POST("", controllers.RefineryController.NewRefineryRecord())
		refinery.PUT("", controllers.RefineryController.EditRefineryRecords())
		refinery.GET("", controllers.RefineryController.GetRefineryRecord())
		refinery.GET("/list", controllers.RefineryController.ListRefineryRecords())
	}

	qualityControl := r.Group("quality-control")
	{
		qualityControl.POST("", controllers.QualityControlController.NewQualityRecord())
		qualityControl.PUT("", controllers.QualityControlController.EditQualityRecords())
		qualityControl.GET("", controllers.QualityControlController.GetQualityRecord())
		qualityControl.GET("/list", controllers.QualityControlController.ListQualityRecords())
	}

}
