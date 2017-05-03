package file

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//打开指定文件，并从指定位置写入数据
func WriteAt(name string, b []byte, off int64) error {
	file, err := os.OpenFile(name, os.O_WRONLY, 0666)
	if hasErr(err) {
		return err
	}
	defer file.Close()
	_, err = file.WriteAt(b, off)
	return err
}

//打开指定文件,并在文件末尾写入数据
func Write(name string, b []byte) error {
	file, err := os.OpenFile(name, os.O_APPEND, 0666)
	if hasErr(err) {
		return err
	}
	defer file.Close()
	_, err = file.Write(b)
	return err
}

//检察文件是否存在
func Exist(name string) bool {
	_, err := os.Stat(name)
	if hasErr(err) {
		return false
	}
	return true
}

//检察文件是否允许读
func AllowRead(name string) bool {
	// os.O_RDONLY 读权限
	_, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if hasErr(err) {
		return false
	}
	return true
}

//检察文件是否允许写
func AllowWrite(name string) bool {
	// os.O_WRONLY 写权限
	_, err := os.OpenFile(name, os.O_WRONLY, 0666)
	if hasErr(err) {
		return false
	}
	return true
}

//复制文件
func Copy(src, dst string) (w int64, err error) {
	//打开源文件
	originalFile, err := os.Open(src)
	if hasErr(err) {
		return 0, err
	}
	defer originalFile.Close()
	// 创建新的文件作为目标文件
	newFile, err := os.Create(dst)
	if hasErr(err) {
		return 0, err
	}
	defer newFile.Close()
	// 从源中复制字节到目标文件
	bytesWritten, err := io.Copy(newFile, originalFile)
	if hasErr(err) {
		return 0, err
	}
	//将文件内容flush到硬盘中
	err = newFile.Sync()
	if hasErr(err) {
		return 0, err
	}
	//最后返回
	return bytesWritten, nil
}

//获取无后缀的文件名
func FileNameNoSuffix(name string) string {
	return strings.TrimSuffix(filepath.Base(name), filepath.Ext(name))
}

//遍历目录及下级目录，查找符合后缀文件
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

//遍历指定目录下的所有文件，查找符合后缀文件,不进入下一级目录搜索。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}
