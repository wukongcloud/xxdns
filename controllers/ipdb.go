package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/models"
	"net/http"
	"strconv"
)

// AddIPDB godoc
// @Tags IPDB
// @Summary Create a ipdb
// @Description Create a ipdb
// @Accept  json
// @Produce  json
// @Param IPDBInfo body models.IPDB{} true "填写ip信息"
// @Success 201 {object} models.IPDB
// @Failure 409 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/ipdb [post]
func AddIPDB(c *gin.Context) {
	var data models.IPDB
	_ = c.ShouldBindJSON(&data)
	if err := models.CreateIPDB(&data); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusCreated, data)
}


// GetIPDB godoc
// @Tags IPDB
// @Summary Get all IPDB
// @Description Get all IPDB
// @Accept  json
// @Produce  json
// @Param pagesize query string false "pagesize"
// @Param pagenum  query string false "pagenum"
// @Success 200 {object} []models.IPDB
// @Failure 401 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/ipdb [get]
func GetIPDB(c *gin.Context) {
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
	data := models.GetIPDB(pageSize, pageNum)
	response(c, http.StatusOK, data)
}