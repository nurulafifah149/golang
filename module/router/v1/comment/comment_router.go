package comment

import (
	"github.com/gin-gonic/gin"
	hdlcomment "github.com/nurulafifah149/golang/module/handle/comment"
	"github.com/nurulafifah149/golang/pkg/middleware"
)

func CommentRoute(v1 *gin.RouterGroup, hdlSoc hdlcomment.CommentHandler) {
	s := v1.Group("/comment")
	s.Use(middleware.JWTMiddleware())
	s.GET("/getall", hdlSoc.GetAll)
	s.GET("/getone/:id", hdlSoc.GetById)
	s.POST("/createcomment", hdlSoc.Create)
	s.PUT("/updatecomment/:id", hdlSoc.Update)
	s.DELETE("/deletecomment/:id", hdlSoc.Delete)

}
