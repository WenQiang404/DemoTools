package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
	"strings"
	"time"
)

func ShowDiskmsg(ctx *gin.Context) gin.H {
	ipAddress := ctx.Param("ip")
	password := ctx.PostForm("pwd")
	diskMsg := gin.H{}
	msg := getDiskMessage(ipAddress, "root", password)
	diskMsg["IpAddress"] = ipAddress
	diskMsg["diskMsg"] = msg

	return diskMsg
}

func getDiskMessage(ip string, user string, pwd string) gin.H {
	// 创建SSH客户端配置
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
		},
		Timeout:         time.Second * 5,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接SSH服务器
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:22", ip), config)
	if err != nil {
		return gin.H{"error": err.Error()}
	}
	defer conn.Close()

	// 创建SSH会话
	session, err := conn.NewSession()
	if err != nil {
		return gin.H{"error": err.Error()}
	}
	defer session.Close()

	// 执行命令获取CPU利用率信息
	cmd := "df -h | awk '{print $1,$2,$3,$4,$5,$6}'"
	output, err := session.Output(cmd)
	if err != nil {

		return gin.H{"error": err.Error()}
	}

	// 解析命令输出，获取硬盘信息
	lines := strings.Split(string(output), "\n")
	disks := make([]map[string]string, 0)
	for _, line := range lines[1:] {
		fields := strings.Fields(line)
		if len(fields) == 6 {
			size := fields[1]
			used := fields[2]
			avail := fields[3]
			usage := fields[4]
			mount := fields[5]
			disks = append(disks, map[string]string{
				"文件系统总大小": size,
				"已用空间":    used,
				"可用空间":    avail,
				"使用率":     usage,
				"挂载地址":    mount,
			})
		}
	}

	// 返回硬盘信息
	return gin.H{"disks": disks}

}
