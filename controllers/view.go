package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/models"
	"net/http"
	"strconv"
)

// GetViews godoc
// @Tags View
// @Summary Get all views
// @Description Get all views
// @Accept  json
// @Produce  json
// @Param pagesize query string false "pagesize"
// @Param pagenum  query string false "pagenum"
// @Success 200 {object} viewList
// @Failure 401 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/views [get]
func GetViews(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data := models.GetViews(pageSize, pageNum)
	c.JSON(http.StatusOK, data)
}

// AddView godoc
// @Tags View
// @Summary Create a views
// @Description Create a views
// @Accept  json
// @Produce  json
// @Param viewInfo body viewCreateForm{name=string,comment=string,disabled=bool} true "填写视图信息"
// @Success 201 {object} viewObject
// @Failure 409 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/views [post]
func AddView(c *gin.Context) {
	var data models.View
	_ = c.ShouldBindJSON(&data)
	if data.CheckViewExist(data.Name) == false {
		models.CreateView(&data)
		response(c,http.StatusCreated,data)
	} else {
		responseError(c,http.StatusConflict,499,"view name already exist")
	}
}
