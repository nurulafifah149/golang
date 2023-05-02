package repository

import "github.com/nurulafifah149/golang/module/model"

type BookRepository interface {
	GetAllBook() ([]model.Book, error)
	GetBookById(id int) (model.Book, error)
	CreateBook(model.Book) (model.Book, error)
	UpdateBook(model.Book) (model.Book, error)
	DeleteBook(id int) error
}
