package file

import (
	"code.google.com/p/mahonia"
	"encoding/binary"
	"encoding/hex"
	"strings"
)

//bytes  型数据读写操作

var gbkDecoder = mahonia.NewDecoder("gbk")

func FetchString(data []byte, nSize int) (string, []byte) {
	s := string(data[:nSize])
	s = strings.TrimRight(s, " ")
	s = GBKToUTF8(s)
	return s, data[nSize:]
}

func FetchUINT16(data []byte) (uint16, []byte) {
	n := binary.LittleEndian.Uint16(data)
	return n, data[2:]
}
func FetchUINT32(data []byte) (uint32, []byte) {
	n := binary.LittleEndian.Uint32(data)
	return n, data[4:]
}

func FetchByte(data []byte) (byte, []byte) {
	return data[0], data[1:]
}

/*
获取未知的字节并把16进制直接展示成文本形式
*/
func FetchUnknownBytes(data []byte, nSize int) (string, []byte) {
	s := strings.ToUpper(hex.EncodeToString(data[:nSize]))
	return s, data[nSize:]
}

/*
把GBK转换为UTF8
*/
func GBKToUTF8(s string) string {
	return gbkDecoder.ConvertString(s)
}
