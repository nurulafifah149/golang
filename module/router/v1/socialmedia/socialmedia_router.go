package socialmedia

import (
	"github.com/gin-gonic/gin"
	hdlsocmed "github.com/nurulafifah149/golang/module/handle/socialmedia"
	"github.com/nurulafifah149/golang/pkg/middleware"
)

func SocialmediaRoute(v1 *gin.RouterGroup, hdlSoc hdlsocmed.SocialmediaHandler) {
	s := v1.Group("/socialmedia")
	s.Use(middleware.JWTMiddleware())
	s.GET("/getall", hdlSoc.GetAll)
	s.GET("/getone/:id", hdlSoc.GetById)
	s.POST("/createsocialmedia", hdlSoc.Create)
	s.PUT("/updatesocialmedia/:id", hdlSoc.Update)
	s.DELETE("/deleteSocialmedia/:id", hdlSoc.Delete)

}
