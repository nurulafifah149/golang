package user

import (
	"github.com/gin-gonic/gin"
	hdluser "github.com/nurulafifah149/golang/module/handle/user"
)

func UserRouter(v1 *gin.RouterGroup, UserHdl hdluser.UserHandler) {
	g := v1.Group("/user")
	g.POST("/register", UserHdl.Register)
	g.POST("/login", UserHdl.Login)
}
