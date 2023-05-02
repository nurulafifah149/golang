package socialmedia

import (
	"time"

	"gorm.io/gorm"
)

type Socialmedia struct {
	Id             int       `json:"id" gorm:"column:id;type:integer;primaryKey;autoIncrement;not null"`
	UserId         int       `json:"user_id" gorm:"column:user_id;type:integer;not null"`
	Name           string    `json:"name" gorm:"column:name;type:varchar(255)" validate:"required"`
	SocialMediaUrl string    `json:"social_media_url" gorm:"column:social_media_url;type:varchar(255)" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      gorm.DeletedAt
}
