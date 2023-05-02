package comment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/helper"
	modelComment "github.com/nurulafifah149/golang/module/model/comment"
	svcComment "github.com/nurulafifah149/golang/module/service/comment"
	"github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

type CommentHandlerImpl struct {
	CommentSVC svcComment.CommentService
}

func NewCommenHandler(cSVC svcComment.CommentService) CommentHandler {
	return &CommentHandlerImpl{
		CommentSVC: cSVC,
	}
}

// @Summary Get All Comment example
// @Schemes
// @Security Bearer
// @Description how to get all Comment
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=[]modelComment.Comment}
// @Router /comment/getall [get]
func (c *CommentHandlerImpl) GetAll(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Comment Handler Invoked", "CommentHandler - Getall", nil)

	logger.LogMyApp("i", "Hit Comment Service", "CommentHandler - Getall", nil)
	data, _ := c.CommentSVC.GetAll(ctx)

	logger.LogMyApp("i", "Render Response", "CommentHandler - Getall", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success GET all Comment",
		Data:    data,
	})

	return
}

// @Summary Get Comment example
// @Schemes
// @Security Bearer
// @Description how to get Comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id photo"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=modelComment.Comment}
// @Router /comment/getone/{id} [get]
func (c *CommentHandlerImpl) GetById(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Comment Handler Invoked", "CommentHandler - GetById", nil)
	idComment, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "CommentHandler - GetById", err)
		return
	}

	//hit service
	data, err := c.CommentSVC.GetById(ctx, idComment)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit Comment Service", "CommentHandler - GetById", err)
		return
	}

	//render
	logger.LogMyApp("i", "Render Response", "CommentHandler - GetById", nil)
	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success GET Comment",
		Data:    data,
	})

	return

}

// @Summary update Comment example
// @Schemes
// @Security Bearer
// @Description how to update Comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param        id    path     int  false  "Id Comment"
// @Param	request	body	modelComment.CommentCreateRequest	true	"Input Data Comment"
// @Success 200 {object} responseTemplate.WebResponseSuccess{data=modelComment.Comment}
// @Router /comment/updatecomment/{id} [put]
func (c *CommentHandlerImpl) Update(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Comment Handler Invoked", "CommentHandler - Update", nil)

	var data modelComment.CommentCreateRequest
	logger.LogMyApp("i", "GET ID Comment FROM PARAMS", "CommentHandler - Update", nil)
	idComment, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "CommentHandler - Update", err)
		return
	}

	logger.LogMyApp("i", "GET Comment DATA FROM JSON", "CommentHandler - Update", nil)
	err = ctx.BindJSON(&data)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "CommentHandler - Update", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "CommentHandler - Update", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT Comment SERVICE", "CommentHandler - Update", nil)
	dataResp, err := c.CommentSVC.Update(ctx, modelComment.Comment{
		Id:      idComment,
		PhotoId: data.PhotoId,
		Message: data.Message,
	}, accessClaim.AccessClaims.UserId)
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT Comment SERVICE", "CommentHandler - Update", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Update Comment",
		Data:    dataResp,
	})
	return
}

// @Summary create Comment example
// @Schemes
// @Security Bearer
// @Description how to Create Comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param	request	body	modelComment.CommentCreateRequest	true	"Input Data Comment"
// @Success 201 {object} responseTemplate.WebResponseSuccess{data=modelComment.Comment}
// @Router /comment/createcomment [post]
func (c *CommentHandlerImpl) Create(ctx *gin.Context) {
	//panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Comment Handler Invoked", "CommentHandler - Create", nil)

	var data modelComment.CommentCreateRequest
	logger.LogMyApp("i", "GET Comment DATA FROM JSON", "CommentHandler - Create", nil)
	err := ctx.BindJSON(&data)
	if err != nil {
		logger.LogMyApp("e", "Error When Get Params data", "CommentHandler - Create", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "CommentHandler - Create", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT Comment SERVICE", "CommentHandler - Create", nil)
	dataResp, err := c.CommentSVC.Create(ctx, modelComment.Comment{
		PhotoId: data.PhotoId,
		Message: data.Message,
		UserId:  accessClaim.AccessClaims.UserId,
	})
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT Comment SERVICE", "CommentHandler - Create", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Create Comment",
		Data:    dataResp,
	})

}

// @Summary delete Comment example
// @Schemes
// @Security Bearer
// @Description how to Delete Comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param        id    path     int  false  "id Comment"
// @Success 200 {object} responseTemplate.WebResponseSuccess{}
// @Router /comment/deletecomment/{id} [delete]
func (c *CommentHandlerImpl) Delete(ctx *gin.Context) {
	// panic("not implemented") // TODO: Implement
	logger.LogMyApp("i", "Comment Handler Invoked", "CommentHandler - Delete", nil)

	logger.LogMyApp("i", "GET ID Comment FROM PARAMS", "CommentHandler - Delete", nil)
	idComment, err := helper.GetIdAndConvertToInt(ctx)
	if err != nil {
		logger.LogMyApp("e", "Error When Hit GetIdAndConvertInt helper", "CommentHandler - Delete", err)
		return
	}

	logger.LogMyApp("i", "GET USERDATA FROM CTX", "CommentHandler - Delete", nil)
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	logger.LogMyApp("i", "HIT Comment SERVICE", "CommentHandler - Delete", nil)
	err = c.CommentSVC.Delete(ctx, idComment, accessClaim.AccessClaims.UserId)
	if err != nil {
		logger.LogMyApp("e", "Error WHEN HIT Comment SERVICE", "CommentHandler - Delete", err)
		return
	}

	ctx.JSON(http.StatusOK, responseTemplate.WebResponseSuccess{
		Message: "Success Delete Comment",
		Data:    nil,
	})
	return
}
