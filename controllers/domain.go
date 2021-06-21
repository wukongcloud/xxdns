package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetDomains godoc
// @Tags Domain
// @Summary Get all domains
// @Description Get all domains
// @Accept  json
// @Produce  json
// @Param pagesize query string false "pagesize"
// @Param pagenum  query string false "pagenum"
// @Success 200 {object} []models.Domain
// @Failure 401 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/domains [get]
func GetDomains(c *gin.Context) {
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
	data := models.GetDomains(pageSize, pageNum)
	response(c, http.StatusOK, data)
}

// GetDomainById godoc
// @Tags Domain
// @Summary Get domain by id
// @Description Get domain by id
// @Accept  json
// @Produce  json
// @Param id path int true "domain id"
// @Success 200 {object} models.Domain
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/domains/{id} [get]
func GetDomainById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if data, err := models.GetDomainById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(c, http.StatusNotFound, 404, err.Error())
			return
		}
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	} else {
		response(c, http.StatusOK, data)
	}
}

// AddDomain godoc
// @Tags Domain
// @Summary Create a domain
// @Description Create a domain
// @Accept  json
// @Produce  json
// @Param domainInfo body domainCreateForm{name=string,comment=string,disabled=bool} true "填写域名信息"
// @Success 201 {object} models.Domain
// @Failure 409 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/domains [post]
func AddDomain(c *gin.Context) {
	var data models.Domain
	_ = c.ShouldBindJSON(&data)
	if data.CheckDomainExist(data.Name) == false {
		if err := models.CreateDomain(&data); err != nil {
			responseError(c, http.StatusInternalServerError, 500, err.Error())
			return
		}
		response(c, http.StatusCreated, data)
	} else {
		responseError(c, http.StatusConflict, 409, "domain name already exist")
	}
}

// DeleteDomain godoc
// @Tags Domain
// @Summary Get domain by id
// @Description Get domain by id
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errResponse "{error:{code:400,message:"bad request"}}"
// @Failure 404 {object} errResponse  "{error:{code:404,message:"record not found"}}"
// @Failure 500 {object} errResponse  "{error:{code:404,message:"server error"}}"
// @Router /api/v1/domains/{id} [delete]
func DeleteDomain(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if err := models.DeleteDomain(id); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusOK, "ok")
}

// UpdateDomainById godoc
// @Tags Domain
// @Summary Create a domain
// @Description Create a domain
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param domainInfo body domainCreateForm{name=string,comment=string,disabled=bool} true "填写域名信息"
// @Success 200 {object} models.Domain
// @Failure 500 {object} errResponse
// @Router /api/v1/domains/{id} [put]
func UpdateDomainById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	var data models.Domain
	if err = c.ShouldBindJSON(&data); err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	data.ID = uint(id)

	if _, err := models.GetDomainById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(c, http.StatusNotFound, 404, err.Error())
			return
		}
	}
	if err := models.UpdateDomainById(id, &data); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusOK, data)
}
