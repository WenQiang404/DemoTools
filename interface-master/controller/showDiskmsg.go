package controller

import (
	"github.com/gin-gonic/gin"
	"interface/response"
	"interface/service"
)

func ShowDiskMsg(ctx *gin.Context) {
	response.Success(ctx, service.ShowDiskmsg(ctx))
}
