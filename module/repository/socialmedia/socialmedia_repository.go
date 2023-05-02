package socialmedia

import (
	"context"

	"github.com/nurulafifah149/golang/module/model/socialmedia"
)

type SocialmediaRepository interface {
	GetAll(ctx context.Context) ([]socialmedia.Socialmedia, error)
	GetById(ctx context.Context, idSoc int) (socialmedia.Socialmedia, error)
	Create(ctx context.Context, socIn socialmedia.Socialmedia) (socialmedia.Socialmedia, error)
	Update(ctx context.Context, socIn socialmedia.Socialmedia) (socialmedia.Socialmedia, error)
	Delete(ctx context.Context, idSoc int) error
}
