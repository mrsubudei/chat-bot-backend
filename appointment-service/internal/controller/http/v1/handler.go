package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/internal/service"
	"github.com/mrsubudei/chat-bot-backend/appointment-service/pkg/logger"
)

func NewRouter(handler *gin.Engine, s service.Service, l logger.Interface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	//swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	//handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	//handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newTranslationRoutes(h, s, l)
	}
}
