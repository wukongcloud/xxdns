package controllers

import "github.com/gin-gonic/gin"

type errResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

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

type viewCreateForm struct {
	Name     string `json:"name" binding:"required"`
	Comment  string `json:"comment"`
	Disabled bool   `json:"disabled"`
}

type recordCreateForm struct {
	Domain  string `json:"domain" binding:"required"`
	View    string `json:"view" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Type    string `json:"type" binding:"required"`
	Content string `json:"content" binding:"required"`
	TTL     uint   `json:"ttl" binding:"required"`
	Priority uint `json:"priority"`
	Disabled bool   `json:"disabled"`
	Comment  string `json:"comment"`
}

type aclCreateForm struct {
	Name     string `json:"name" binding:"required"`
	Path     string `json:"path" binding:"required"`
	Comment  string `json:"comment"`
	Disabled bool   `json:"disabled"`
}
