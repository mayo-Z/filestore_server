package meta

import (
	mydb "filestore_server/db"
	"sort"
)

//filemeta:文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//新增/更新元信息
func UpdateFileMeta(meta FileMeta) {
	fileMetas[meta.FileSha1] = meta
}

func UpdateFileMetaDB(meta FileMeta) bool {
	return mydb.OnFileUploadFinished(meta.FileSha1, meta.FileName, meta.FileSize, meta.Location)

}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
func GetFileMetaDB(fileSha1 string) (*FileMeta, error) {
	tfile, err := mydb.GetFileMeta(fileSha1)
	if tfile == nil || err != nil {
		return nil, err
	}

	fmeta := FileMeta{
		FileSha1: tfile.FileHash,
		FileName: tfile.FileName.String,
		FileSize: tfile.FileSize.Int64,
		Location: tfile.FileAddr.String,
	}
	return &fmeta, nil

}

//GetLastFileMetas:获取批量文件元信息列表
func GetLastFileMetas(count int) []FileMeta {
	fMetaArray := make([]FileMeta, len(fileMetas))
	for _, v := range fileMetas {
		fMetaArray = append(fMetaArray, v)
	}

	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}

//删除信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
