package service

import (
	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/helper"
	"github.com/nurulafifah149/golang/module/model"
	"github.com/nurulafifah149/golang/module/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookService(BookRepo repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: BookRepo,
	}
}

func (bs *BookServiceImpl) Insert(ctx *gin.Context, request model.BookCreateRequest) (model.BookResponse, error) {
	var data model.Book
	data.Author = request.Author
	data.Desc = request.Desc
	data.Title = request.Title
	data, err := bs.BookRepository.CreateBook(ctx, data)
	if err != nil {
		return helper.BookDomainToResp(data), err
	}
	return helper.BookDomainToResp(data), nil
}

func (bs *BookServiceImpl) Update(ctx *gin.Context, request model.BookUpdateRequest) (bookOut model.BookResponse, err error) {
	var data model.Book
	data.Author = request.Author
	data.Desc = request.Desc
	data.Title = request.Title
	data.Id = request.Id

	_, err = bs.BookRepository.GetBookById(ctx, data.Id)
	if err != nil {
		return
	}

	data, err = bs.BookRepository.UpdateBook(ctx, data)
	if err != nil {
		return
	}
	bookOut = helper.BookDomainToResp(data)

	return
}

func (bs *BookServiceImpl) Delete(ctx *gin.Context, id int) error {
	err := bs.BookRepository.DeleteBook(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (bs *BookServiceImpl) GetAll(ctx *gin.Context) []model.BookResponse {
	data, err := bs.BookRepository.GetAllBook(ctx)
	if err != nil {
		panic(err)
	}

	var resp []model.BookResponse
	for _, vals := range data {
		resp = append(resp, helper.BookDomainToResp(vals))
	}
	return resp
}

func (bs *BookServiceImpl) GetById(ctx *gin.Context, id int) (model.BookResponse, error) {

	data, err := bs.BookRepository.GetBookById(ctx, id)

	if err != nil {
		return helper.BookDomainToResp(data), err
	}

	return helper.BookDomainToResp(data), err
}
