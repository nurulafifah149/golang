package server

import (
	"fmt"

	"github.com/nurulafifah149/golang/docs"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/gin-gonic/gin"
	"github.com/nurulafifah149/golang/config"

	modelComment "github.com/nurulafifah149/golang/module/model/comment"
	modelPhoto "github.com/nurulafifah149/golang/module/model/photo"
	modelSocmed "github.com/nurulafifah149/golang/module/model/socialmedia"
	modelUser "github.com/nurulafifah149/golang/module/model/user"

	"github.com/go-playground/validator/v10"

	repositoryComment "github.com/nurulafifah149/golang/module/repository/comment"
	repositoryPhoto "github.com/nurulafifah149/golang/module/repository/photo"
	repositorySocmed "github.com/nurulafifah149/golang/module/repository/socialmedia"
	repositoryUser "github.com/nurulafifah149/golang/module/repository/user"

	serviceComment "github.com/nurulafifah149/golang/module/service/comment"
	servicePhoto "github.com/nurulafifah149/golang/module/service/photo"
	serviceSocmed "github.com/nurulafifah149/golang/module/service/socialmedia"
	serviceUser "github.com/nurulafifah149/golang/module/service/user"

	handlerComment "github.com/nurulafifah149/golang/module/handle/comment"
	handlerPhoto "github.com/nurulafifah149/golang/module/handle/photo"
	handlerSocmed "github.com/nurulafifah149/golang/module/handle/socialmedia"
	handlerUser "github.com/nurulafifah149/golang/module/handle/user"

	routerComment "github.com/nurulafifah149/golang/module/router/v1/comment"
	routerPhoto "github.com/nurulafifah149/golang/module/router/v1/photo"
	routerSocmed "github.com/nurulafifah149/golang/module/router/v1/socialmedia"
	routerUser "github.com/nurulafifah149/golang/module/router/v1/user"
)

func Serve() {
	var validate *validator.Validate
	//Load Model
	pgConn := config.NewPostgresGormConn()
	MPhoto := modelPhoto.Photo{}
	MSocial := modelSocmed.Socialmedia{}
	MUser := modelUser.User{}
	MComment := modelComment.Comment{}

	if config.Load.DataSource.Migrate {
		pgConn.AutoMigrate(&MPhoto, &MSocial, &MUser, &MComment)
	}
	//bookService := service.NewBookService(bookRepo)
	//bookController := controller.NewBookController(bookService)

	//load Repo
	RPhoto := repositoryPhoto.NewPhotoRepository(pgConn)
	RSocial := repositorySocmed.NewSocialmediaRepository(pgConn)
	RUser := repositoryUser.NewUserRepository(pgConn)
	RComment := repositoryComment.NewCommentRepository(pgConn)

	//load Services
	SPhoto := servicePhoto.NewPhotoService(RPhoto, validate)
	SComment := serviceComment.NewCommentService(RComment, validate)
	SUser := serviceUser.NewUserService(RUser, validate)
	SSocial := serviceSocmed.NewSocialmediaService(RSocial, validate)

	//load Handler
	hUser := handlerUser.NewUserHandler(SUser)
	hSocmed := handlerSocmed.NewSocialmediaHandler(SSocial)
	hPhoto := handlerPhoto.NewPhotoHandler(SPhoto)
	hComment := handlerComment.NewCommenHandler(SComment)

	ginServer := gin.Default()

	if config.Load.Server.Env == config.ENV_PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	//###init middleware
	ginServer.Use(
		gin.Logger(),   // untuk log request yang masuk
		gin.Recovery(), // untuk auto restart kalau panic
	)
	//###swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginServer.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ###mendaftarkan route

	v1 := ginServer.Group("/api/v1")

	routerUser.UserRouter(v1, hUser)
	routerSocmed.SocialmediaRoute(v1, hSocmed)
	routerPhoto.PhotoRoute(v1, hPhoto)
	routerComment.CommentRoute(v1, hComment)

	///run server
	ginServer.Run(fmt.Sprintf(":%v", config.Load.Server.Http.Port))
}
