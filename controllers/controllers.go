package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/tejiriaustin/slabmark-api/env"
	"github.com/tejiriaustin/slabmark-api/models"
	"net/http"
)

type Controller struct {
	conf                     *env.Environment
	AccountsController       *AccountsController
	QualityControlController *QualityControlController
	FractionationController  *FractionationController
	RefineryController       *RefineryController
}

func BuildNewController(ctx context.Context) *Controller {
	return &Controller{
		AccountsController:       NewAccountController(),
		QualityControlController: NewQualityControlController(),
		FractionationController:  NewFractionationController(),
		RefineryController:       NewRefineryController(),
	}
}

func (r *Controller) SetCookieHandlers(c *gin.Context, token string) {
	c.SetCookie("auth", token, -1, "/", r.conf.GetAsString(env.FrontendUrl), false, true)
	c.String(http.StatusOK, "Cookie has been set")
}

func (r *Controller) extractToken(ctx *gin.Context) {

}

func GetAccountInfo(ctx *gin.Context, jwtSecret string) (*models.AccountInfo, error) {
	tokenString, err := GetCookieHandler(ctx)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("there's an error with the signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, errors.New("error signing token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accountInfo := claims["account_info"].(models.AccountInfo)
		return &accountInfo, nil
	}
	return nil, fmt.Errorf("unable to get account")
}

func GetCookieHandler(c *gin.Context) (string, error) {
	cookie, err := c.Cookie("auth")
	if err != nil {
		return "", errors.New("cookie not found")
	}
	return cookie, nil
}
