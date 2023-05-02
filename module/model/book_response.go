package model

import "time"

type BookResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
