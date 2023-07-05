package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
	"time"
)

func ShowMemMessage(ctx *gin.Context) gin.H {
	ipAddress := ctx.Param("ip")
	password := ctx.PostForm("pwd")
	fmt.Println(ipAddress)
	fmt.Println(password)
	MemStatus := gin.H{}
	cpuUsage := getMemUsage(ipAddress, "root", password)

	MemStatus["IpAddress:"] = ipAddress
	MemStatus["Utilization"] = cpuUsage + "%"
	return MemStatus
}

func getMemUsage(server string, user string, password string) string {

	// 创建SSH客户端配置
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout:         time.Second * 5,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接SSH服务器
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", server), config)
	if err != nil {
		return "error to connect ssh server"
	}
	defer conn.Close()

	// 创建SSH会话
	session, err := conn.NewSession()
	if err != nil {
		return "error to start ssh session"
	}
	defer session.Close()

	// 执行命令获取内存利用率信息

	output, err := session.Output("free | grep Mem | awk '{print $3/$2 * 100.0}'")
	if err != nil {
		return "err to execute command"
	}

	// 解析CPU利用率信息

	cpuUsage := string(output)[:2]

	return cpuUsage
}
