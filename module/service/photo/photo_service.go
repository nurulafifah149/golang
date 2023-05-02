package photo

import (
	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/model/photo"
)

type PhotoService interface {
	GetAll(ctx *gin.Context) ([]photo.Photo, error)
	GetById(ctx *gin.Context, idPhoto int) (photo.Photo, error)
	Create(ctx *gin.Context, photoIn photo.Photo) (photoOut photo.Photo, err error)
	Update(ctx *gin.Context, photoIn photo.Photo, idUser int) (photoOut photo.Photo, err error)
	Delete(ctx *gin.Context, idPhoto int, idUser int) error
}
