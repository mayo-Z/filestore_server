package meta

import "time"

const baseFormat = "2006-01-02 15:04:05"

type ByUploadTime []FileMeta

func (receiver ByUploadTime) Len() int {
	return len(receiver)
}

func (receiver ByUploadTime) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}

func (receiver ByUploadTime) Less(i, j int) bool {
	iTime, _ := time.Parse(baseFormat, receiver[i].UploadAt)
	jTime, _ := time.Parse(baseFormat, receiver[j].UploadAt)
	return iTime.UnixNano() > jTime.UnixNano()

}
