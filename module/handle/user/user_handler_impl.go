package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/model/user"
	usrSVC "github.com/nurulafifah149/golang/module/service/user"
	"github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

type UserHandlerImpl struct {
	UserService usrSVC.UserService
}

func NewUserHandler(USvc usrSVC.UserService) UserHandler {
	return &UserHandlerImpl{
		UserService: USvc,
	}
}

// @Summary Register example
// @Schemes
// @Description do register
// @Tags user
// @Accept json
// @Produce json
// @Param	request	body	user.UserCreateRequest	true	"Input Data User"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=user.UserCreateResponse}
// @Failure      400  {object}  responseTemplate.WebResponseFailed{} "Input tidak valid"
// @Failure      500  {object}  responseTemplate.WebResponseFailed{} "Error Server Side"
// @Router /user/register [post]
func (u *UserHandlerImpl) Register(ctx *gin.Context) {
	//register begin
	logger.LogMyApp("i", "User Handler Invoked", "UserHandler - Register", nil)

	//catch data
	logger.LogMyApp("i", "Binding JSON data", "UserHandler - Register", nil)
	var dataInput user.UserCreateRequest
	err := ctx.BindJSON(&dataInput)
	if err != nil {
		logger.LogMyApp("e", "Error When Binding JSON data", "UserHandler - Register", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	//hit service
	logger.LogMyApp("i", "Hit User Service", "UserHandler - Register", nil)
	UserDataResp, err := u.UserService.CreateUser(ctx, dataInput)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit User Service", "UserHandler - Register", err)
		return
	}

	//Selesai
	logger.LogMyApp("i", "rendering response User Service", "UserHandler - Register", nil)
	ctx.JSON(http.StatusCreated, responseTemplate.WebResponseSuccess{
		Message: "Success Create Account",
		Data:    UserDataResp,
	})
	logger.LogMyApp("i", "User Service Success", "UserHandler - Register", nil)
}

// @Summary Login example
// @Schemes
// @Description do login
// @Tags user
// @Accept json
// @Produce json
// @Param	request	body	user.UserAuthenticate	true	"Login User"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=token.Tokens}
// @Failure      400  {object}  responseTemplate.WebResponseFailed{} "Input tidak valid"
// @Failure      500  {object}  responseTemplate.WebResponseFailed{} "Error Server Side"
// @Router /user/login [post]
func (u *UserHandlerImpl) Login(ctx *gin.Context) {
	logger.LogMyApp("i", "User Handler Invoked", "UserHandler - Login", nil)

	//catch data
	logger.LogMyApp("i", "Binding JSON data", "UserHandler - Login", nil)
	var dataInput user.UserAuthenticate
	err := ctx.BindJSON(&dataInput)
	if err != nil {
		logger.LogMyApp("e", "Error When Binding JSON data", "UserHandler - Login", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	// call service\
	logger.LogMyApp("i", "Hit User Service", "UserHandler - Register", nil)
	dataResponse, err := u.UserService.AuthenticateUser(ctx, dataInput)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit User Service", "UserHandler - Login", err)
		return
	}

	logger.LogMyApp("i", "rendering response User Service", "UserHandler - Register", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Login",
		Data:    dataResponse,
	})
}
