package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, httpStatus int, data gin.H) {
	ctx.JSON(httpStatus, data)
}

func Success(ctx *gin.Context, data gin.H) {
	Response(ctx, http.StatusOK, data)
}
