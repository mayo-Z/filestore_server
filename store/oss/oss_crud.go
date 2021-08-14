package oss

import (
	cfg "filestore_server/config"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//Bucket :获取bucket存储空间
func Bucket() *oss.Bucket {
	cli := Client()
	if cli != nil {
		bucket, err := cli.Bucket(cfg.OSSBucket)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		return bucket
	}
	return nil
}

//DownloadUrl :临时授权下载url
func DownloadUrl(objName string) string {
	signURL, err := Bucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return signURL

}
