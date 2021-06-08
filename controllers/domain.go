package controllers

import (
"github.com/gin-gonic/gin"
"net/http"
)


// GetDomains godoc
// @Tags Domain
// @Summary Get all domains
// @Description Get all domains
// @Accept  json
// @Produce  json
// @Success 200 {object} domainResponse
// @Failure 401 {object} domainResponse
// @Failure 500 {object} domainResponse
// @Router /api/v1/domains [get]
func GetDomains(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"domains":"aaaaa"})
}