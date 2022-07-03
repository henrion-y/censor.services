package providers

import (
	"context"
	"log"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func StartService(lifecycle fx.Lifecycle, web *gin.Engine, config *viper.Viper) {
	listenHTTP := config.GetString("app.ListenHTTP")
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("lifecycle.OnStart")
			return web.Run(listenHTTP)
		},
	})
}
