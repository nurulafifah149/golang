package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id          int       `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	Title       string    `json:"title" gorm:"column:title;type:varchar(255)"`
	Description string    `json:"description" gorm:"column:description;type:varchar(255)"`
	UserId      int       `json:"user_id" gorm:"column:user_id;type:integer"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt
}
