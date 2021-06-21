package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/server/models"
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
// @Param country query string false "country"
// @Param province query string false "province"
// @Param isp query string false "isp"
// @Param pagesize query string false "pagesize"
// @Param pagenum  query string false "pagenum"
// @Success 200 {object} []models.IPDB
// @Failure 401 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/ipdb [get]
func GetIPDB(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	country := c.Query("country")
	province := c.Query("province")
	isp := c.Query("isp")
	switch {
	case pageSize >= 200000:
		pageSize = 200000
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data := models.GetIPDB(pageSize, pageNum, country, province, isp)
	response(c, http.StatusOK, data)
}
