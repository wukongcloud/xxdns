package controllers
import (
	_ "github.com/gin-gonic/gin"
	"github.com/wukongcloud/xxdns/models"
)

type response struct {
	Code int         `json:"code"`
	Err  string      `json:"error"`
	Data interface{} `json:"data"`
}

type viewResponse struct {
	response
	Data []models.View
}

type domainResponse struct {
	response
	Data []models.Domain
}