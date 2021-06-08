package controllers

import (
"github.com/gin-gonic/gin"
"net/http"
)

// GetViews godoc
// @Tags View
// @Summary Get all views
// @Description Get all views
// @Accept  json
// @Produce  json
// @Param search query string false "keyword"
// @Param Authorization  header string true "auth header"
// @Success 200 {object} viewResponse
// @Failure 401 {object} viewResponse
// @Failure 500 {object} viewResponse
// @Router /api/v1/views [get]
func GetViews(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"views":"aaaaa"})
}