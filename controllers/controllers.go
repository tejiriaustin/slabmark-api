package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"github.com/tejiriaustin/slabmark-api/services"
	"net/http"
)

type Controller struct {
	conf                     *env.Environment
	AccountsController       *AccountsController
	QualityControlController *QualityControlController
	FractionationController  *FractionationController
	RefineryController       *RefineryController
}

func BuildNewController(ctx context.Context, conf *env.Environment) *Controller {
	return &Controller{
		AccountsController:       NewAccountController(conf),
		QualityControlController: NewQualityControlController(conf),
		FractionationController:  NewFractionationController(conf),
		RefineryController:       NewRefineryController(conf),
	}
}

func (r *Controller) SetCookieHandlers(c *gin.Context, token string) {
	c.SetCookie("auth", token, -1, "/", r.conf.GetAsString(env.FrontendUrl), false, true)
	c.String(http.StatusOK, "Cookie has been set")
}

func (r *Controller) extractToken(ctx *gin.Context) {

}

func GetAccountInfo(ctx *gin.Context, jwtSecret []byte) (*models.AccountInfo, error) {
	tokenString, err := GetAuthHeader(ctx)
	if err != nil {
		return nil, err
	}
	claims := &services.Claims{}

	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.AccountInfo, nil
}

func GetAuthHeader(c *gin.Context) (string, error) {
	return c.GetHeader("x-token-user"), nil
}
