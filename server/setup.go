package server

import (
	"github.com/nurulafifah149/golang/config"
	"github.com/nurulafifah149/golang/module/controller"
	"github.com/nurulafifah149/golang/module/repository"
	"github.com/nurulafifah149/golang/module/service"
)

func initDI() controller.BookController {
	var bookRepo repository.BookRepository

	switch config.Load.DataSource.Mode {
	case config.MODE_GORM:
		pgConn := config.NewPostgresGormConn()
		bookRepo = repository.NewBookRepositoryGorm(pgConn)
	case config.MODE_MAP:
		bookRepo = repository.NewBookRepositoryMap()
	case config.MODE_POSTGRE:
		pgConn := config.NewPostgresConn()
		bookRepo = repository.NewBookRepositoryPostgre(pgConn)
	}

	bookService := service.NewBookService(bookRepo)
	bookController := controller.NewBookController(bookService)

	return bookController
}
