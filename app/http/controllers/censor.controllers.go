package controllers

import (
	"censor.services/app/domain/dtos"
	"censor.services/app/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/henrion-y/base.services/http/errors"
)

type CensorController struct {
	censorService services.CensorService
}

func NewCensorController(censorService services.CensorService) *CensorController {
	return &CensorController{
		censorService: censorService,
	}
}

// @Tags 内容审核
// @Summary 文本审核
// @Description 文本审核
// @Id CensorText
// @Produce  json
// @Param        CensorTextDto  body   dtos.CensorTextDto  true  "文本审核"
// @Success 200 {object} ResponseData
// @Router /censors/censor_text [post]
func (h *CensorController) CensorText(c *gin.Context) {
	req := &dtos.CensorTextDto{}
	if err := c.Bind(req); err != nil {
		responseError(c, errors.New(errors.ErrParamRequired).WithRaw(err))
		return
	}

	data, err := h.censorService.CensorText(c, req)
	if err != nil {
		responseError(c, err)
		return
	}

	responseData(c, data)
}

// @Tags 内容审核
// @Summary 图片审核
// @Description 图片审核
// @Id CensorImage
// @Produce  json
// @Param        CensorImageDto  body   dtos.CensorImageDto  true  "图片审核"
// @Success 200 {object} ResponseData
// @Router /censors/censor_image [post]
func (h *CensorController) CensorImage(c *gin.Context) {
	req := &dtos.CensorImageDto{}
	if err := c.Bind(req); err != nil {
		responseError(c, errors.New(errors.ErrParamRequired).WithRaw(err))
		return
	}

	data, err := h.censorService.CensorImage(c, req)
	if err != nil {
		responseError(c, err)
		return
	}

	responseData(c, data)
}
