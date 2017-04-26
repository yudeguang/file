package file

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 创建文件
func Create(name string) (*os.File, error) {
	return os.Create(name)
}

//只读方式打开文件
func OpenFileReadOnly(name string) (*os.File, error) {
	return os.Open(name)
}

//以可读写方式打开
func OpenFile(name string) (*os.File, error) {
	return os.OpenFile(name, os.O_APPEND, 0666)
	// OpenFile提供更多的选项。
	// 最后一个参数是权限模式permission mode
	// 第二个是打开时的属性
	// 下面的属性可以单独使用，也可以组合使用。
	// 组合使用时可以使用 OR 操作设置 OpenFile的第二个参数，例如：
	// os.O_CREATE|os.O_APPEND
	// 或者 os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // 只读
	// os.O_WRONLY // 只写
	// os.O_RDWR // 读写
	// os.O_APPEND // 往文件中添建（Append）
	// os.O_CREATE // 如果文件不存在则先创建
	// os.O_TRUNC // 文件打开时裁剪文件
	// os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC // 以同步I/O的方式打开
}

//打开并读取全部内容到内存
func ReadAll(name string) ([]byte, error) {
	return ioutil.ReadFile(name)
}

//截取文件，清空文件
func Truncate(name string, size int64) error {
	// 裁剪一个文件到100个字节。
	// 如果文件本来就少于100个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	// 如果文件本来超过100个字节，则超过的字节会被抛弃。
	// 这样我们总是得到精确的100个字节的文件。
	// 传入0则会清空文件。
	return os.Truncate(name, size)
}

//获得文件基本信息
func Stat(name string) (os.FileInfo, error) {
	fileInfo, err := os.Stat(name)
	if hasErr(err) {
		return fileInfo, err
	}
	log.Println("File name:", fileInfo.Name())        //文件名
	log.Println("Size in bytes:", fileInfo.Size())    //大小
	log.Println("Permissions:", fileInfo.Mode())      //权限
	log.Println("Last modified:", fileInfo.ModTime()) //最后更新时间
	log.Println("Is Directory: ", fileInfo.IsDir())   //是否是目录
	log.Printf("System interface type: %T\n", fileInfo.Sys())
	log.Printf("System info: %+v\n\n", fileInfo.Sys())
	return fileInfo, err
}

//重命名
func Rename(originalPath, newPath string) error {
	return os.Rename(originalPath, newPath)
}

//移动文件
func Move(originalPath, newPath string) error {
	//移动和重命名实质上是同一个函数
	return Rename(originalPath, newPath)
}

//删除文件
func Remove(name string) error {
	return os.Remove(name)
}

//文件是否存在
func Exist(name string) bool {
	_, err := os.Stat(name)
	if hasErr(err) {
		return false
	}
	return true
}

//文件是否允许读
func AllowRead(name string) bool {
	// os.O_RDONLY 读权限
	_, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if hasErr(err) {
		return false
	}
	return true
}

//文件是否允许写
func AllowWrite(name string) bool {
	// os.O_WRONLY 写权限
	_, err := os.OpenFile(name, os.O_WRONLY, 0666)
	if hasErr(err) {
		return false
	}
	return true
}

//修改文件权限
func Chmod(name string, mode os.FileMode) error {
	// err := os.Chmod("test.txt", 0777); 0777表示权限
	return os.Chmod(name, mode)
}

//修改文件拥有者
func Chown(name string) error {
	return os.Chown(name, os.Getuid(), os.Getgid())
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

//写文件,追加方式写
func Write(name string, b []byte) error {
	//例：Write("test.txt", []byte("hello,word"))
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if hasErr(err) {
		return err
	}
	defer file.Close()
	_, err = file.Write(b)
	return err
}

//快写文件
func WriteFile(name string, b []byte) error {
	//ioutil包有一个非常有用的方法WriteFile()可以处理创建／打开文件、写字节slice
	//和关闭文件一系列的操作。如果你需要简洁快速地写字节slice到文件中，你可以使用它。
	return ioutil.WriteFile(name, b, 0666)
}

//使用缓存写
func WriteWithBufer(name string, b []byte) error {
	// bufio包提供了带缓存功能的writer，所以你可以在写字节到硬盘前使用内存缓存。
	// 当你处理很多的数据很有用，因为它可以节省操作硬盘I/O的时间。在其它一些情况
	// 下它也很有用，比如你每次写一个字节，把它们攒在内存缓存中，然后一次写入到硬
	// 盘中，减少硬盘的磨损以及提升性能。

	// 打开文件，只写
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if hasErr(err) {
		return err
	}
	defer file.Close()
	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)
	// 写字节到buffer
	_, err = bufferedWriter.Write(b)
	if hasErr(err) {
		return err
	}
	// 写内存buffer到硬盘
	return bufferedWriter.Flush()
}

//遍历目录及下级目录，查找符合后缀文件
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToUpper(suffix)                                                     //忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
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
