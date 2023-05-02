package photo

import (
	"time"

	"github.com/nurulafifah149/golang/module/model/comment"
	"gorm.io/gorm"
)

type Photo struct {
	Id        int               `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	UserId    int               `json:"user_id" gorm:"column:user_id;type:integer"`
	Title     string            `json:"title" gorm:"column:title;type:varchar(255)" validate:"required"`
	Caption   string            `json:"caption" gorm:"column:caption;type:varchar(255)"`
	PhotoUrl  string            `json:"photo_url" gorm:"column:caption;type:varchar(255)" validate:"required"`
	Comments  []comment.Comment `json:"comments" gorm:"foreignKey:PhotoId"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
