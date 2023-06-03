package initalize

import (
	"github.com/h2non/filetype"
	"mime/multipart"
)

var goType = filetype.NewType("go", "go")

// go文件头： 112 97 99 107 97 103 101 32
func goMatcher(buf []byte) bool {
	return len(buf) > 1 && buf[0] == 112 && buf[1] == 97 && buf[2] == 99 && buf[3] == 107 && buf[4] == 97 && buf[5] == 103 && buf[6] == 101 && buf[7] == 32
}

// 初始化自定义文件头
func initFileType() {
	filetype.AddMatcher(goType, goMatcher)
}

// GetFileType  获取文件类型
func GetFileType(file *multipart.FileHeader) (string, string, error) {
	f, _ := file.Open()
	head := make([]byte, 261)
	_, err := f.Read(head)
	if err != nil {
		return "", "", err
	}
	kind, _ := filetype.Match(head)
	if kind == filetype.Unknown {
		return "", "", nil
	}
	return kind.Extension, kind.MIME.Value, nil
}
