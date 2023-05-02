package repository

import (
	"errors"

	"github.com/nurulafifah149/golang/module/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImplGorm struct {
	db *gorm.DB
}

func NewBookRepositoryGorm(db *gorm.DB) BookRepository {
	return &UserRepositoryImplGorm{
		db: db,
	}
}

func (b *UserRepositoryImplGorm) GetAllBook() (books []model.Book, err error) {

	tx := b.db.Model(&model.Book{}).Find(&books).Order("created_at ASC")

	if err = tx.Error; err != nil {
		err = errors.New("ISE")
		return
	}

	return
}

func (b *UserRepositoryImplGorm) GetBookById(id int) (book model.Book, err error) {
	//panic("not implemented") // TODO: Implement
	tx := b.db.Table("books").Model(&model.Book{}).Where("id = ?", id).Find(&book)

	if err = tx.Error; err != nil {
		err = errors.New("ISE")
		return
	}

	if book.Id <= 0 {
		err = errors.New("NF")
	}

	return
}

func (b *UserRepositoryImplGorm) CreateBook(book model.Book) (model.Book, error) {
	result := b.db.Table("books").Create(&book)
	if result.Error != nil {
		return book, errors.New("ISE")
	}

	return book, nil
}

func (b *UserRepositoryImplGorm) UpdateBook(bookin model.Book) (bookOut model.Book, err error) {
	tx := b.db.Table("books").Model(&bookOut).Clauses(clause.Returning{}).Where("id = ?", bookin.Id).Updates(&bookin)
	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("user is not found")
		return
	}

	return
}

func (b *UserRepositoryImplGorm) DeleteBook(id int) (err error) {
	//panic("not implemented") // TODO: Implement
	tx := b.db.Table("books").Model(&model.Book{}).Where("id = ?", id).Delete(&model.Book{})
	if err = tx.Error; err != nil {
		err = errors.New("ISE")
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("NF")
		return
	}
	return

}
