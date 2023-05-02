package product

import (
	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/model/product"
)

type ProductService interface {
	GetAll(ctx *gin.Context, idUser int, role string) ([]product.Product, error)
	GetById(ctx *gin.Context, idProduct int) (product.Product, error)
	Create(ctx *gin.Context, productIn product.Product) (productOut product.Product, err error)
	Update(ctx *gin.Context, productIn product.Product, idUser int) (productOut product.Product, err error)
	Delete(ctx *gin.Context, idProduct int, idUser int) error
}
