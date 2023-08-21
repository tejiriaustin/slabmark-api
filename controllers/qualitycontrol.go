package controllers

import "github.com/gin-gonic/gin"

type QualityControlController struct {
}

func NewQualityControlController() *QualityControlController {
	return &QualityControlController{}
}

func (c *QualityControlController) NewQualityRecord() gin.HandlerFunc {
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
