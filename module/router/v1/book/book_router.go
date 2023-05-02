package book

import (
	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/controller"
)

func BookRouter(v1 *gin.RouterGroup, bc controller.BookController) {

	g := v1.Group("/book")
	g.GET("", bc.GetAllBook)
	g.GET("/:id", bc.GetBookById)
	g.DELETE("/:id", bc.DeleteBook)
	g.POST("", bc.CreateBook)
	g.PUT("/:id", bc.UpdateBook)
}
