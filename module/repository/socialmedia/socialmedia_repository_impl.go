package socialmedia

import (
	"context"
	"errors"

	"github.com/nurulafifah149/golang/module/model/socialmedia"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SocialmediaRepositoryImpl struct {
	db *gorm.DB
}

func NewSocialmediaRepository(db *gorm.DB) SocialmediaRepository {
	return &SocialmediaRepositoryImpl{
		db: db,
	}
}

func (s *SocialmediaRepositoryImpl) GetAll(ctx context.Context) (socOut []socialmedia.Socialmedia, err error) {
	//	panic("not implemented") // TODO: Implement
	tx := s.db.Model(&socialmedia.Socialmedia{}).Find(&socOut).Order("update_at ASC")

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (s *SocialmediaRepositoryImpl) GetById(ctx context.Context, idSoc int) (socOut socialmedia.Socialmedia, err error) {
	//panic("not implemented") // TODO: Implement
	tx := s.db.Model(&socialmedia.Socialmedia{}).Where("id = ?", idSoc).Find(&socOut)

	if err = tx.Error; err != nil {
		return
	}

	if socOut.Id <= 0 {
		err = errors.New("NF")
		return
	}

	return
}

func (s *SocialmediaRepositoryImpl) Create(ctx context.Context, socIn socialmedia.Socialmedia) (socialmedia.Socialmedia, error) {
	// panic("not implemented") // TODO: Implement
	tx := s.db.Model(&socIn).Clauses(clause.Returning{}).Create(&socIn)

	if err := tx.Error; err != nil {
		return socIn, err
	}

	return socIn, nil
}

func (s *SocialmediaRepositoryImpl) Update(ctx context.Context, socIn socialmedia.Socialmedia) (socOut socialmedia.Socialmedia, err error) {
	tx := s.db.Model(&socOut).Clauses(clause.Returning{}).Where("id = ?", socIn.Id).Updates(&socIn)

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (s *SocialmediaRepositoryImpl) Delete(ctx context.Context, idSoc int) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := s.db.Model(&socialmedia.Socialmedia{}).Where("id = ?", idSoc).Delete(&socialmedia.Socialmedia{})

	if err = tx.Error; err != nil {
		return
	}

	return
}
