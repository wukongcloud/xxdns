package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/models"
	"gorm.io/gorm"
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
// @Success 200 {object} []models.View
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
	response(c, http.StatusOK, data)
}

// GetViewById godoc
// @Tags View
// @Summary Get view by id
// @Description Get view by id
// @Accept  json
// @Produce  json
// @Param id path int true "view id"
// @Success 200 {object} models.View
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/views/{id} [get]
func GetViewById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if data, err := models.GetViewById(id); err != nil {
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

// AddView godoc
// @Tags View
// @Summary Create a view
// @Description Create a view
// @Accept  json
// @Produce  json
// @Param viewInfo body viewCreateForm{name=string,comment=string,disabled=bool} true "填写视图信息"
// @Success 201 {object} models.View
// @Failure 409 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/views [post]
func AddView(c *gin.Context) {
	var data models.View
	_ = c.ShouldBindJSON(&data)
	if data.CheckViewExist(data.Name) == false {
		if err := models.CreateView(&data); err != nil {
			responseError(c, http.StatusInternalServerError, 500, err.Error())
			return
		}
		response(c, http.StatusCreated, data)
	} else {
		responseError(c, http.StatusConflict, 409, "view name already exist")
	}
}

// DeleteView godoc
// @Tags View
// @Summary Get view by id
// @Description Get view by id
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errResponse "{error:{code:400,message:"bad request"}}"
// @Failure 404 {object} errResponse  "{error:{code:404,message:"record not found"}}"
// @Failure 500 {object} errResponse  "{error:{code:404,message:"server error"}}"
// @Router /api/v1/views/{id} [delete]
func DeleteView(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if err := models.DeleteView(id); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusOK, "ok")
}

// UpdateViewById godoc
// @Tags View
// @Summary Create a view
// @Description Create a view
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param viewInfo body viewCreateForm{name=string,comment=string,disabled=bool} true "填写视图信息"
// @Success 200 {object} models.View
// @Failure 500 {object} errResponse
// @Router /api/v1/views/{id} [put]
func UpdateViewById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	var data models.View
	if err = c.ShouldBindJSON(&data); err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	data.ID = id

	if _, err := models.GetViewById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(c, http.StatusNotFound, 404, err.Error())
			return
		}
	}
	if err := models.UpdateViewById(id, &data); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusOK, data)
}
