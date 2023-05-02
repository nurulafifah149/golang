package user

import (
	"time"

	"github.com/nurulafifah149/golang/module/model/product"
	"gorm.io/gorm"
)

type User struct {
	Id        int               `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	Username  string            `json:"username" gorm:"column:username;type:varchar(255);unique;not null" validate:"required"`
	Email     string            `json:"email" gorm:"column:email;type:varchar(255);unique;not null" validate:"required,email"`
	Password  string            `json:"password" gorm:"column:password;type:varchar(60);not null" validate:"required"`
	Age       int               `json:"age" gorm:"column:age;type:integer;not null" validate:"required,gt=8"`
	Role      string            `json:"role" gorm:"column:role;type:varchar(255);unique;not null" validate:"required"`
	Products  []product.Product `json:"products"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"-"`
}
