package controller

import (
	"github.com/gin-gonic/gin"
	"interface/response"
	"interface/service"
)

func ShowCPUMsg(ctx *gin.Context) {
	response.Success(ctx, service.ShowCPUMessage(ctx))
}
