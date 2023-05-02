package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	ProductModel "github.com/nurulafifah149/golang/module/model/product"
	ProductRepo "github.com/nurulafifah149/golang/module/repository/product"
	MyLog "github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

type ProductServiceImpl struct {
	ProductRepo ProductRepo.ProductRepository
	Validate    *validator.Validate
}

func NewProductService(ProductRepo ProductRepo.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepo: ProductRepo,
		Validate:    validate,
	}
}

func (Cs *ProductServiceImpl) GetAll(ctx *gin.Context, idUser int, role string) (Photos []ProductModel.Product, err error) {
	//logging
	MyLog.LogMyApp("i", "Product Service Invoked", "ProductService - GetAll", nil)

	if role == "user" {
		Photos, err = Cs.ProductRepo.GetAllByUserId(ctx, idUser)
	} else if role == "admin" {
		Photos, err = Cs.ProductRepo.GetAll(ctx)
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.SomethingWentWrong,
			Error:   responseTemplate.SomethingWentWrong,
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "ProductService - GetAll", err)
		return
	}

	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "ProductService - GetAll", err)
		return
	}

	return
}

func (Cs *ProductServiceImpl) GetById(ctx *gin.Context, idProduct int) (photoOut ProductModel.Product, err error) {
	// panic("not implemented") // TODO: Implement
	MyLog.LogMyApp("i", "Product Service Invoked", "ProductService - GetById", nil)

	photoOut, err = Cs.ProductRepo.GetById(ctx, idProduct)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "ProductService - GetById", err)
		return
	}

	return
}

func (Cs *ProductServiceImpl) Create(ctx *gin.Context, photoIn ProductModel.Product) (photoOut ProductModel.Product, err error) {
	// panic("not implemented") // TODO: Implement
	MyLog.LogMyApp("i", "Product Service Invoked", "ProductService - Create", nil)

	//validasi input
	MyLog.LogMyApp("i", "Validating Process invoked", "ProductService - Create", nil)
	Cs.Validate = validator.New()
	err = Cs.Validate.Struct(photoIn)

	if err != nil {
		MyLog.LogMyApp("e", "Validating Process Error", "ProductService - Create", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	photoOut, err = Cs.ProductRepo.Create(ctx, photoIn)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "ProductService - Create", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}

func (Cs *ProductServiceImpl) Update(ctx *gin.Context, photoIn ProductModel.Product, idUser int) (photoOut ProductModel.Product, err error) {
	MyLog.LogMyApp("i", "Product Service Invoked", "ProductService - Update", nil)

	//autorisasi
	MyLog.LogMyApp("i", "Autorisasi kepemilikan Product", "ProductService - Update", nil)
	photoOut, err = Cs.ProductRepo.GetById(ctx, photoIn.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "ProductService - Update", err)
		return
	}

	MyLog.LogMyApp("i", "Hit Repository For Update Proccess", "ProductService - Update", nil)
	photoOut, err = Cs.ProductRepo.Update(ctx, photoIn)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "ProductService - Update", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}

func (Cs *ProductServiceImpl) Delete(ctx *gin.Context, idProduct int, idUser int) (err error) {
	MyLog.LogMyApp("i", "Product Service Invoked", "Product Service - Delete", nil)

	//autorisasi
	MyLog.LogMyApp("i", "Autorisasi kepemilikan Product", "ProductService - Delete", nil)

	MyLog.LogMyApp("i", "hit repository", "ProductService - Delete", nil)
	err = Cs.ProductRepo.Delete(ctx, idProduct)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "ProductService - Delete", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}
