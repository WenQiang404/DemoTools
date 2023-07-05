package controller

import (
	"github.com/gin-gonic/gin"
	"interface/response"
	"interface/service"
)

func ShowPortMessage(c *gin.Context) {
	response.Success(c, service.ShowPortMessage(c))

}
