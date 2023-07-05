package router

import (
	"DemoTools/interface-master/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api/Mem/:ip", controller.ShowMemoryMessage)
	router.POST("/api/Disk/:ip", controller.ShowDiskMsg)
	router.POST("/api/CPU/:ip", controller.ShowCPUMsg)
	router.POST("/api/port/:ip/:port", controller.ShowPortMessage)

	return router
}
