package photo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/helper"
	modelPhoto "github.com/nurulafifah149/golang/module/model/photo"
	svcPhoto "github.com/nurulafifah149/golang/module/service/photo"
	"github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

type PhotoHandlerImpl struct {
	PhotoSVC svcPhoto.PhotoService
}

func NewPhotoHandler(photoSVC svcPhoto.PhotoService) PhotoHandler {
	return &PhotoHandlerImpl{
		PhotoSVC: photoSVC,
	}
}

// @Summary Get All Photo example
// @Schemes
// @Security Bearer
// @Description how to get all Photo
// @Tags Photos
// @Accept json
// @Produce json
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=[]modelPhoto.Photo}
// @Router /photo/getall [get]
func (p *PhotoHandlerImpl) GetAll(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Photo Handler Invoked", "PhotoHandler - Getall", nil)

	logger.LogMyApp("i", "Hit Photo Service", "PhotoHandler - Getall", nil)
	data, _ := p.PhotoSVC.GetAll(ctx)

	logger.LogMyApp("i", "Render Response", "PhotoHandler - Getall", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success GET all Photo",
		Data:    data,
	})

	return
}

// @Summary Get Photo example
// @Schemes
// @Security Bearer
// @Description how to get Photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id photo"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=modelPhoto.Photo}
// @Router /photo/getone/{id} [get]
func (p *PhotoHandlerImpl) GetById(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Photo Handler Invoked", "PhotoHandler - GetById", nil)
	idPhoto, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "PhotoHandler - GetById", err)
		return
	}

	//hit service
	data, err := p.PhotoSVC.GetById(ctx, idPhoto)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit Photo Service", "PhotoHandler - GetById", err)
		return
	}

	//render
	logger.LogMyApp("i", "Render Response", "PhotoHandler - GetById", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success GET Photo",
		Data:    data,
	})

	return

}

// @Summary update Photo example
// @Schemes
// @Security Bearer
// @Description how to update photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id photo"
// @Param	request	body	modelPhoto.PhotoCreateRequest	true	"Input Data Photo"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=modelPhoto.Photo}
// @Router /photo/updatephoto/{id} [put]
func (p *PhotoHandlerImpl) Update(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Photo Handler Invoked", "PhotoHandler - Update", nil)

	var data modelPhoto.PhotoCreateRequest
	logger.LogMyApp("i", "GET ID PHOTO FROM PARAMS", "PhotoHandler - Update", nil)
	idPhoto, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "PhotoHandler - Update", err)
		return
	}

	logger.LogMyApp("i", "GET PHOTO DATA FROM JSON", "PhotoHandler - Update", nil)
	err = ctx.BindJSON(&data)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "PhotoHandler - Update", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "PhotoHandler - Update", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT PHOTO SERVICE", "PhotoHandler - Update", nil)
	dataResp, err := p.PhotoSVC.Update(ctx, modelPhoto.Photo{
		Id:       idPhoto,
		Title:    data.Title,
		Caption:  data.Caption,
		PhotoUrl: data.PhotoUrl,
	}, accessClaim.AccessClaims.UserId)

	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT PHOTO SERVICE", "PhotoHandler - Update", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Update Photo",
		Data:    dataResp,
	})
	return
}

// @Summary create Photo example
// @Schemes
// @Security Bearer
// @Description how to Create photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param	request	body	modelPhoto.PhotoCreateRequest	true	"Input Data Photo"
// @Success 201 {object} responseTemplate.WebResponseSuccess{data=modelPhoto.Photo}
// @Router /photo/createphoto [post]
func (p *PhotoHandlerImpl) Create(ctx *gin.Context) {
	//panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Photo Handler Invoked", "PhotoHandler - Create", nil)

	var data modelPhoto.PhotoCreateRequest
	logger.LogMyApp("i", "GET PHOTO DATA FROM JSON", "PhotoHandler - Create", nil)
	err := ctx.BindJSON(&data)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "PhotoHandler - Create", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "PhotoHandler - Create", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT PHOTO SERVICE", "PhotoHandler - Create", nil)
	dataResp, err := p.PhotoSVC.Create(ctx, modelPhoto.Photo{
		UserId:   accessClaim.AccessClaims.UserId,
		Title:    data.Title,
		Caption:  data.Caption,
		PhotoUrl: data.PhotoUrl,
	})
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT PHOTO SERVICE", "PhotoHandler - Create", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Create Photo",
		Data:    dataResp,
	})

}

// @Summary delete Photo example
// @Schemes
// @Security Bearer
// @Description how to Delete photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id Photo"
// @Success 200 {object} responseTemplate.WebResponseSuccess{}
// @Router /photo/deletephoto/{id} [delete]
func (p *PhotoHandlerImpl) Delete(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Photo Handler Invoked", "PhotoHandler - Delete", nil)

	logger.LogMyApp("i", "GET ID PHOTO FROM PARAMS", "PhotoHandler - Delete", nil)
	idPhoto, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "PhotoHandler - Delete", err)
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "PhotoHandler - Delete", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT PHOTO SERVICE", "PhotoHandler - Delete", nil)
	err = p.PhotoSVC.Delete(ctx, idPhoto, accessClaim.AccessClaims.UserId)
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT PHOTO SERVICE", "PhotoHandler - Delete", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Delete Photo",
		Data:    nil,
	})
	return
}
