package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/models"
)

type errResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type viewCreateForm struct {
	Name     string `json:"name" binding:"required"`
	Comment  string `json:"comment" binding:"required"`
	Disabled bool   `json:"disabled"`
}

type viewObject models.View

type viewList []models.View

// response 响应数据
func response(c *gin.Context, httpCode int, v interface{}) {
	c.IndentedJSON(httpCode, v)
}

// responseError 响应错误数据
func responseError(c *gin.Context, httpCode int, errCode int, err string) {
	c.IndentedJSON(
		httpCode,
		errResponse{
			Code:    errCode,
			Message: err,
		},
	)
}
