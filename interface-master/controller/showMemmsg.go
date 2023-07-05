package controller

import (
	"DemoTools/interface-master/response"
	"DemoTools/interface-master/service"
	"github.com/gin-gonic/gin"
)

func ShowMemoryMessage(ctx *gin.Context) {
	response.Success(ctx, service.ShowMemMessage(ctx))
}
