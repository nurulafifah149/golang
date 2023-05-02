package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id        int       `json:"id" gorm:"column:id;type:integer;primarykey;autoIncrement;not null"`
	Title     string    `json:"title" gorm:"column:title;not null"`
	Author    string    `json:"author" gorm:"column:author;not null"`
	Desc      string    `json:"desc" gorm:"column:dsc;not null"`
	Deleted   bool      `json:"-" gorm:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}

func (Book) TableName() string {
	return "books"
}
