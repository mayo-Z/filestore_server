package handler

import (
	rPool "filestore_server/cache/redis"
	dblayer "filestore_server/db"
	"filestore_server/util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

//MultipartUploadInfo :初始化信息
type MultipartUploadInfo struct {
	FileHash   string
	FileSize   int
	UploadID   string
	ChunkSize  int
	ChunkCount int
}

//InitialMultipartUploadHandler :初始化分块上传
func InitialMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	//	1.解析用户请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize, err := strconv.Atoi(r.Form.Get("filesize"))

	if err != nil {
		w.Write(util.NewRespMsg(-1, "params invalid", nil).JSONBytes())
		return
	}

	// 	2.获得redis的一个连接
	rConn := rPool.RedisPool().Get()
	defer rConn.Close()
	//	3.生成分块上传的初始化信息
	upInfo := MultipartUploadInfo{
		FileHash:  filehash,
		FileSize:  filesize,
		UploadID:  username + fmt.Sprintf("%x", time.Now().UnixNano()),
		ChunkSize: 5 * 1024 * 1024, //5MB
		//Ceil这个函数是对数值向上取整
		ChunkCount: int(math.Ceil(float64(filesize) / (5 * 1024 * 1024))),
	}

	//	4.将初始化信息写入到redis缓存
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "chunkcount", upInfo.ChunkCount)
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "filehash", upInfo.FileHash)
	rConn.Do("HSET", "MP_"+upInfo.UploadID, "filesize", upInfo.FileSize)
	//	5.将响应初始化数据返回到客户端
	w.Write(util.NewRespMsg(0, "OK", upInfo).JSONBytes())
}

// UploadPartHandler : 上传文件分块
func UploadPartHandler(w http.ResponseWriter, r *http.Request) {
	//	1.解析用户请求参数
	r.ParseForm()
	uploadID := r.Form.Get("uploadid")
	chunkIndex := r.Form.Get("index")

	//	2.获得redis连接池中的一个连接
	rConn := rPool.RedisPool().Get()
	defer rConn.Close()
	//  3.获得文件句柄
	fpath := "./data/" + uploadID + "/" + chunkIndex
	//fpath := "C:\\Users\\blackcat\\Desktop\\goland_work\\filestore_server/data/" + uploadID + "/" + chunkIndex
	os.MkdirAll(path.Dir(fpath), os.ModePerm)
	fd, err := os.Create(fpath)
	if err != nil {
		w.Write(util.NewRespMsg(-1, "upload part failed", nil).JSONBytes())
		return
	}
	defer fd.Close()

	buf := make([]byte, 1024*1024)
	for {
		n, err := r.Body.Read(buf)
		fd.Write(buf[:n])
		if err != nil {
			break
		}
	}
	//	4.更新redis缓存状态
	rConn.Do("HSet", "MP_"+uploadID, "chkidx_"+chunkIndex, 1)

	//	5.返回处理结果到客户端
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())

}

// CompleteUploadHandler : 通知上传合并
func CompleteUploadHandler(w http.ResponseWriter, r *http.Request) {
	//1.解析请求参数
	r.ParseForm()
	upid := r.Form.Get("uploadid")
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filesize := r.Form.Get("filesize")
	filename := r.Form.Get("filename")

	//2.获得redis连接池中的一个连接
	rConn := rPool.RedisPool().Get()
	defer rConn.Close()

	//3.通过uploadid查询redis并判断是否所有分块上传完成

	//Redis Hgetall 命令用于返回哈希表中，所有的字段和值。
	//在返回值里，紧跟每个字段名(field name)之后是字段的值(value)，所以返回值的长度是哈希表大小的两倍。
	data, err := redis.Values(rConn.Do("Hgetall", "MP_"+upid))
	if err != nil {
		fmt.Println(err.Error())
		w.Write(util.NewRespMsg(-1, "complete upload failed", nil).JSONBytes())
		return
	}
	totalCount := 0
	chunkCount := 0
	for i := 0; i < len(data); i += 2 {
		k := string(data[i].([]byte))
		v := string(data[i+1].([]byte))
		if k == "chunkcount" {
			totalCount, err = strconv.Atoi(v)
			//	strings.HasPrefix()函数用来检测字符串是否以指定的前缀开头
		} else if strings.HasPrefix(k, "chkidx_") && v == "1" {
			chunkCount += 1
		}
	}
	if totalCount != chunkCount {
		w.Write(util.NewRespMsg(-2, "invalid request", nil).JSONBytes())
		return
	}

	//4.TODO: 合并分块

	//5.更新唯一文件表及用户表
	fsize, _ := strconv.Atoi(filesize)
	dblayer.OnFileUploadFinished(filehash, filename, int64(fsize), "")
	dblayer.OnUserFileUploadFinished(username, filehash, filename, int64(fsize))
	//6.响应处理结果
	w.Write(util.NewRespMsg(0, "OK", nil).JSONBytes())
}

//CancelUploadPartHandler :通知取消上传
func CancelUploadPartHandler(w http.ResponseWriter, r *http.Request) {
	//	删除已存在的分块文件

	//  删除redis缓存状态

	//	更新mysql文件status

}

//MultipartUploadStatusHandler :查看分块上传状态
func MultipartUploadStatusHandler(w http.ResponseWriter, r *http.Request) {
	//	检查分块上传状态是否有效

	//	获取分块初始化信息

	//	获取已上传的分块信息

}
