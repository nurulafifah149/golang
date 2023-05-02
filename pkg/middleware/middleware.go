package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/nurulafifah149/golang/config"
	"github.com/nurulafifah149/golang/module/helper"
	repositoryProduct "github.com/nurulafifah149/golang/module/repository/product"
	svcProduct "github.com/nurulafifah149/golang/module/service/product"
	"github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	HeaderKey  string
	ContextKey string
)

func (h HeaderKey) String() string {
	return string(h)
}

func (h ContextKey) String() string {
	return string(h)
}

const (
	Authorization HeaderKey = "Authorization"

	AccessClaim ContextKey = "access_claim"
	BearerAuth  string     = "Bearer "
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.LogMyApp("i", "Middleware Invoked", "Middleware - JWTMiddleware", nil)

		logger.LogMyApp("i", "GET TOKEN FROM REQUEST HEADER", "Middleware - JWTMiddleware", nil)
		header := ctx.GetHeader(Authorization.String())
		if header == "" {
			logger.LogMyApp("e", "TOKEN IS NOT FOUND", "Middleware - JWTMiddleware", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseTemplate.WebResponseFailed{
				Message: responseTemplate.Unauthorized,
				Error:   "token is not found",
			})
			return
		}

		token := strings.Split(header, BearerAuth)
		if len(token) != 2 {
			logger.LogMyApp("e", "TOKEN IS NOT FOUND", "Middleware - JWTMiddleware", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseTemplate.WebResponseFailed{
				Message: responseTemplate.Unauthorized,
				Error:   "token is not found",
			})
			return
		}

		var claim helper.Claims
		err := helper.VerifyAndParse(token[1], &claim)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseTemplate.WebResponseFailed{
				Message: responseTemplate.Unauthorized,
				Error:   "invalid token",
			})
			return
		}
		ctx.Set(AccessClaim.String(), claim)
		ctx.Next()
	}
}

func AuthorizationById() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		idProduct, err := helper.GetIdAndConvertToInt(ctx)
		if err != nil {
			logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "AuthorizationMiddleware", err)
			return
		}

		var validate *validator.Validate
		//Load Model
		pgConn := config.NewPostgresGormConn()
		RProduct := repositoryProduct.NewProductRepository(pgConn)
		SProduct := svcProduct.NewProductService(RProduct, validate)
		data, err := SProduct.GetById(ctx, idProduct)
		if err != nil {
			logger.LogMyApp("e", "Error When Hit Photo Service", "ProductHandler - GetById", err)
			return
		}

		accessClaim, err := helper.GetIdentityFromCtx(ctx)
		if err != nil {
			return
		}

		if data.Id != accessClaim.AccessClaims.UserId && accessClaim.AccessClaims.Role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, responseTemplate.WebResponseFailed{
				Message: responseTemplate.Unauthorized,
				Error:   "You are not allowed to access this data",
			})
			return
		}

		ctx.Next()
	}
}

func AuthorizationByRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		accessClaim, err := helper.GetIdentityFromCtx(ctx)
		if err != nil {
			return
		}

		if accessClaim.AccessClaims.Role != "admin" {
			fmt.Println(accessClaim.AccessClaims.Role)
			ctx.AbortWithStatusJSON(http.StatusForbidden, responseTemplate.WebResponseFailed{
				Message: responseTemplate.Unauthorized,
				Error:   "You are not allowed to access this endpoint",
			})
			return
		}

		ctx.Next()
	}
}
