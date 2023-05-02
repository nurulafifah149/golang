package photo

import (
	"github.com/gin-gonic/gin"
	hdlphoto "github.com/nurulafifah149/golang/module/handle/photo"
	"github.com/nurulafifah149/golang/pkg/middleware"
)

func PhotoRoute(v1 *gin.RouterGroup, hdlSoc hdlphoto.PhotoHandler) {
	s := v1.Group("/photo")
	s.Use(middleware.JWTMiddleware())
	s.GET("/getall", hdlSoc.GetAll)
	s.GET("/getone/:id", hdlSoc.GetById)
	s.POST("/createphoto", hdlSoc.Create)
	s.PUT("/updatephoto/:id", hdlSoc.Update)
	s.DELETE("/deletephoto/:id", hdlSoc.Delete)

}
