package comment

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nurulafifah149/golang/module/model/comment"
	CommentModel "github.com/nurulafifah149/golang/module/model/comment"
	CommentRepo "github.com/nurulafifah149/golang/module/repository/comment"
	MyLog "github.com/nurulafifah149/golang/pkg/logger"
	responseTemplate "github.com/nurulafifah149/golang/pkg/response"
)

type CommentServiceImpl struct {
	CommentRepo CommentRepo.CommentRepository
	Validate    *validator.Validate
}

func NewCommentService(comentrepo CommentRepo.CommentRepository, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepo: comentrepo,
		Validate:    validate,
	}
}

func (Cs *CommentServiceImpl) GetAll(ctx *gin.Context) (Comments []CommentModel.Comment, err error) {
	//logging
	MyLog.LogMyApp("i", "Comment Service Invoked", "CommentService - GetAll", nil)

	Comments, err = Cs.CommentRepo.GetAll(ctx)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "CommentService - GetAll", err)
		return
	}

	return
}

func (Cs *CommentServiceImpl) GetById(ctx *gin.Context, idComment int) (comOut CommentModel.Comment, err error) {
	// panic("not implemented") // TODO: Implement
	MyLog.LogMyApp("i", "Comment Service Invoked", "CommentService - GetById", nil)

	comOut, err = Cs.CommentRepo.GetById(ctx, idComment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "CommentService - GetById", err)
		return
	}

	return
}

func (Cs *CommentServiceImpl) Create(ctx *gin.Context, comIn CommentModel.Comment) (comOut comment.Comment, err error) {
	// panic("not implemented") // TODO: Implement
	MyLog.LogMyApp("i", "Comment Service Invoked", "CommentService - Create", nil)

	//validasi input
	MyLog.LogMyApp("i", "Validating Process invoked", "CommentService - Create", nil)
	Cs.Validate = validator.New()
	err = Cs.Validate.Struct(comIn)

	if err != nil {
		MyLog.LogMyApp("e", "Validating Process Error", "CommentService - Create", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InvalidBody,
			Error:   err.Error(),
		})
		return
	}

	MyLog.LogMyApp("i", "Hit Repository For Update Proccess", "CommentService - Create", nil)
	comOut, err = Cs.CommentRepo.Create(ctx, comIn)

	if err != nil {
		MyLog.LogMyApp("e", "Validating Process Error", "CommentService - Create", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		return
	}

	return
}

func (Cs *CommentServiceImpl) Update(ctx *gin.Context, comIn comment.Comment, idUser int) (comOut comment.Comment, err error) {
	MyLog.LogMyApp("i", "Comment Service Invoked", "CommentService - Update", nil)

	//autorisasi
	MyLog.LogMyApp("i", "Autorisasi kepemilikan comment", "CommentService - Update", nil)
	comOut, err = Cs.CommentRepo.GetById(ctx, comIn.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "CommentService - Update", err)
		return
	}

	if comOut.UserId != idUser {
		MyLog.LogMyApp("e", "Unauthorized", "CommentService - Update", err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, responseTemplate.WebResponseFailed{
			Message: responseTemplate.Forbidden,
			Error:   responseTemplate.Forbidden,
		})
		return
	}

	MyLog.LogMyApp("i", "Hit Repository For Update Proccess", "CommentService - Update", nil)
	comOut, err = Cs.CommentRepo.Update(ctx, comIn)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "CommentService - Update", err)
		return
	}

	return
}

func (Cs *CommentServiceImpl) Delete(ctx *gin.Context, idComment int, idUser int) (err error) {
	MyLog.LogMyApp("i", "Comment Service Invoked", "CommentService - Delete", nil)

	//autorisasi
	MyLog.LogMyApp("i", "Autorisasi kepemilikan comment", "CommentService - Delete", nil)
	comOut, err := Cs.CommentRepo.GetById(ctx, idComment)
	if err != nil {
		MyLog.LogMyApp("e", "Repository Returning Error", "CommentService - Delete", err)
		ctx.AbortWithStatusJSON(http.StatusNotFound, responseTemplate.WebResponseFailed{
			Message: "Data tidak di temukan",
			Error:   err.Error(),
		})
		return
	}

	if comOut.UserId != idUser {
		MyLog.LogMyApp("e", "Unauthorized", "CommentService - Delete", err)
		ctx.AbortWithStatusJSON(http.StatusForbidden, responseTemplate.WebResponseFailed{
			Message: responseTemplate.Forbidden,
			Error:   responseTemplate.Forbidden,
		})
		err = errors.New("Unauthorized")
		return
	}

	MyLog.LogMyApp("i", "hit repository", "CommentService - Delete", nil)
	err = Cs.CommentRepo.Delete(ctx, idComment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responseTemplate.WebResponseFailed{
			Message: responseTemplate.InternalServer,
			Error:   err.Error(),
		})
		MyLog.LogMyApp("e", "Repository Returning Error", "CommentService - Delete", err)
		return
	}

	return
}
