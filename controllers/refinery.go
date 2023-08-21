package controllers

import "github.com/gin-gonic/gin"

type RefineryController struct {
}

func NewRefineryController() *RefineryController {
	return &RefineryController{}
}

func (c *RefineryController) NewRefineryRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func (c *RefineryController) EditRefineryRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *RefineryController) GetRefineryRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (c *RefineryController) ListRefineryRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
