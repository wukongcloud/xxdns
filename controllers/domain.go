package controllers

import (
"github.com/gin-gonic/gin"
"net/http"
)

func GetDomains(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"domains":"aaaaa"})
}