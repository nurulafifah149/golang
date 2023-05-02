package socialmedia

import (
	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/module/model/socialmedia"
)

type SocialmediaService interface {
	GetAll(ctx *gin.Context) ([]socialmedia.Socialmedia, error)
	GetById(ctx *gin.Context, idPhoto int) (socialmedia.Socialmedia, error)
	Create(ctx *gin.Context, socIn socialmedia.Socialmedia) (socOut socialmedia.Socialmedia, err error)
	Update(ctx *gin.Context, socIn socialmedia.Socialmedia, idUser int) (socOut socialmedia.Socialmedia, err error)
	Delete(ctx *gin.Context, idSoc int, idUser int) error
}
