package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// GetAcls godoc
// @Tags Acl
// @Summary Get all ACLs
// @Description Get all ACLs
// @Accept  json
// @Produce  json
// @Param pagesize query string false "pagesize"
// @Param pagenum  query string false "pagenum"
// @Success 200 {object} []models.Acl
// @Failure 401 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/acls [get]
func GetAcls(c *gin.Context) {
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
	if err,data := models.GetAcls(pageSize, pageNum);err !=nil{
		responseError(c,http.StatusInternalServerError,5000,err.Error())
		return
	}else {
		response(c, http.StatusOK, data)
	}
}

// GetAclById godoc
// @Tags Acl
// @Summary Get acl by id
// @Description Get acl by id
// @Accept  json
// @Produce  json
// @Param id path int true "acl id"
// @Success 200 {object} models.Acl
// @Failure 404 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/acls/{id} [get]
func GetAclById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if data, err := models.GetAclById(id); err != nil {
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

// AddAcl godoc
// @Tags Acl
// @Summary Create a acl
// @Description Create a acl
// @Accept  json
// @Produce  json
// @Param AclInfo body aclCreateForm{name=string,comment=string,disabled=bool} true "填写ACL信息"
// @Success 201 {object} models.Acl
// @Failure 409 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/acls [post]
func AddAcl(c *gin.Context) {
	var data models.Acl
	_ = c.ShouldBindJSON(&data)
	if data.CheckAclExist(data.Name) == false {
		if err := models.CreateAcl(&data); err != nil {
			responseError(c, http.StatusInternalServerError, 500, err.Error())
			return
		}
		response(c, http.StatusCreated, data)
	} else {
		responseError(c, http.StatusConflict, 409, "acl name already exist")
	}
}

// DeleteAcl godoc
// @Tags Acl
// @Summary Get acl by id
// @Description Get acl by id
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errResponse "{error:{code:400,message:"bad request"}}"
// @Failure 404 {object} errResponse  "{error:{code:404,message:"record not found"}}"
// @Failure 500 {object} errResponse  "{error:{code:404,message:"server error"}}"
// @Router /api/v1/acls/{id} [delete]
func DeleteAcl(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	if err := models.DeleteAcl(id); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusOK, "ok")
}

// UpdateAclById godoc
// @Tags Acl
// @Summary Create a acl
// @Description Create a acl
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param aclInfo body aclCreateForm{name=string,comment=string,disabled=bool} true "填写Acl信息"
// @Success 200 {object} models.Acl
// @Failure 500 {object} errResponse
// @Router /api/v1/acls/{id} [put]
func UpdateAclById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	var data models.Acl
	if err = c.ShouldBindJSON(&data); err != nil {
		responseError(c, http.StatusBadRequest, 400, err.Error())
		return
	}
	data.ID = id

	if _, err := models.GetAclById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			responseError(c, http.StatusNotFound, 404, err.Error())
			return
		}
	}
	if err := models.UpdateAclById(id, &data); err != nil {
		responseError(c, http.StatusInternalServerError, 500, err.Error())
		return
	}
	response(c, http.StatusOK, data)
}
