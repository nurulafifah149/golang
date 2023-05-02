package repository

import (
	"context"
	"errors"

	"github.com/nurulafifah149/golang/module/model"
)

type BookRepositoryImplMap struct {
	Data map[uint64]model.Book
}

func NewBookRepositoryMap() BookRepository {
	DataStore := map[uint64]model.Book{}
	return &BookRepositoryImplMap{
		Data: DataStore,
	}
}

// b BookRepositoryImpl
func (b *BookRepositoryImplMap) GetAllBook(ctx context.Context) ([]model.Book, error) {
	var data []model.Book
	for _, vals := range b.Data {
		if vals.Deleted != true {
			data = append(data, vals)
		}
	}
	return data, nil
}

func (b *BookRepositoryImplMap) GetBookById(ctx context.Context, id int) (model.Book, error) {
	var data model.Book
	valid := false

	for _, vals := range b.Data {
		if id == int(vals.Id) && vals.Deleted != true {
			data = model.Book{
				Id:     vals.Id,
				Title:  vals.Title,
				Author: vals.Author,
				Desc:   vals.Desc,
			}
			valid = true
			break
		}
	}

	if valid {
		return data, nil
	} else {
		err := errors.New("NF")
		return data, err
	}
}

func (b *BookRepositoryImplMap) CreateBook(ctx context.Context, book model.Book) (model.Book, error) {
	counter := len(b.Data)
	book.Id = counter
	book.Deleted = false
	b.Data[uint64(counter)] = book
	return b.Data[uint64(counter)], nil
}

func (b *BookRepositoryImplMap) UpdateBook(ctx context.Context, book model.Book) (bookOut model.Book, err error) {
	if b.Data[uint64(book.Id)].Deleted {
		err = errors.New("NF")
		return
	} else {
		b.Data[uint64(book.Id)] = book
		bookOut = b.Data[uint64(book.Id)]
		return
	}
}

func (b *BookRepositoryImplMap) DeleteBook(ctx context.Context, id int) error {

	if id <= len(b.Data)-1 && b.Data[uint64(id)].Deleted == false {
		book := b.Data[uint64(id)]
		book.Deleted = true
		b.Data[uint64(id)] = book
		return nil
	} else {
		return errors.New("NF")
	}
}
