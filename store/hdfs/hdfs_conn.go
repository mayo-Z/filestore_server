package hdfs

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
)

//var client *gowfs.FileSystem

func GetConnection() *gowfs.FileSystem {
	//这是配置，传入Addr: "ip: 9870", User: "随便写一个英文名就行"
	addr := "127.0.0.1:9870"
	user := "blackcat"
	config := gowfs.Configuration{Addr: addr, User: user}
	//返回一个客户端(这里面叫文件系统)和error
	client, err := gowfs.NewFileSystem(config)
	if err != nil {
		fmt.Println("hdfs连接出现异常,异常信息为:", err)
		return client
	}
	return client
}
