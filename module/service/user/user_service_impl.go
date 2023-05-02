package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nurulafifah149/golang/module/helper"
	"github.com/nurulafifah149/golang/module/model/token"
	ModelUser "github.com/nurulafifah149/golang/module/model/user"
	RepoUser "github.com/nurulafifah149/golang/module/repository/user"
	MyLog "github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

type UserServiceImpl struct {
	UserRepository RepoUser.UserRepository
	Validate       *validator.Validate
}

func NewUserService(userrepo RepoUser.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userrepo,
		Validate:       validate,
	}
}

func (u *UserServiceImpl) CreateUser(ctx *gin.Context, userIn ModelUser.UserCreateRequest) (UserOut ModelUser.UserCreateResponse, err error) {
	//logging
	MyLog.LogMyApp("i", "User Service Invoked", "UserService - CreateUser", nil)

	//konversi create request ke model
	MyLog.LogMyApp("i", "Parse Data from Request to Model", "UserService - CreateUser", nil)
	data := ModelUser.User{
		Username: userIn.Username,
		Password: userIn.Password,
		Email:    userIn.Email,
		Age:      userIn.Age,
		Role:     userIn.Role,
	}

	//validasi input
	MyLog.LogMyApp("i", "Validating Process invoked", "UserService - CreateUser", nil)
	u.Validate = validator.New()
	err = u.Validate.Struct(data)

	if err != nil {
		MyLog.LogMyApp("e", "Validating Process Error", "UserService - CreateUser", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	//hashing Password
	MyLog.LogMyApp("i", "Hashing Password invoked", "UserService - CreateUser", nil)
	data.Password = helper.HashPass(data.Password)

	//hit repository
	MyLog.LogMyApp("i", "Hit UserRepository.Create ", "UserService - CreateUser", nil)
	err = u.UserRepository.Create(ctx, data)
	if err != nil {
		MyLog.LogMyApp("e", "Creating User is error", "UserService - CreateUser", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}
	UserOut.Username = data.Username
	UserOut.Email = data.Email
	UserOut.Age = data.Age
	return UserOut, nil
}

func (u *UserServiceImpl) AuthenticateUser(ctx *gin.Context, userIn ModelUser.UserAuthenticate) (token token.Tokens, err error) {
	MyLog.LogMyApp("i", "User Service Invoked", "UserService - AuthenticateUser", nil)
	data := ModelUser.User{
		Username: userIn.Username,
	}

	//get user by username
	MyLog.LogMyApp("i", "Hit Repository to Get user by username", "UserService - AuthenticateUser", nil)
	dataUser, err := u.UserRepository.GetByUsername(ctx, data.Username)

	if err != nil {
		MyLog.LogMyApp("e", "User Is Not Found", "UserService - AuthenticateUser", err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		return
	}

	//compare password
	MyLog.LogMyApp("i", "Comparing Password", "UserService - AuthenticateUser", nil)
	err = helper.ComparePass(userIn.Password, dataUser.Password)

	//render
	if err != nil {
		MyLog.LogMyApp("e", "Invalid Username Or Password", "UserService - AuthenticateUser", err)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseTemplate.WebResponseFailed{
			Message: "Invalid Username Or Password",
			Error:   responseTemplate.Unauthorized,
		})
		return
	}

	//Generating token
	MyLog.LogMyApp("i", "Generating Token", "UserService - AuthenticateUser", nil)
	token, err = helper.GenerateToken(dataUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   "Gagal Generate Token",
		})
		MyLog.LogMyApp("e", "Failed Generate Token", "UserService - AuthenticateUser", err)
		return
	}

	return
}
