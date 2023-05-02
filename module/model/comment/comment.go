package comment

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id        int       `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	UserId    int       `json:"user_id" gorm:"column:user_id;type:integer"`
	PhotoId   int       `json:"photo_id" gorm:"column:photo_id;type:integer;not null"`
	Message   string    `json:"message" gorm:"column:message;type:varchar(255)" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
