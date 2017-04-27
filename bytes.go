package file

import (
	"encoding/binary"
	"encoding/hex"
	"strings"
)

//获取nSize长度字节
func FetchByte(data []byte, nSize int) ([]byte, []byte) {
	return data[:nSize], data[nSize:]
}

//data第1个字节为后续字节长度
func FetchByteUINT8(data []byte) ([]byte, []byte) {
	nSize, data := FetchUINT8(data)
	return FetchByte(data, int(nSize))
}

//data前面2个字节为后续字节长度
func FetchByteUINT16(data []byte) ([]byte, []byte) {
	nSize, data := FetchUINT16(data)
	return FetchByte(data, int(nSize))
}

//data前面2个字节为后续字节长度
func FetchByteUINT16BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := FetchUINT16BigEndian(data)
	return FetchByte(data, int(nSize))
}

//data前面4个字节为后续字节长度
func FetchByteUINT32(data []byte) ([]byte, []byte) {
	nSize, data := FetchUINT32(data)
	return FetchByte(data, int(nSize))
}

//data前面4个字节为后续字节长度
func FetchByteUINT32BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := FetchUINT32BigEndian(data)
	return FetchByte(data, int(nSize))
}

//data前面8个字节为后续字节长度 注意
func FetchByteUINT64(data []byte) ([]byte, []byte) {
	nSize, data := FetchUINT64(data)
	return data[0:nSize], data[nSize:]
}

//data前面8个字节为后续字节长度 注意
func FetchByteUINT64BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := FetchUINT64BigEndian(data)
	return data[0:nSize], data[nSize:]
}

//获取nSize长度文本
func FetchString(data []byte, nSize int) (string, []byte) {
	return string(data[:nSize]), data[nSize:]
}

//data第1个字节为后续文本长度
func FetchStringUINT8(data []byte) (string, []byte) {
	nSize, data := FetchUINT8(data)
	return FetchString(data, int(nSize))
}

//data前2个字节为后续文本长度
func FetchStringUIN16(data []byte) (string, []byte) {
	nSize, data := FetchUINT16(data)
	return FetchString(data, int(nSize))
}

//data前2个字节为后续文本长度
func FetchStringUIN16BigEndian(data []byte) (string, []byte) {
	nSize, data := FetchUINT16BigEndian(data)
	return FetchString(data, int(nSize))
}

//data前4个字节为后续文本长度
func FetchStringUIN32(data []byte) (string, []byte) {
	nSize, data := FetchUINT32(data)
	return FetchString(data, int(nSize))
}

//data前4个字节为后续文本长度
func FetchStringUIN32BigEndian(data []byte) (string, []byte) {
	nSize, data := FetchUINT32BigEndian(data)
	return FetchString(data, int(nSize))
}

//data前4个字节为后续文本长度
func FetchStringUIN64(data []byte) (string, []byte) {
	nSize, data := FetchUINT64(data)
	return string(data[0:nSize]), data[nSize:]
}

//data前4个字节为后续文本长度
func FetchStringUIN64BigEndian(data []byte) (string, []byte) {
	nSize, data := FetchUINT64BigEndian(data)
	return string(data[0:nSize]), data[nSize:]
}

//读取uint8
func FetchUINT8(data []byte) (uint8, []byte) {
	return uint8(data[0:1][0]), data[1:]
}

//读取uint16
func FetchUINT16(data []byte) (uint16, []byte) {
	n := binary.LittleEndian.Uint16(data)
	return n, data[2:]
}

//读取uint16 BigEndian
func FetchUINT16BigEndian(data []byte) (uint16, []byte) {
	n := binary.BigEndian.Uint16(data)
	return n, data[2:]
}

//读取uint32
func FetchUINT32(data []byte) (uint32, []byte) {
	n := binary.LittleEndian.Uint32(data)
	return n, data[4:]
}

//读取uint32 BigEndian
func FetchUINT32BigEndian(data []byte) (uint32, []byte) {
	n := binary.BigEndian.Uint32(data)
	return n, data[4:]
}

//读取uint64
func FetchUINT64(data []byte) (uint64, []byte) {
	n := binary.LittleEndian.Uint64(data)
	return n, data[8:]
}

//读取uint64 BigEndian
func FetchUINT64BigEndian(data []byte) (uint64, []byte) {
	n := binary.BigEndian.Uint64(data)
	return n, data[8:]
}

//把16进制转化为文本形式
func FetchHexToString(data []byte, nSize int) (string, []byte) {
	s := strings.ToUpper(hex.EncodeToString(data[:nSize]))
	return s, data[nSize:]
}
