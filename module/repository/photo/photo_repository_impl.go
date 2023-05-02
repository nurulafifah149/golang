package photo

import (
	"context"
	"errors"

	"github.com/nurulafifah149/golang/module/model/photo"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepositoryImpl struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{
		db: db,
	}
}

func (p *PhotoRepositoryImpl) GetAll(ctx context.Context) (photoOut []photo.Photo, err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&photo.Photo{}).Find(&photoOut).Order("updated_at ASC")

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (p *PhotoRepositoryImpl) GetById(ctx context.Context, idPhoto int) (photoOut photo.Photo, err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&photo.Photo{}).Preload("Comments").Where("id = ?", idPhoto).Find(&photoOut)

	if err = tx.Error; err != nil {
		return
	}

	if photoOut.Id <= 0 {
		err = errors.New("NF")
		return
	}

	return
}

func (p *PhotoRepositoryImpl) Create(ctx context.Context, photoIn photo.Photo) (photo.Photo, error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&photoIn).Clauses(clause.Returning{}).Create(&photoIn)

	if err := tx.Error; err != nil {
		return photoIn, err
	}

	return photoIn, nil
}

func (p *PhotoRepositoryImpl) Update(ctx context.Context, photoIn photo.Photo) (photoOut photo.Photo, err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&photoOut).Clauses(clause.Returning{}).Where("id = ?", photoIn.Id).Updates(&photoIn)

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (p *PhotoRepositoryImpl) Delete(ctx context.Context, idPhoto int) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&photo.Photo{}).Where("id = ?", idPhoto).Delete(&photo.Photo{})

	if err = tx.Error; err != nil {
		return
	}

	return
}
