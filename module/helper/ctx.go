package helper

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

func GetIdentityFromCtx(ctx *gin.Context) (dataOut Claims, err error) {
	accessClaimI, ok := ctx.Get("access_claim")
	if !ok {
		err = errors.New("error get claim from context")
		logger.LogMyApp("e", "ERROR WHEN GET USERDATA FROM CTX", "Helper GetIdentityFromCtx", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidPayload,
			Error:   "INVALID USERID",
		})
		return
	}

	if err = ObjectMapper(accessClaimI, &dataOut); err != nil {
		logger.LogMyApp("e", "ERROR WHEN MAPPING OBJECT", "SocialmediaHandler - Update", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidPayload,
			Error:   "invalid payload",
		})
		return
	}

	return
}

func GetIdAndConvertToInt(ctx *gin.Context) (idOut int, err error) {
	id := ctx.Param("id")
	idOut, err = strconv.Atoi(id)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "Helper - GetIdAndConvertToInt", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}
	return
}
