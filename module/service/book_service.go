package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/model"
)

type BookService interface {
	Insert(ctx *gin.Context, request model.BookCreateRequest) (model.BookResponse, error)
	Update(ctx *gin.Context, request model.BookUpdateRequest) (model.BookResponse, error)
	Delete(ctx *gin.Context, id int) error
	GetAll(ctx *gin.Context) []model.BookResponse
	GetById(ctx *gin.Context, id int) (model.BookResponse, error)
}
