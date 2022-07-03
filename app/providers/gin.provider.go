package providers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewGinProvider(config *viper.Viper) *gin.Engine {
	env := config.GetString("app.Env")
	gin.SetMode(env)
	return gin.New()
}
