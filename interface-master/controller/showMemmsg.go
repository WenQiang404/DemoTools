package controller

import (
	"github.com/gin-gonic/gin"
	"interface/response"
	"interface/service"
)

func ShowMemoryMessage(ctx *gin.Context) {
	response.Success(ctx, service.ShowMemMessage(ctx))
}
