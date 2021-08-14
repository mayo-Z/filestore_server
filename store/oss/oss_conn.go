package oss

import (
	cfg "filestore_server/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossCli *oss.Client

//Client ：创建oss client对象
func Client() *oss.Client {
	if ossCli != nil {
		return ossCli

	}

	ossCli, err := oss.New(cfg.OSSEndpoint, cfg.OSSAccesskeyID, cfg.OSSAccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil
	}
	return ossCli
}
