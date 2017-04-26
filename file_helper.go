package file

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

// 文件读取对象
type FileHelper struct {
	pFile io.ReadSeeker
}

//OS文件打开
func (this *FileHelper) OpenOsFile(filepath string) error {
	var err error = nil
	if this.pFile, err = os.OpenFile(filepath, os.O_RDONLY, 0x666); err != nil {
		return err
	}
	return nil
}

//字节数组的打开
func (this *FileHelper) OpenBytes(p *bytes.Reader) error {
	this.pFile = p
	return nil
	//打开案例
	// m := &FileHelper{}
	// m.OpenBytes(bytes.NewReader([]byte{0xda}))
}

//字符串类型的打开
func (this *FileHelper) OpenStrings(p *strings.Reader) error {
	this.pFile = p
	return nil
	//打开案例
	// m := &FileHelper{}
	// m.OpenStrings(strings.NewReader("hello world"))
}

//关闭文件
func (this *FileHelper) Close() {
	if this.pFile != nil {
		if osFile, ok := this.pFile.(*os.File); ok {
			osFile.Close()
		}
		this.pFile = nil
	}
}

//移动到文件什么地方
func (this *FileHelper) MoveTo(npos int) error {
	_, err := this.pFile.Seek(int64(npos), os.SEEK_SET)
	if err != nil {
		return err
	}
	return nil
}

//移动到文件末尾
func (this *FileHelper) MoveToEnd() error {
	_, err := this.pFile.Seek(int64(0), os.SEEK_END)
	if err != nil {
		return err
	}
	return nil
}

//从当前位置移动多少个
func (this *FileHelper) Move(npos int) error {
	_, err := this.pFile.Seek(int64(npos), os.SEEK_CUR)
	if err != nil {
		return err
	}
	return nil
}

//获得当前位置
func (this *FileHelper) GetFilePos() (int64, error) {
	return this.pFile.Seek(0, os.SEEK_CUR)
}

//获得文件末尾位置
func (this *FileHelper) GetFileEndPos() (int64, error) {
	return this.pFile.Seek(0, os.SEEK_END)
}

//获得文件长度
func (this *FileHelper) GetFileLength() int64 {
	n, _ := this.GetFileEndPos()
	return n + int64(1)
}

//读取多少个字节
func (this *FileHelper) ReadByte(nsize int) ([]byte, error) {
	bt := make([]byte, nsize)
	n, err := this.pFile.Read(bt)
	if err != nil {
		return nil, err
	}
	if n != nsize {
		return nil, fmt.Errorf("wish read not match real read")
	}
	return bt, nil
}

//读取转换为STRING
func (this *FileHelper) ReadString(nsize int) (string, error) {
	bt, err := this.ReadByte(nsize)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

//读取转换为STRING
func (this *FileHelper) ReadStringTrimSpace(nsize int) (string, error) {
	bt, err := this.ReadByte(nsize)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bt)), nil

}

//读取一个INT16数
func (this *FileHelper) ReadInt16() (uint16, error) {
	bt := make([]byte, 2)
	_, err := this.pFile.Read(bt)
	if err != nil {
		return 0, err
	}
	nval := binary.LittleEndian.Uint16(bt)
	return nval, nil
}

//读取一个INT32数
func (this *FileHelper) ReadInt32() (uint32, error) {
	bt := make([]byte, 4)
	_, err := this.pFile.Read(bt)
	if err != nil {
		return 0, err
	}

	nval := binary.LittleEndian.Uint32(bt)
	return nval, nil
}

//读取一个INT64数
func (this *FileHelper) ReadInt64() (uint64, error) {
	bt := make([]byte, 8)
	_, err := this.pFile.Read(bt)
	if err != nil {
		return 0, err
	}
	nval := binary.LittleEndian.Uint64(bt)
	return nval, nil
}

//查找数据,从文件的BEGIN位置开始查找S子串，第1次出现的位置,N大于0
func (this *FileHelper) Index(npos int64, sep []byte) int64 {
	//防止npos超出
	if _, err := this.pFile.Seek(int64(npos), os.SEEK_SET); err != nil {
		return -1
	}
	//取得截取的待搜索数据长度 20M
	var nMaxSize int = 20480
	lenSep := len(sep)
	if nMaxSize < lenSep*2 {
		nMaxSize = lenSep * 2
	}
	for {
		curPos, _ := this.GetFilePos()
		data, _ := this.ReadByte(nMaxSize)
		newPos := bytes.Index(data, sep)
		if len(data) < nMaxSize { // 数据太短，显然已经读取到最后一次
			if newPos == -1 {
				return -1
			} else {
				return curPos + int64(newPos)
			}
		} else { // 中间的相关数据
			if newPos >= 0 {
				return curPos + int64(newPos)
			}
		}
		this.Move(0 - len(sep))
	}
	return -1
}

//查找数据,从文件的BEGIN位置开始查找S子串，第N次出现的位置,N大于0
func (this *FileHelper) IndexN(npos int64, sep []byte, n int) int64 {
	if n <= 0 {
		return -1
	}
	//防止npos超出
	if _, err := this.pFile.Seek(int64(npos), os.SEEK_SET); err != nil {
		return -1
	}
	var (
		findPos int64
		err     error
	)

	for i := 0; i < n; i++ {
		curPos, _ := this.GetFilePos()
		findPos = this.Index(curPos, sep)
		if err != nil {
			return -1
		} else {
			//找到一次，那么要移动到这个位置加SEP长度
			this.MoveTo(int(findPos) + len(sep))
		}
	}
	return findPos
}
