package http

import (
	"go-log-saas/internal/adapter/config"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type Router struct {
	*gin.Engine
}

func NewRouter(config *config.Cfg, handler Handler, logger *zap.SugaredLogger) *Router {
	if config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.HTTP.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(gin.Recovery(), cors.New(ginConfig))

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		logs := v1.Group("/logs")
		{

			logsHandler := logs.Group("/")
			{
				logsHandler.POST("/", handler.IngestLog)
				logsHandler.GET("/", handler.SearchLog)
				logsHandler.GET("/:id", handler.SearchLogById)
			}
		}
	}

	return &Router{
		router,
	}
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
