package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/nurulafifah149/golang/module/model"
)

type BookRepositoryImplPostgre struct {
	db *sql.DB
}

func NewBookRepositoryPostgre(db *sql.DB) BookRepository {
	return &BookRepositoryImplPostgre{
		db: db,
	}
}

func (b *BookRepositoryImplPostgre) GetAllBook(ctx context.Context) (BookOut []model.Book, err error) {
	//panic("not implemented") // TODO: Implement

	sqlQuery := `SELECT id,title,author,dsc,created_at,updated_at,deleted_at from books WHERE deleted_at is null`
	rows, err := b.db.QueryContext(ctx, sqlQuery)

	if err != nil {
		err = errors.New("ISE")
		return
	}

	defer rows.Close()

	for rows.Next() {
		book := model.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt)

		if err != nil {
			err = errors.New("ISE")
			return
		}

		BookOut = append(BookOut, book)
	}

	return
}

func (b *BookRepositoryImplPostgre) GetBookById(ctx context.Context, id int) (BookOut model.Book, err error) {
	// panic("not implemented") // TODO: Implement
	sqlQuery := `SELECT id,title,author,dsc,created_at,updated_at,deleted_at from books WHERE id = $1 and deleted_at is null`

	stmt, err := b.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		err = errors.New("ISE")
		return
	}

	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		err = errors.New("ISE")
		return
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&BookOut.Id, &BookOut.Title, &BookOut.Author, &BookOut.Desc, &BookOut.CreatedAt, &BookOut.UpdatedAt, &BookOut.DeletedAt)
		if err != nil {
			err = errors.New("ISE")
			return
		}
	}

	if BookOut.Id <= 0 {
		err = errors.New("NF")
	}

	fmt.Println("done all vallidator")
	return
}

// Implement nanti subuh
func (b *BookRepositoryImplPostgre) CreateBook(ctx context.Context, bookIn model.Book) (BookOut model.Book, err error) {
	//panic("not implemented") // TODO: Implement
	queryStr := `INSERT into books(title,author,dsc) VALUES($1,$2,$3) RETURNING id,title,author,dsc;`

	stmt, err := b.db.PrepareContext(ctx, queryStr)
	if err != nil {
		return
	}

	rows, err := stmt.QueryContext(ctx, bookIn.Title, bookIn.Author, bookIn.Desc)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(&BookOut.Id, &BookOut.Title, &BookOut.Author, &BookOut.Desc); err != nil {
			return
		}
	}

	return
}

// Implement nanti subuh
func (b *BookRepositoryImplPostgre) UpdateBook(ctx context.Context, bookIn model.Book) (BookOut model.Book, err error) {
	//panic("not implemented") // TODO: Implement
	queryStr := `UPDATE books SET title=$1,author=$2,dsc=$3,updated_at=now() WHERE id=$4 AND deleted_at is null RETURNING id,title,author,dsc,updated_at,created_at;`
	stmt, err := b.db.PrepareContext(ctx, queryStr)
	if err != nil {
		err = errors.New("ISE")
		return
	}

	rows, err := stmt.QueryContext(ctx, bookIn.Title, bookIn.Author, bookIn.Desc, bookIn.Id)
	if err != nil {
		err = errors.New("ISE")
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err = rows.Scan(&BookOut.Id, &BookOut.Title, &BookOut.Author, &BookOut.Desc, &BookOut.UpdatedAt, &BookOut.CreatedAt); err != nil {
			err = errors.New("ISE")
			return
		}
	}

	return
}

// Implement nanti subuh
func (b *BookRepositoryImplPostgre) DeleteBook(ctx context.Context, id int) (err error) {
	//panic("not implemented") // TODO: Implement
	queryStr := `UPDATE books SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL`
	stmt, err := b.db.PrepareContext(ctx, queryStr)
	if err != nil {
		err = errors.New("ISE")
		return
	}

	res, err := stmt.ExecContext(ctx, id)

	if err != nil {
		err = errors.New("ISE")
		return
	}

	terubah, err := res.RowsAffected()
	if terubah <= 0 {
		err = errors.New("NF")
	}

	if err != nil {
		err = errors.New("NF")
	}

	return
}
