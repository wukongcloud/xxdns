package controllers

import (
"github.com/gin-gonic/gin"
"net/http"
)

func GetRecords(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"records":"aaaaa"})
}