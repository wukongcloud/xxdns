package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetRecords godoc
// @Tags Record
// @Summary Get all records
// @Description Get all records
// @Accept  json
// @Produce  json
// @Param pagesize query string false "pagesize"
// @Param pagenum  query string false "pagenum"
// @Success 200 {object} []models.Record
// @Failure 401 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/records [get]
func GetRecords(c *gin.Context) {
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
	data := models.GetRecords(pageSize, pageNum)
	response(c, http.StatusOK, data)
}

// GetRecordById godoc
// @Tags Record
// @Summary Get record by id
// @Description Get record by id
// @Accept  json
// @Produce  json
// @Param id path int true "record id"
// @Success 200 {object} models.Record
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/records/{id} [get]
func GetRecordById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if data, err := models.GetRecordById(id); err != nil {
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

// AddRecord godoc
// @Tags Record
// @Summary Create a record
// @Description Create a record
// @Accept  json
// @Produce  json
// @Param RecordInfo body recordCreateForm{domain=string,name=string,type=string,content=string,comment=string,disabled=bool} true "填写视图信息"
// @Success 201 {object} models.Record
// @Failure 409 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/records [post]
func AddRecord(c *gin.Context) {
	var data models.Record
	_ = c.ShouldBindJSON(&data)
	if data.CheckRecordExist(data.Name) == false {
		if err := models.CreateRecord(&data); err != nil {
			responseError(c, http.StatusInternalServerError, 500, err.Error())
			return
		}
		response(c, http.StatusCreated, data)
	} else {
		responseError(c, http.StatusConflict, 409, "record name already exist")
	}
}

// DeleteRecord godoc
// @Tags Record
// @Summary Get record by id
// @Description Get record by id
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errResponse "{error:{code:400,message:"bad request"}}"
// @Failure 404 {object} errResponse  "{error:{code:404,message:"record not found"}}"
// @Failure 500 {object} errResponse  "{error:{code:404,message:"server error"}}"
// @Router /api/v1/records/{id} [delete]
func DeleteRecord(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if err := models.DeleteRecord(id); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusOK, "ok")
}

// UpdateRecordById godoc
// @Tags Record
// @Summary Create a record
// @Description Create a record
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param RecordInfo body recordCreateForm{domain=string,name=string,type=string,content=string,comment=string,disabled=bool} true "填写视图信息"
// @Success 200 {object} models.Record
// @Failure 500 {object} errResponse
// @Router /api/v1/records/{id} [put]
func UpdateRecordById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	var data models.Record
	if err = c.ShouldBindJSON(&data); err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	data.ID = id

	if _, err := models.GetRecordById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(c, http.StatusNotFound, 404, err.Error())
			return
		}
	}
	if err := models.UpdateRecordById(id, &data); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusOK, data)
}
