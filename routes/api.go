package routes

import (
	"censor.services/app/http/controllers"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func APIRoutes(c RouterContext) {
	api := c.Router.Group("/api/v1/censor")
	{
		articles := api.Group("/censors")
		{
			articles.POST("/censor_text", c.CensorController.CensorText)   // 审核文本
			articles.POST("/censor_image", c.CensorController.CensorImage) // 审核图片
		}
	}
}

type RouterContext struct {
	dig.In
	Router           *gin.Engine
	CensorController *controllers.CensorController
}
