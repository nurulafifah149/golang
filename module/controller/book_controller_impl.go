package controller

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nurulafifah149/golang/module/helper"
	"github.com/nurulafifah149/golang/module/model"
	"github.com/nurulafifah149/golang/module/service"
)

type BookControllerImpl struct {
	BookService service.BookService
}

var validate = validator.New()

func NewBookController(bs service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bs,
	}
}

// @BasePath /api/v1

// Book godoc
// @Summary Book
// @Schemes
// @Description Get all data Book
// @Tags Books
// @Accept json
// @Produce json
// @Param	request	body	model.BookCreateRequest	true	"Input Data Buku"
// @Success 200 {object} model.WebResponse{data=model.BookResponse} "desc"
// @Router /book [post]
func (bc *BookControllerImpl) CreateBook(ctx *gin.Context) {
	var req model.BookCreateRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		helper.CatchError(ctx, "BR")
		return
	}
	var resp model.WebResponse
	resp.Message = "Success Create Book"

	err = validate.Struct(req)
	if err != nil {
		helper.CatchError(ctx, errors.New("BR").Error())
		return
	}

	resp.Data, err = bc.BookService.Insert(ctx, req)
	if err != nil {
		helper.CatchError(ctx, "BR")
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

// Book godoc
// @Summary Book
// @Schemes
// @Description Get all data Book
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.BookResponse} "desc"
// @Router /book [get]
func (bc *BookControllerImpl) GetAllBook(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	var response model.WebResponse

	data := bc.BookService.GetAll(ctx)
	response.Data = data
	response.Message = "Sukses Mengambil Data Buku"
	ctx.JSON(http.StatusOK, response)

}

// Book godoc
// @Summary Book
// @Schemes
// @Description Get data Book
// @Tags Books
// @Accept json
// @Produce json
// @Param        id    path     int  false  "book search by id"
// @Success 200 {object} model.WebResponse{data=model.BookResponse} "Sukses Mengambil data Buku"
// @Router /book/{id} [get]
func (bc *BookControllerImpl) GetBookById(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	var response model.WebResponse
	parameter := ctx.Param("id")
	var id int
	id, err := strconv.Atoi(parameter)

	if err != nil {
		helper.CatchError(ctx, "BR")
		return
	}

	data, err := bc.BookService.GetById(ctx, id)
	if err != nil {
		helper.CatchError(ctx, err.Error())
		return
	}

	response.Message = "Sukses Mengambil Data Buku"
	response.Data = data
	ctx.JSON(http.StatusOK, response)

}

// Book godoc
// @Summary Book
// @Schemes
// @Description Get data Book
// @Tags Books
// @Accept json
// @Produce json
// @Param	id		path 	int  false  "book update by id"
// @Param	request	body	model.BookCreateRequest	true	"Input Data Buku"
// @Success 200 {object} model.WebResponse{data=model.BookResponse} "Sukses Mengambil data Buku"
// @Router /book/{id} [put]
func (bc *BookControllerImpl) UpdateBook(ctx *gin.Context) {
	var req model.BookUpdateRequest

	err := ctx.BindJSON(&req)
	if err != nil {
		helper.CatchError(ctx, "BR")
		return
	}

	parameter := ctx.Param("id")
	var id int
	id, err = strconv.Atoi(parameter)

	if err != nil {
		helper.CatchError(ctx, "BR")
		return
	}

	req.Id = id

	var resp model.WebResponse
	resp.Message = "Success Update Book"
	err = validate.Struct(req)
	if err != nil {
		helper.CatchError(ctx, errors.New("BR").Error())
		return
	}

	data, err := bc.BookService.Update(ctx, req)

	if err != nil {
		helper.CatchError(ctx, err.Error())
		return
	}

	resp.Data = data
	ctx.JSON(http.StatusOK, resp)
}

// Book godoc
// @Summary Book
// @Schemes
// @Description Delete data Book By Id
// @Tags Books
// @Accept json
// @Produce json
// @Param        id    path     int  false  "Delete book by id"
// @Success 200 {object} model.WebResponse{} "desc"
// @Router /book/{id} [delete]
func (bc *BookControllerImpl) DeleteBook(ctx *gin.Context) {
	var response model.WebResponse
	parameter := ctx.Param("id")
	var id int
	id, err := strconv.Atoi(parameter)

	if err != nil {
		helper.CatchError(ctx, "BR")
		return
	}

	err = bc.BookService.Delete(ctx, id)
	if err != nil {
		helper.CatchError(ctx, err.Error())
		return
	}

	response.Message = "Sukses Menghapus Data Buku"
	response.Data = nil
	ctx.JSON(http.StatusOK, response)
}
