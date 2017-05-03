package file

import (
	"encoding/binary"
	"encoding/hex"
	"strings"
)

//获取nSize长度字节
func ReadByte(data []byte, nSize int) ([]byte, []byte) {
	return data[:nSize], data[nSize:]
}

//data第1个字节为后续字节长度
func ReadByteUINT8(data []byte) ([]byte, []byte) {
	nSize, data := ReadUINT8(data)
	return ReadByte(data, int(nSize))
}

//data前面2个字节为后续字节长度
func ReadByteUINT16(data []byte) ([]byte, []byte) {
	nSize, data := ReadUINT16(data)
	return ReadByte(data, int(nSize))
}

//data前面2个字节为后续字节长度
func ReadByteUINT16BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := ReadUINT16BigEndian(data)
	return ReadByte(data, int(nSize))
}

//data前面4个字节为后续字节长度
func ReadByteUINT32(data []byte) ([]byte, []byte) {
	nSize, data := ReadUINT32(data)
	return ReadByte(data, int(nSize))
}

//data前面4个字节为后续字节长度
func ReadByteUINT32BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := ReadUINT32BigEndian(data)
	return ReadByte(data, int(nSize))
}

//data前面8个字节为后续字节长度 注意
func ReadByteUINT64(data []byte) ([]byte, []byte) {
	nSize, data := ReadUINT64(data)
	return data[0:nSize], data[nSize:]
}

//data前面8个字节为后续字节长度 注意
func ReadByteUINT64BigEndian(data []byte) ([]byte, []byte) {
	nSize, data := ReadUINT64BigEndian(data)
	return data[0:nSize], data[nSize:]
}

//获取nSize长度文本
func ReadString(data []byte, nSize int) (string, []byte) {
	return string(data[:nSize]), data[nSize:]
}

//data第1个字节为后续文本长度
func ReadStringUINT8(data []byte) (string, []byte) {
	nSize, data := ReadUINT8(data)
	return ReadString(data, int(nSize))
}

//data前2个字节为后续文本长度
func ReadStringUIN16(data []byte) (string, []byte) {
	nSize, data := ReadUINT16(data)
	return ReadString(data, int(nSize))
}

//data前2个字节为后续文本长度
func ReadStringUIN16BigEndian(data []byte) (string, []byte) {
	nSize, data := ReadUINT16BigEndian(data)
	return ReadString(data, int(nSize))
}

//data前4个字节为后续文本长度
func ReadStringUIN32(data []byte) (string, []byte) {
	nSize, data := ReadUINT32(data)
	return ReadString(data, int(nSize))
}

//data前4个字节为后续文本长度
func ReadStringUIN32BigEndian(data []byte) (string, []byte) {
	nSize, data := ReadUINT32BigEndian(data)
	return ReadString(data, int(nSize))
}

//data前4个字节为后续文本长度
func ReadStringUIN64(data []byte) (string, []byte) {
	nSize, data := ReadUINT64(data)
	return string(data[0:nSize]), data[nSize:]
}

//data前4个字节为后续文本长度
func ReadStringUIN64BigEndian(data []byte) (string, []byte) {
	nSize, data := ReadUINT64BigEndian(data)
	return string(data[0:nSize]), data[nSize:]
}

//读取uint8
func ReadUINT8(data []byte) (uint8, []byte) {
	return uint8(data[0:1][0]), data[1:]
}

//读取uint16
func ReadUINT16(data []byte) (uint16, []byte) {
	n := binary.LittleEndian.Uint16(data)
	return n, data[2:]
}

//读取uint16 BigEndian
func ReadUINT16BigEndian(data []byte) (uint16, []byte) {
	n := binary.BigEndian.Uint16(data)
	return n, data[2:]
}

//读取uint32
func ReadUINT32(data []byte) (uint32, []byte) {
	n := binary.LittleEndian.Uint32(data)
	return n, data[4:]
}

//读取uint32 BigEndian
func ReadUINT32BigEndian(data []byte) (uint32, []byte) {
	n := binary.BigEndian.Uint32(data)
	return n, data[4:]
}

//读取uint64
func ReadUINT64(data []byte) (uint64, []byte) {
	n := binary.LittleEndian.Uint64(data)
	return n, data[8:]
}

//读取uint64 BigEndian
func ReadUINT64BigEndian(data []byte) (uint64, []byte) {
	n := binary.BigEndian.Uint64(data)
	return n, data[8:]
}

//把16进制转化为文本形式
func ReadHexToString(data []byte, nSize int) (string, []byte) {
	s := strings.ToUpper(hex.EncodeToString(data[:nSize]))
	return s, data[nSize:]
}
