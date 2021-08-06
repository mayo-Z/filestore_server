package main

import (
	. "filestore_server/store/hdfs"
)

func main() {
	//HdfsDeleteFile("/filestore_server/admin",true)
	//HdfsUploadFile("./test2/test.txt","/filestore_server/tmp",true)
	HdfsViewFile("/filestore_server/admin")
	//HdfsReadFile("/filestore_server/tmp/test.txt")
	//HdfsDownloadFile("/filestore_server/tmp/test.txt","./test2/test2.txt")
}
