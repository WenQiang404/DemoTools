package controller

import (
	"DemoTools/interface-master/response"
	"DemoTools/interface-master/service"
	"github.com/gin-gonic/gin"
)

func ShowPortMessage(c *gin.Context) {
	response.Success(c, service.ShowPortMessage(c))

}
