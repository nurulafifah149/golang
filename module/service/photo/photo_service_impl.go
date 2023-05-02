package photo

import (
	"errors"
	"net/http"

	PhotoModel "github.com/nurulafifah149/golang/module/model/photo"
	photoRepo "github.com/nurulafifah149/golang/module/repository/photo"
	MyLog "github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PhotoServiceImpl struct {
	PhotoRepo photoRepo.PhotoRepository
	Validate  *validator.Validate
}

func NewPhotoService(photorepo photoRepo.PhotoRepository, validate *validator.Validate) PhotoService {
	return &PhotoServiceImpl{
		PhotoRepo: photorepo,
		Validate:  validate,
	}
}

func (Cs *PhotoServiceImpl) GetAll(ctx *gin.Context) (Photos []PhotoModel.Photo, err error) {
	//logging
	MyLog.LogMyApp("i", "Photo Service Invoked", "PhotoService - GetAll", nil)

	Photos, err = Cs.PhotoRepo.GetAll(ctx)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "PhotoService - GetAll", err)
		return
	}

	return
}

func (Cs *PhotoServiceImpl) GetById(ctx *gin.Context, idPhoto int) (photoOut PhotoModel.Photo, err error) {
	// panic("not implemented") // TODO: Implement
	MyLog.LogMyApp("i", "Photo Service Invoked", "PhotoService - GetById", nil)

	photoOut, err = Cs.PhotoRepo.GetById(ctx, idPhoto)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "PhotoService - GetById", err)
		return
	}

	return
}

func (Cs *PhotoServiceImpl) Create(ctx *gin.Context, photoIn PhotoModel.Photo) (photoOut PhotoModel.Photo, err error) {
	// panic("not implemented") // TODO: Implement
	MyLog.LogMyApp("i", "Comment Service Invoked", "PhotoService - Create", nil)

	//validasi input
	MyLog.LogMyApp("i", "Validating Process invoked", "PhotoService - Create", nil)
	Cs.Validate = validator.New()
	err = Cs.Validate.Struct(photoIn)

	if err != nil {
		MyLog.LogMyApp("e", "Validating Process Error", "PhotoService - Create", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	photoOut, err = Cs.PhotoRepo.Create(ctx, photoIn)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "PhotoService - Create", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}

func (Cs *PhotoServiceImpl) Update(ctx *gin.Context, photoIn PhotoModel.Photo, idUser int) (photoOut PhotoModel.Photo, err error) {
	MyLog.LogMyApp("i", "Photo Service Invoked", "PhotoService - Update", nil)

	//autorisasi
	MyLog.LogMyApp("i", "Autorisasi kepemilikan Photo", "PhotoService - Update", nil)
	photoOut, err = Cs.PhotoRepo.GetById(ctx, photoIn.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "PhotoService - Update", err)
		return
	}

	if photoOut.UserId != idUser {
		MyLog.LogMyApp("e", "Unauthorized", "PhotoService - Update", err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, responseTemplate.WebResponseFailed{
			Message: responseTemplate.Forbidden,
			Error:   responseTemplate.Forbidden,
		})
		err = errors.New("Unauthorized")
		return
	}

	MyLog.LogMyApp("i", "Hit Repository For Update Proccess", "PhotoService - Update", nil)
	photoOut, err = Cs.PhotoRepo.Update(ctx, photoIn)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "PhotoService - Update", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}

func (Cs *PhotoServiceImpl) Delete(ctx *gin.Context, idPhoto int, idUser int) (err error) {
	MyLog.LogMyApp("i", "Photo Service Invoked", "Photo Service - Delete", nil)

	//autorisasi
	MyLog.LogMyApp("i", "Autorisasi kepemilikan Photo", "PhotoService - Delete", nil)
	photoOut, err := Cs.PhotoRepo.GetById(ctx, idPhoto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "PhotoService - Delete", err)
		return
	}

	if photoOut.UserId != idUser {
		MyLog.LogMyApp("e", "Unauthorized", "PhotoService - Delete", err)
		err = errors.New("Unauthorized")
		ctx.AbortWithStatusJSON(http.StatusForbidden, responseTemplate.WebResponseFailed{
			Message: responseTemplate.Forbidden,
			Error:   responseTemplate.Forbidden,
		})
		return
	}

	MyLog.LogMyApp("i", "hit repository", "PhotoService - Delete", nil)
	err = Cs.PhotoRepo.Delete(ctx, idPhoto)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "PhotoService - Delete", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}
