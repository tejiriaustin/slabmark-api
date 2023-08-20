package controllers

import "github.com/gin-gonic/gin"

type FractionationController struct {
}

func NewFractionationController() *FractionationController {
	return &FractionationController{}
}

func (c *FractionationController) NewFractionationRecord() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
