package file

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 文件读取对象
type FileHelper struct {
	readSeekerFile io.ReadSeeker
}

//OS文件打开
func (f *FileHelper) OpenOsFile(filepath string) error {
	var err error = nil
	if f.readSeekerFile, err = os.OpenFile(filepath, os.O_RDONLY, 0x666); err != nil {
		return err
	}
	return nil
}

//字节数组的打开
func (f *FileHelper) OpenBytes(p *bytes.Reader) error {
	f.readSeekerFile = p
	return nil
	//打开案例
	// m := &FileHelper{}
	// m.OpenBytes(bytes.NewReader([]byte{0xda}))
}

//字符串类型的打开
func (f *FileHelper) OpenStrings(p *strings.Reader) error {
	f.readSeekerFile = p
	return nil
	//打开案例
	// m := &FileHelper{}
	// m.OpenStrings(strings.NewReader("hello world"))
}

//关闭文件
func (f *FileHelper) Close() {
	if f.readSeekerFile != nil {
		if osFile, ok := f.readSeekerFile.(*os.File); ok {
			osFile.Close()
		}
		f.readSeekerFile = nil
	}
}

//移动到文件什么地方
func (f *FileHelper) MoveTo(npos int64) error {
	_, err := f.readSeekerFile.Seek(npos, os.SEEK_SET)
	if err != nil {
		return err
	}
	return nil
}

//移动到文件末尾
func (f *FileHelper) MoveToEnd() error {
	_, err := f.readSeekerFile.Seek(int64(0), os.SEEK_END)
	if err != nil {
		return err
	}
	return nil
}

//从当前位置移动多少个
func (f *FileHelper) Move(npos int64) error {
	_, err := f.readSeekerFile.Seek(npos, os.SEEK_CUR)
	if err != nil {
		return err
	}
	return nil
}

//获得当前位置
func (f *FileHelper) GetFilePos() (int64, error) {
	return f.readSeekerFile.Seek(0, os.SEEK_CUR)
}

//获得文件末尾位置
func (f *FileHelper) GetFileEndPos() (int64, error) {
	return f.readSeekerFile.Seek(0, os.SEEK_END)
}

//获得文件长度
func (f *FileHelper) GetFileLength() int64 {
	n, _ := f.GetFileEndPos()
	return n + int64(1)
}

//读取多少个字节
func (f *FileHelper) ReadByte(nsize int) ([]byte, error) {
	//nsize太大，可能会出现内存不足的问题，对于大于100M的内存分配，我们都预先检察一下
	if nsize > 102400 {
		currentPos, err := f.GetFilePos()
		if err != nil {
			return nil, err
		}
		if SurplusLen := f.GetFileLength() - currentPos; SurplusLen < int64(nsize) {
			return nil, fmt.Errorf(strconv.Itoa(nsize) + "is too long for this readSeekerFile,it has only " +
				strconv.Itoa(int(SurplusLen)) + " byte left,the current position is:" + strconv.Itoa(int(currentPos)))
		}
	}
	bt := make([]byte, nsize)
	n, err := f.readSeekerFile.Read(bt)
	if err != nil {
		return nil, err
	}
	if n != nsize {
		return nil, fmt.Errorf("wish read not match real read")
	}
	return bt, nil
}

//读取转换为STRING
func (f *FileHelper) ReadString(nsize int) (string, error) {
	bt, err := f.ReadByte(nsize)
	if err != nil {
		return "", err
	}
	return string(bt), nil
}

//读取转换为STRING
func (f *FileHelper) ReadStringTrimSpace(nsize int) (string, error) {
	bt, err := f.ReadByte(nsize)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bt)), nil

}

//读取一个INT16数
func (f *FileHelper) ReadInt16() (uint16, error) {
	bt := make([]byte, 2)
	_, err := f.readSeekerFile.Read(bt)
	if err != nil {
		return 0, err
	}
	nval := binary.LittleEndian.Uint16(bt)
	return nval, nil
}

//读取一个INT32数
func (f *FileHelper) ReadInt32() (uint32, error) {
	bt := make([]byte, 4)
	_, err := f.readSeekerFile.Read(bt)
	if err != nil {
		return 0, err
	}

	nval := binary.LittleEndian.Uint32(bt)
	return nval, nil
}

//读取一个INT64数
func (f *FileHelper) ReadInt64() (uint64, error) {
	bt := make([]byte, 8)
	_, err := f.readSeekerFile.Read(bt)
	if err != nil {
		return 0, err
	}
	nval := binary.LittleEndian.Uint64(bt)
	return nval, nil
}

//查找数据,从文件的npos位置开始查找sep子串，返回第1次出现的位置,npos大于等于0
func (f *FileHelper) Index(npos int64, sep []byte) int64 {
	//先算出最大长度
	endpos, _ := f.GetFileEndPos()
	//防止npos超出
	if _, err := f.readSeekerFile.Seek(npos, os.SEEK_SET); err != nil {
		return -1
	}
	//取得截取的待搜索数据长度 20M
	var nMaxSize int = 20480
	lenSep := len(sep)
	if nMaxSize < lenSep*2 {
		nMaxSize = lenSep * 2
	}

	for {
		var data []byte
		curPos, _ := f.GetFilePos()
		// 数据太短，显然已经读取到最后一次,最后一次则全部读取
		if int(endpos-curPos) < nMaxSize {
			data, _ = f.ReadByte(int(endpos - curPos))
		} else {
			data, _ = f.ReadByte(nMaxSize)
		}

		newPos := bytes.Index(data, sep)
		// 数据太短，显然已经读取到最后一次
		if len(data) < nMaxSize {
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
		f.Move(int64(0 - len(sep)))
	}
	return -1
}

//查找数据,从文件的npos位置开始查找S子串，第N次出现的位置,N大于0，npos大于等于0
func (f *FileHelper) IndexN(npos int64, sep []byte, n int) int64 {
	if n <= 0 {
		return -1
	}
	//防止npos超出
	if _, err := f.readSeekerFile.Seek(int64(npos), os.SEEK_SET); err != nil {
		return -1
	}
	var (
		findPos int64
		err     error
	)

	for i := 0; i < n; i++ {
		curPos, _ := f.GetFilePos()
		findPos = f.Index(curPos, sep)
		if err != nil {
			return -1
		} else {
			//找到一次，那么要移动到这个位置加SEP长度
			f.MoveTo(findPos + int64(len(sep)))
		}
	}
	return findPos
}
