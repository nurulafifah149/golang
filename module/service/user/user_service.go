package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/model/token"
	ModelUser "github.com/nurulafifah149/golang/module/model/user"
)

type UserService interface {
	CreateUser(ctx *gin.Context, userIn ModelUser.UserCreateRequest) (UserOut ModelUser.UserCreateResponse, err error)
	AuthenticateUser(ctx *gin.Context, userIn ModelUser.UserAuthenticate) (tokenOut token.Tokens, err error)
}
