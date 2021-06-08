package controllers

import (
"github.com/gin-gonic/gin"
"net/http"
)

// GetRecords godoc
// @Tags Record
// @Summary Get all records
// @Description Get all records
// @Accept  json
// @Produce  json
// @Success 200 {object} viewResponse
// @Failure 401 {object} viewResponse
// @Failure 500 {object} viewResponse
// @Router /api/v1/records [get]
func GetRecords(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"records":"aaaaa"})
}