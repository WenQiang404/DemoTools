package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
)

func ShowPortMessage(ctx *gin.Context) gin.H {
	hostName := ctx.Param("ip")
	portName := ctx.Param("port")
	ipStatus := checkPort(hostName, portName)
	return ipStatus
}

func checkPort(ip string, port string) gin.H {
	status := gin.H{}
	conn, err := net.Dial("tcp", net.JoinHostPort(ip, port))
	status["IpAddress"] = ip
	status["port"] = port
	if err != nil {
		fmt.Printf("Port %s:%s closed\n", ip, port)
		status["PortStatus:"] = "Closed"
	} else {
		conn.Close()
		status["PortStatus"] = "Listening"
	}
	return status
}
