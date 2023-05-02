package user

import (
	"context"
	"errors"

	"github.com/nurulafifah149/golang/module/model/user"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, userIn user.User) (err error) {
	// panic("not implemented") // TODO: Implement
	tx := u.db.Model(&user.User{}).Create(&userIn)

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (u *UserRepositoryImpl) GetByUsername(ctx context.Context, username string) (userOut user.User, err error) {
	// panic("not implemented") // TODO: Implement
	tx := u.db.Model(&user.User{}).Where("username = ?", username).Find(&userOut)

	if err = tx.Error; err != nil {
		return
	}

	if userOut.Id <= 0 {
		err = errors.New("NF")
		return
	}

	return
}
