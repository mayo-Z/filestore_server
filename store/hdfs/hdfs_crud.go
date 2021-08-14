package hdfs

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"io/ioutil"
)

func UploadFile(localFile, hdfsPath string, overwrite bool) bool {
	client := GetConnection()
	//创建一个shell,可以使用下面shell进行操作
	shell := gowfs.FsShell{FileSystem: client}
	//本地路径，hdfs路径，是否重写
	_, err := shell.Put(localFile, hdfsPath, overwrite)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func ViewFile(hdfsPath string) bool {
	client := GetConnection()
	path := gowfs.Path{Name: hdfsPath}
	fsArr, err := client.ListStatus(path)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	// 返回[]FileStatus和error
	//这个FileStatus是什么？我们看一下源码
	/*
		type FileStatus struct {
			AccesTime        int64  访问时间
			BlockSize        int64  块大小,只针对文件(134217728 Bytes,128 MB)，目录的话为0
			Group            string 所属组
			Length           int64  文件的字节数(目录为0)
			ModificationTime int64  修改时间
			Owner            string 所有者
			PathSuffix       string 文件后缀，说白了就是文件名
			Permission       string 权限
			Replication      int64  副本数
			Type             string 类型，文本的话是FILE，目录的话是DIRECTORY
		}
	*/
	fmt.Println(fsArr)
	// [{0 0 supergroup 0 1570359570447 dr.who tkinter 755 0 DIRECTORY} {0 134217728 supergroup 184 1570359155457 root whitealbum.txt 644 1 FILE}]
	for _, fs := range fsArr {
		fmt.Println("文件名：", fs.PathSuffix)
	}
	//FileStatus里面包含了文件的详细信息，如果想查看某个文件的详细信息
	//可以使用fs, err := client.GetFileStatus(path)
	//fs, err := client.GetFileStatus(path)
	return true
}

func ReadFile(file string) bool {
	client := GetConnection()
	path := gowfs.Path{Name: file}
	//接收如下参数：gowfs.Path,offset(偏移量),长度(显然是字节的长度), 容量(自己的cap)
	//返回一个io.ReadCloser，这是需要实现io.Reader和io.Closer的接口
	reader, err := client.Open(path, 0, 512, 2048)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	//可以使用reader.Read(buf)的方式循环读取，也可以丢给ioutil。ReadAll，一次性全部读取
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Println(string(data))

	return true
}

func DeleteFile(hdfsPath string, recursive bool) bool {
	client := GetConnection()
	path := gowfs.Path{Name: hdfsPath}
	//路径，是否递归
	flag, err := client.Delete(path, recursive)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return flag
}

func DownloadFile(hdfsPath, localFile string) bool {
	client := GetConnection()
	shell := gowfs.FsShell{FileSystem: client}

	flag, err := shell.Get(hdfsPath, localFile)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return flag
}
