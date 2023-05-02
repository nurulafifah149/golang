package model

type BookCreateRequest struct {
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Desc   string `json:"desc" validate:"required"`
}
