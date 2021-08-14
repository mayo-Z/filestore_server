package main

import (
	. "filestore_server/store/hdfs"
)

func main() {
	//DeleteFile("/filestore_server/admin",true)
	//UploadFile("./test2/test.txt","/filestore_server/tmp",true)
	ViewFile("/filestore_server/admin")
	//ReadFile("/filestore_server/tmp/test.txt")
	//DownloadFile("/filestore_server/tmp/test.txt","./test2/test2.txt")
}
