package helper

import "github.com/nurulafifah149/golang/module/model"

func BookDomainToResp(data model.Book) model.BookResponse {
	resp := model.BookResponse{
		Id:        data.Id,
		Title:     data.Title,
		Author:    data.Author,
		Desc:      data.Desc,
		UpdatedAt: data.UpdatedAt,
		CreatedAt: data.CreatedAt,
	}
	return resp
}
