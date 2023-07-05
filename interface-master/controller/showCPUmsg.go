package controller

import (
	"DemoTools/interface-master/response"
	"DemoTools/interface-master/service"
	"github.com/gin-gonic/gin"
)

func ShowCPUMsg(ctx *gin.Context) {
	response.Success(ctx, service.ShowCPUMessage(ctx))
}
