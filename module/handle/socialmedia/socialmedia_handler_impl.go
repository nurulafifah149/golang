package socialmedia

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/helper"
	modelSocialmedia "github.com/nurulafifah149/golang/module/model/socialmedia"
	svcSocialmedia "github.com/nurulafifah149/golang/module/service/socialmedia"
	"github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

type SocialmediaHandlerImpl struct {
	socmedSVC svcSocialmedia.SocialmediaService
}

func NewSocialmediaHandler(socmedSvc svcSocialmedia.SocialmediaService) SocialmediaHandler {
	return &SocialmediaHandlerImpl{
		socmedSVC: socmedSvc,
	}
}

// @Summary Get All Social Media example
// @Schemes
// @Security Bearer
// @Description how to get all social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=[]modelSocialmedia.Socialmedia}
// @Router /socialmedia/getall [get]
func (s *SocialmediaHandlerImpl) GetAll(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Socialmedia Handler Invoked", "SocialmediaHandler - Getall", nil)

	logger.LogMyApp("i", "Hit Socialmedia Service", "SocialmediaHandler - Getall", nil)
	data, _ := s.socmedSVC.GetAll(ctx)

	logger.LogMyApp("i", "Render Response", "SocialmediaHandler - Getall", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success GET all Social Media",
		Data:    data,
	})

	return
}

// @Summary Get Social Media example
// @Schemes
// @Security Bearer
// @Description how to get social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id social media"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=modelSocialmedia.Socialmedia}
// @Router /socialmedia/getone/{id} [get]
func (s *SocialmediaHandlerImpl) GetById(ctx *gin.Context) {
	logger.LogMyApp("i", "Socialmedia Handler Invoked", "SocialmediaHandler - GetById", nil)
	idSocialmedia, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "SocialmediaHandler - GetById", err)
		return
	}
	logger.LogMyApp("i", "Hit Socialmedia Service", "SocialmediaHandler - GetById", nil)
	data, err := s.socmedSVC.GetById(ctx, idSocialmedia)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit SocialMedia Service", "SocialmediaHandler - GetById", err)
		return
	}

	logger.LogMyApp("i", "Render Response", "SocialmediaHandler - GetById", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success GET Social Media",
		Data:    data,
	})

	return
}

// @Summary update Social Media example
// @Schemes
// @Security Bearer
// @Description how to update social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id social media"
// @Param	request	body	modelSocialmedia.SocialmediaCreateRequest	true	"Input Data User"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=modelSocialmedia.Socialmedia}
// @Router /socialmedia/updatesocialmedia/{id} [put]
func (s *SocialmediaHandlerImpl) Update(ctx *gin.Context) {
	logger.LogMyApp("i", "Socialmedia Handler Invoked", "SocialmediaHandler - GetById", nil)

	var dataRequest modelSocialmedia.SocialmediaCreateRequest
	logger.LogMyApp("i", "GET ID SOCIAL MEDIA FDROM PARAMS", "SocialmediaHandler - GetById", nil)
	idSocialmedia, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "GET SOCIAL MEDIA DATA FROM JSON", "SocialmediaHandler - GetById", nil)
	err = ctx.BindJSON(&dataRequest)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "SocialmediaHandler - GetById", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	///get userid from accestoken
	logger.LogMyApp("i", "GET USERDATA FROM CTX", "SocialmediaHandler - Update", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT SOCIALMEDIA SERVICE", "SocialmediaHandler - Update", nil)
	dataResp, err := s.socmedSVC.Update(ctx, modelSocialmedia.Socialmedia{
		Id:             idSocialmedia,
		Name:           dataRequest.Name,
		SocialMediaUrl: dataRequest.SocialMediaUrl,
	}, accessClaim.AccessClaims.UserId)

	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT SOCIALMEDIA SERVICE", "SocialmediaHandler - Update", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "SUCCESS UPDATE SOCIAL MEDIA",
		Data:    dataResp,
	})
	return
}

// @Summary Create Social Media example
// @Schemes
// @Security Bearer
// @Description how to create social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param	request	body	modelSocialmedia.SocialmediaCreateRequest	true	"Input Data Social Media"
// @Success 201 {object} responseTemplate.WebResponseSuccess{data=modelSocialmedia.Socialmedia}
// @Router /socialmedia/createsocialmedia/ [post]
func (s *SocialmediaHandlerImpl) Create(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Socialmedia Handler Invoked", "SocialmediaHandler - CREATE", nil)
	var data modelSocialmedia.Socialmedia

	// Get data
	logger.LogMyApp("i", "GET Socialmedia FROM USER REQUEST", "SocialmediaHandler - CREATE", nil)
	err := ctx.BindJSON(&data)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "SocialmediaHandler - GetById", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	//getdata USER from ctx
	logger.LogMyApp("i", "GET USERDATA FROM CTX", "SocialmediaHandler - Update", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}
	// Create Data
	data.UserId = accessClaim.AccessClaims.UserId
	logger.LogMyApp("i", "HIT SOCIALMEDIA REPOSITORY", "SocialmediaHandler - CREATE", nil)
	data, err = s.socmedSVC.Create(ctx, data)
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT SOCIALMEDIA SERVICE", "SocialmediaHandler - DELETE", err)
		return
	}

	//success
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Create Social Media",
		Data:    data,
	})
}

// @Summary Delete Social Media example
// @Schemes
// @Security Bearer
// @Description how to delete social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id social media"
// @Success 200 {object} responseTemplate.WebResponseSuccess{}
// @Router /socialmedia/deleteocialmedia/{id} [delete]
func (s *SocialmediaHandlerImpl) Delete(ctx *gin.Context) {
	logger.LogMyApp("i", "Socialmedia Handler Invoked", "SocialmediaHandler - DELETE", nil)

	logger.LogMyApp("i", "GET ID SOCIAL MEDIA FROM PARAMS", "SocialmediaHandler - DELETE", nil)
	idSocialmedia, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		return
	}
	///get userid from accestoken
	logger.LogMyApp("i", "GET USERDATA FROM CTX", "SocialmediaHandler - Delete", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT SOCIALMEDIA SERVICE", "SocialmediaHandler - DELETE", nil)
	err = s.socmedSVC.Delete(ctx, idSocialmedia, accessClaim.AccessClaims.UserId)
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT SOCIALMEDIA SERVICE", "SocialmediaHandler - DELETE", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Delete Social Media",
		Data:    nil,
	})
	return
}
