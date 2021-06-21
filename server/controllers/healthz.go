package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type healthResponse struct {
	status string `json:"status"`
	err    string `json:"error"`
}

// GetHealth godoc
// @Tags Health
// @Summary Get health status
// @Description Get health status
// @Accept  json
// @Produce  json
// @Success 200 {object} healthResponse{status=string}
// @Failure 500 {object} healthResponse{status=string,error=string}
// @Router /healthz [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
