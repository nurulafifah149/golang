package user

import (
	"time"

	"github.com/nurulafifah149/golang/module/model/comment"
	"github.com/nurulafifah149/golang/module/model/photo"
	"github.com/nurulafifah149/golang/module/model/socialmedia"
	"gorm.io/gorm"
)

type User struct {
	Id          int                       `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	Username    string                    `json:"username" gorm:"column:username;type:varchar(255);unique;not null" validate:"required"`
	Email       string                    `json:"email" gorm:"column:email;type:varchar(255);unique;not null" validate:"required,email"`
	Password    string                    `json:"password" gorm:"column:password;type:varchar(60);not null" validate:"required"`
	Age         int                       `json:"age" gorm:"column:age;type:integer;not null" validate:"required,gt=8"`
	SocialMedia []socialmedia.Socialmedia `gorm:"foreignKey:UserId"`
	Photo       []photo.Photo             `gorm:"foreignKey:UserId"`
	Comment     []comment.Comment         `gorm:"foreignKey:UserId"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedAt   time.Time                 `json:"updated_at"`
	DeletedAt   gorm.DeletedAt            `json:"-"`
}
