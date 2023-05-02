package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/helper"
	"github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
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
