package comment

import (
	"context"
	"errors"

	"github.com/nurulafifah149/golang/module/model/comment"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{
		db: db,
	}
}

func (c *CommentRepositoryImpl) GetAll(ctx context.Context) (comOut []comment.Comment, err error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(&comment.Comment{}).Find(&comOut).Order("updated_at ASC")

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (c *CommentRepositoryImpl) GetById(ctx context.Context, idComment int) (comOut comment.Comment, err error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(&comment.Comment{}).Where("id = ?", idComment).Find(&comOut)

	if err = tx.Error; err != nil {
		return
	}

	if comOut.Id <= 0 {
		err = errors.New("NF")
		return
	}

	return
}

func (c *CommentRepositoryImpl) Create(ctx context.Context, comIn comment.Comment) (comment.Comment, error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(&comIn).Clauses(clause.Returning{}).Create(&comIn)

	if err := tx.Error; err != nil {
		return comIn, err
	}

	return comIn, nil
}

func (c *CommentRepositoryImpl) Update(ctx context.Context, comIn comment.Comment) (comOut comment.Comment, err error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(&comOut).Clauses(clause.Returning{}).Where("id = ?", comIn.Id).Updates(&comIn)

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (c *CommentRepositoryImpl) Delete(ctx context.Context, idComment int) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := c.db.Model(&comment.Comment{}).Where("id = ?", idComment).Delete(&comment.Comment{})

	if err = tx.Error; err != nil {
		return
	}

	return
}
