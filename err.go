package file

import (
	"log"
	"runtime"
	"strings"
)

//判断有无错误仅返回true或者false
func hasErr(err error) bool {
	if err != nil {
		return true
	}
	return false
}

//返回true或者false，有错误时打印错误
func hasErrPrintln(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		file = file[strings.LastIndex(file, `/`)+1:]
		log.Printf("%v,第%v行,错误类型:%v", file, line, err)
		return true
	}
	return false
}

//返回true或者false，有错误时程序退出
func hasErrFatal(err error) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		file = file[strings.LastIndex(file, `/`)+1:]
		log.Fatalf("%v,第%v行,错误类型:%v", file, line, err)
		return true
	}
	return false
}
