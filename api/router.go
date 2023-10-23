package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/casbin/casbin/v2"
	_ "gitlab.com/tizim-back/api/docs"
	v1 "gitlab.com/tizim-back/api/handlers/v1"
	"gitlab.com/tizim-back/config"
	"gitlab.com/tizim-back/pkg/logger"
	"gitlab.com/tizim-back/storage"

	jwthandler "gitlab.com/tizim-back/api/tokens"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RoutetOptions struct {
	Cfg            *config.Config
	Storage        storage.StorageI
	Log            logger.Logger
	CasbinEnforcer *casbin.Enforcer
}

// New ...
// @Description Created by Otajonov Quvonchbek
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option RoutetOptions) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowHeaders = append(corConfig.AllowHeaders, "*")
	router.Use(cors.New(corConfig))

	jwt := jwthandler.JWTHandler{
		SigninKey: option.Cfg.SignKey,
		Log:       option.Log,
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "App is running...",
		})
	})

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:        option.Log,
		Cfg:        option.Cfg,
		Storage:    &option.Storage,
		JWTHandler: jwt,
	})

	//router.Use(middleware.NewAuth(option.CasbinEnforcer, jwt, config.Load(".")))

	router.Static("/media", "./media")
	api := router.Group("/v1")

	// User
	api.POST("/auth/login", handlerV1.Login)

	// Workers
	api.POST("/worker", handlerV1.CreateWorker)
	api.DELETE("/worker/:id", handlerV1.DeleteWorker)
	api.POST("/worker/update", handlerV1.UpdateWorker)
	api.GET("/workers", handlerV1.GetAllWorkers)
	api.GET("/worker/:id", handlerV1.GetWorker)
	api.GET("/workers/at-work", handlerV1.GetWorkersAtWork)
	api.GET("/workers/:gender", handlerV1.GetWorkersByGender)

	// Workers-history
	api.GET("get-workers-by-month/:date", handlerV1.GetAllWorkersByMonth)
	//api.GET("get-workers-by-month/:date", handlerV1.GetAllWorkersByMonth)

	// File - upload
	api.POST("/file-upload", handlerV1.UploadFile)

	// Daily
	api.POST("/daily", handlerV1.CreateAttendance)
	api.GET("/daily/portion", handlerV1.GetAttendancePortion)


	
	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
