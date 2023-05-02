package product

import (
	"github.com/gin-gonic/gin"
	hdlproduct "github.com/nurulafifah149/golang/module/handle/product"
	"github.com/nurulafifah149/golang/pkg/middleware"
)

func ProductRoute(v1 *gin.RouterGroup, hdlSoc hdlproduct.ProductHandler) {
	s := v1.Group("/product")
	s.Use(middleware.JWTMiddleware())
	s.GET("/getall", hdlSoc.GetAll)
	s.GET("/getone/:id", middleware.AuthorizationById(), hdlSoc.GetById)
	s.POST("/createproduct", hdlSoc.Create)
	s.PUT("/updateproduct/:id", middleware.AuthorizationByRole(), hdlSoc.Update)
	s.DELETE("/deleteproduct/:id", middleware.AuthorizationByRole(), hdlSoc.Delete)

}
