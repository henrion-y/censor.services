package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func GlobalMiddlewares(server *gin.Engine, logMiddleware *LogMiddleware) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.ExposeHeaders = []string{"Content-Length", "Authorization", "Content-Type"}
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	server.Use(
		logMiddleware.Handler(),
		cors.New(corsConfig),
		// limit.MaxAllowed(config.Connection),
		gzip.Gzip(gzip.DefaultCompression),
	)
}
