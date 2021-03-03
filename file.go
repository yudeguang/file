// Copyright 2020 file Author(https://github.com/yudeguang/file). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/yudeguang/file.
package file

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
)

//检察文件是否允许读
func AllowRead(path string) bool {
	_, err := os.OpenFile(path, os.O_RDONLY, 0666)
	return err == nil
}

//检察文件是否允许写
func AllowWrite(path string) bool {
	_, err := os.OpenFile(path, os.O_WRONLY, 0666)
	return err == nil
}

//复制文件，目标文件所在目录不存在，则创建目录后再复制
//Copy(`d:\test\hello.txt`,`c:\test\hello.txt`)
func Copy(dstFileName, srcFileName string) (w int64, err error) {
	//打开源文件
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()
	// 创建新的文件作为目标文件
	dstFile, err := os.Create(dstFileName)
	if err != nil {
		//如果出错，很可能是目标目录不存在，需要先创建目标目录
		err = os.MkdirAll(filepath.Dir(dstFileName), 0666)
		if err != nil {
			return 0, err
		}
		//再次尝试创建
		dstFile, err = os.Create(dstFileName)
		if err != nil {
			return 0, err
		}
	}
	defer dstFile.Close()
	//通过bufio实现对大文件复制的自动支持
	dst := bufio.NewWriter(dstFile)
	defer dst.Flush()
	src := bufio.NewReader(srcFile)
	w, err = io.Copy(dst, src)
	if err != nil {
		return 0, err
	}
	return w, err
}

//复制目录
func CopyDir(dst string, src string) (err error) {
	if src == "" {
		return fmt.Errorf("source directory cannot be empty")
	}
	if dst == "" {
		return fmt.Errorf("destination directory cannot be empty")
	}
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)
	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}
	if !Exist(dst) {
		err = os.MkdirAll(dst, si.Mode())
		if err != nil {
			return
		}
	}
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			_, err = Copy(dstPath, srcPath)
			if err != nil {
				return
			}
		}
	}
	return
}

//获得程序所在当前文件路径 注意末尾不包含/
//d:\test\hello.exe  ==>     d:\test
func CurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

//获取可执行文件的绝对路径（包括文件名）
//d:\test\hello.exe
func CurrentExePath() string {
	exeName, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	return strings.Replace(exeName, "\\", "/", -1)
}

//检察文件是或者目录否存在
func Exist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

//返回不包含文件路径的文件名
//c:\test\hello.txt ==> hello.txt
func FileName(path string) string {
	return filepath.Base(path)
}

//返回文件的后缀名
//c:\test\hello.txt ==> .txt
func FileNameExt(path string) string {
	return filepath.Ext(path)
}

//获取无后缀的文件名
//c:\test\hello.txt ==> hello
func FileNameNoExt(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

//获取文件所在目录, 注意末尾不包含/
func FileNameDir(path string) string {
	return filepath.Dir(path)
}

//获得文件的修改时间
func FileModTime(path string) (int64, error) {
	f, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

//返回文件的大小
func FileSize(path string) (int64, error) {
	f, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

//遍历目录及下级目录，查找符合后缀文件,如果suffix为空，则查找所有文件
//如果后缀为空，则查找任何后缀的文件
func GetFileListBySuffix(dirPath, suffix string) (files []string, err error) {
	if !IsDir(dirPath) {
		return nil, fmt.Errorf("given path does not exist: %s", dirPath)
	}
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                      //忽略后缀匹配的大小写
	err = filepath.Walk(dirPath, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if suffix == "" {
			files = append(files, filename)
		} else {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				files = append(files, filename)
			}
		}
		return nil
	})
	return files, err
}

//遍历指定目录下的所有文件，查找符合后缀文件,不进入下一级目录搜索
//如果后缀为空，则查找任何后缀的文件
func GetFileListJustCurrentDirBySuffix(dirPath string, suffix string) (files []string, err error) {
	if !IsDir(dirPath) {
		return nil, fmt.Errorf("given path does not exist: %s", dirPath)
	}
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	PathSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if suffix == "" {
			files = append(files, dirPath+PathSep+fi.Name())
		} else {
			if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
				files = append(files, dirPath+PathSep+fi.Name())
			}
		}
	}
	return files, nil
}

//把文件大小转换成人更加容易看懂的文本
// HumaneFileSize calculates the file size and generate user-friendly string.
func HumaneFileSize(s uint64) string {
	logn := func(n, b float64) float64 {
		return math.Log(n) / math.Log(b)
	}
	humanateBytes := func(s uint64, base float64, sizes []string) string {
		if s < 10 {
			return fmt.Sprintf("%dB", s)
		}
		e := math.Floor(logn(float64(s), base))
		suffix := sizes[int(e)]
		val := float64(s) / math.Pow(base, math.Floor(e))
		f := "%.0f"
		if val < 10 {
			f = "%.1f"
		}
		return fmt.Sprintf(f+"%s", val, suffix)
	}
	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1024, sizes)
}

//判断是否是目录
func IsDir(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return f.IsDir()
}

//返回是否是文件
func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

//打开指定文件，并从指定位置写入数据
func WriteAt(path string, b []byte, off int64) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteAt(b, off)
	return err
}

//打开指定文件,并在文件末尾写入数据
func WriteAppend(path string, b []byte) error {
	file, err := os.OpenFile(path, os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(b)
	return err
}
