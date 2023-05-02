package controller

import "github.com/gin-gonic/gin"

type BookController interface {
	CreateBook(ctx *gin.Context)
	GetAllBook(ctx *gin.Context)
	GetBookById(ctx *gin.Context)
	UpdateBook(ctx *gin.Context)
	DeleteBook(ctx *gin.Context)
}
