package repository

import (
	"context"

	"github.com/nurulafifah149/golang/module/model"
)

type BookRepository interface {
	GetAllBook(ctx context.Context) ([]model.Book, error)
	GetBookById(ctx context.Context, id int) (model.Book, error)
	CreateBook(ctx context.Context, bookIn model.Book) (model.Book, error)
	UpdateBook(ctx context.Context, bookIn model.Book) (model.Book, error)
	DeleteBook(ctx context.Context, id int) error
}
