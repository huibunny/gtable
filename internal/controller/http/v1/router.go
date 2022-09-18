// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	"gtable/config"
	_ "gtable/docs"
	"gtable/internal/usecase"
	"gtable/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       gtable API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8820
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Table, cfg config.Config) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe & consul
	handler.GET(cfg.Consul.CheckApi, func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newtableRoutes(h, t, l, cfg.PG.Tables)
	}
}
