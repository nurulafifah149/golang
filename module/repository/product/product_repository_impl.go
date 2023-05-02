package product

import (
	"context"
	"errors"

	"github.com/nurulafifah149/golang/module/model/product"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

func (p *ProductRepositoryImpl) GetAll(ctx context.Context) (productOut []product.Product, err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&product.Product{}).Find(&productOut).Order("updated_at ASC")

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (p *ProductRepositoryImpl) GetAllByUserId(ctx context.Context, userId int) (productOut []product.Product, err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&product.Product{}).Where("user_id = ?", userId).Find(&productOut).Order("updated_at ASC")

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (p *ProductRepositoryImpl) GetById(ctx context.Context, idProduct int) (productOut product.Product, err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&product.Product{}).Where("id = ?", idProduct).Find(&productOut)

	if err = tx.Error; err != nil {
		return
	}

	if productOut.Id <= 0 {
		err = errors.New("Data is Not Found")
		return
	}

	return
}

func (p *ProductRepositoryImpl) Create(ctx context.Context, ProductIn product.Product) (product.Product, error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&ProductIn).Clauses(clause.Returning{}).Create(&ProductIn)

	if err := tx.Error; err != nil {
		return ProductIn, err
	}

	return ProductIn, nil
}

func (p *ProductRepositoryImpl) Update(ctx context.Context, ProductIn product.Product) (productOut product.Product, err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&productOut).Clauses(clause.Returning{}).Where("id = ?", ProductIn.Id).Updates(&ProductIn)

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (p *ProductRepositoryImpl) Delete(ctx context.Context, idProduct int) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := p.db.Model(&product.Product{}).Where("id = ?", idProduct).Delete(&product.Product{})

	if err = tx.Error; err != nil {
		return
	}

	return
}
