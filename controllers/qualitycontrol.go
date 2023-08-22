package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/repository"
	"github.com/tejiriaustin/slabmark-api/services"
)

type QualityControlController struct {
}

func NewQualityControlController() *QualityControlController {
	return &QualityControlController{}
}

func (c *QualityControlController) CreateQualityControlRecord(
	qcService services.QualityControlServiceInterface,
	qcRepository *repository.Repository[models.HourlyQualityReadings],
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func (c *QualityControlController) EditQualityRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *QualityControlController) GetQualityRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *QualityControlController) ListQualityRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
