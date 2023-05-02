package photo

import (
	"context"

	"github.com/nurulafifah149/golang/module/model/photo"
)

type PhotoRepository interface {
	GetAll(ctx context.Context) (photoOut []photo.Photo, err error)
	GetById(ctx context.Context, idPhoto int) (photoOut photo.Photo, err error)
	Create(ctx context.Context, comIn photo.Photo) (photoOut photo.Photo, err error)
	Update(ctx context.Context, comIn photo.Photo) (photoOut photo.Photo, err error)
	Delete(ctx context.Context, idPhoto int) (err error)
}
