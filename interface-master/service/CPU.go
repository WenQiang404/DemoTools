package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
	"strconv"
	"strings"
	"time"
)

func ShowCPUMessage(ctx *gin.Context) gin.H {
	ipAddress := ctx.Param("ip")
	password := ctx.PostForm("pwd")
	CpuStatus := gin.H{}
	cpuUsage, err := getCPUUsage(ipAddress, "root", password)
	if err != nil {
		fmt.Printf("获取服务器 %s 的CPU利用率失败：%s\n", ipAddress, err.Error())

	}

	CpuStatus["IpAddress:"] = ipAddress
	CpuStatus["Utilization"] = cpuUsage
	return CpuStatus
}

func getCPUUsage(server string, user string, password string) (float64, error) {

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
		return -999, err
	}
	defer conn.Close()

	// 创建SSH会话
	session, err := conn.NewSession()
	if err != nil {
		return -9, err
	}
	defer session.Close()

	// 执行命令获取CPU利用率信息
	//free | grep Mem | awk '{print $3/$2 * 100.0}'
	output, err := session.Output("top -b -n 1 | grep 'Cpu(s)'")
	if err != nil {
		return -99, err
	}

	// 解析CPU利用率信息
	fields := strings.Fields(string(output))
	if len(fields) < 8 {
		return -1, fmt.Errorf("无法解析CPU利用率信息：%s", string(output))
	}
	cpuUsageStr := strings.TrimSuffix(fields[1], "%")
	cpuUsage, err := strconv.ParseFloat(cpuUsageStr, 64)
	if err != nil {
		return 0, err
	}

	return cpuUsage, nil
}
