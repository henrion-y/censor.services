package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/henrion-y/base.services/http/errors"
)

func responseError(c *gin.Context, err error) {
	if xerr, ok := err.(*errors.XError); ok {
		fmt.Printf("%+v\n", xerr)
		c.AbortWithStatusJSON(xerr.Status, xerr)
	} else {
		responseError(c, errors.New(errors.ErrRuntime))
	}
}

type ResponseData struct {
	Code int32       `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

func responseData(c *gin.Context, data interface{}) {
	b, _ := json.Marshal(data)
	fmt.Printf("%s %s response: %s\n", strings.ToUpper(c.Request.Method), c.Request.URL.Path, string(b))
	c.JSON(http.StatusOK, ResponseData{Code: 0, Data: data})
}

func responseSuccess(c *gin.Context) {
	fmt.Printf("%s %s response success\n", strings.ToUpper(c.Request.Method), c.Request.URL.Path)
	c.JSON(http.StatusOK, ResponseData{Code: 0})
}
