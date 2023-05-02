package comment

import (
	"context"

	"github.com/nurulafifah149/golang/module/model/comment"
)

type CommentRepository interface {
	GetAll(ctx context.Context) ([]comment.Comment, error)
	GetById(ctx context.Context, idComment int) (comment.Comment, error)
	Create(ctx context.Context, comIn comment.Comment) (comOut comment.Comment, err error)
	Update(ctx context.Context, comIn comment.Comment) (comOut comment.Comment, err error)
	Delete(ctx context.Context, idComment int) error
}
