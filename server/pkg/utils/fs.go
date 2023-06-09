package utils

import (
	"crypto/md5"
	"fiber-layout/validator/form"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

// GetDirDataList  获取目录下文件/文件夹
func GetDirDataList(path string) ([]form.ListResponse, error) {
	var l []form.ListResponse
	pwd, _ := os.Getwd()
	url := pwd + "/static" + path
	exists, err := PathExists(url)
	if err != nil {
		return nil, err
	}
	if exists {
		//获取文件或目录相关信息
		fileInfoList, err := ioutil.ReadDir(url)
		if err != nil {
			log.Fatal(err)
		}
		for i := range fileInfoList {
			l = append(l, form.ListResponse{
				Path:  path,
				IsDir: fileInfoList[i].IsDir(),
				Time:  fileInfoList[i].ModTime().Format("2006-01-02 15:04:05"),
				Name:  fileInfoList[i].Name(),
				Size:  fileInfoList[i].Size(),
			})
		}
		return l, err
	}
	return nil, nil
}

// GetDirFile 获取目录下文件信息
func GetDirFile(path string) (form.GetResponse, error) {
	var l form.GetResponse
	pwd, _ := os.Getwd()
	url := pwd + "/static" + path
	exists, err := PathExists(url)
	if err != nil {
		return l, err
	}
	fmt.Println(exists, "文件是否存在")
	if exists {
		//获取文件或目录相关信息
		fileInfoList, err := os.Stat(url)
		if err != nil {
			return l, err
		}
		l = form.GetResponse{
			Path:  path,
			IsDir: fileInfoList.IsDir(),
			Time:  fileInfoList.ModTime().Format("2006-01-02 15:04:05"),
			Name:  fileInfoList.Name(),
			Size:  fileInfoList.Size(),
		}
		return l, err
	}
	return l, nil
}

// PathExists 输入文件路径，根据返回的bool值来判断文件或文件夹是否存在  true存在文件夹/false不存在文件夹
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// AllExtMap 效验文件后缀
func AllExtMap(extMap []string, extName string) bool {
	allowExtMap := make(map[string]bool)
	for _, val := range extMap {
		allowExtMap[val] = true
	}
	// 判断excel上传是否合法
	if _, ok := allowExtMap[extName]; !ok {
		return false
	}
	return true
}

// ApiUpload api上传
func ApiUpload(extName, route string) (error, string, string) {
	pwd, _ := os.Getwd()
	// 组成 文件路径
	dir := pwd + "/static/upload/" + GetFileDay() + route
	// 创建文件路径
	if err := os.MkdirAll(dir, 0666); err != nil {
		return err, "", ""
	}
	//生成文件名称   144325235235.png
	fileUnixName := strconv.FormatInt(GetUnixNano(), 10)
	saveDir := path.Join(dir, fileUnixName+"--"+extName)
	return nil, saveDir, fileUnixName + "--" + extName
}

// Upload 文件上传
func Upload(extName, route string) (error, string, string) {
	pwd, _ := os.Getwd()
	// 组成 文件路径
	dir := pwd + "/static" + route
	// 创建文件路径
	if err := os.MkdirAll(dir, 0666); err != nil {
		return err, "", ""
	}
	saveDir := path.Join(dir, extName)
	return nil, saveDir, extName
}

// GetFileMd5 获取文件md5
func GetFileMd5(path string) string {
	md5hash := md5.New()
	f, _ := os.Open(path)
	io.Copy(md5hash, f)
	has := md5hash.Sum(nil)
	return fmt.Sprintf("%x", has)
}
