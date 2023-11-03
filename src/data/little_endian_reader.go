package data

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/25 11:29
  @describe :
*/

type LittleEndianReader struct {
	reader *bytes.Reader
}

// NewLittleEndianReader 返回一个小端数据访问器
func NewLittleEndianReader(data []byte) *LittleEndianReader {
	return &LittleEndianReader{
		reader: bytes.NewReader(data),
	}
}

// NewLittleEndianReaderFromReader 根据一个读取者返回一个小端数据访问器
func NewLittleEndianReaderFromReader(in io.Reader) (*LittleEndianReader, error) {
	buffer, ok := in.(*bytes.Reader)
	if !ok {
		data, err := io.ReadAll(in)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewReader(data)
	}
	return &LittleEndianReader{
		reader: buffer,
	}, nil
}

// ReadByteAsInt 读取一个字节并转换成整形
func (g *LittleEndianReader) ReadByteAsInt() (int, error) {
	d, err := g.ReadByte()
	if err != nil {
		return 0, err
	}
	return int(d), nil
}

// ReadByte 读取一个字节
func (g *LittleEndianReader) ReadByte() (byte, error) {
	b, err := g.reader.ReadByte()
	if err != nil {
		return 0, err
	}
	return b, nil
}

// ReadInt 读取一个整形
func (g *LittleEndianReader) ReadInt() (int, error) {
	var data [4]byte
	_, err := io.ReadFull(g.reader, data[:])
	if err != nil {
		return 0, err
	}
	return int(binary.LittleEndian.Uint32(data[:])), nil
}

// ReadShort 读取一个短整形 init16
func (g *LittleEndianReader) ReadShort() (int16, error) {
	var data [2]byte
	_, err := io.ReadFull(g.reader, data[:])
	if err != nil {
		return 0, err
	}
	return int16(binary.LittleEndian.Uint16(data[:])), nil
}

// ReadChar 读取一个字符 4字节
// 存疑
func (g *LittleEndianReader) ReadChar() (rune, error) {
	//d, err := g.ReadShort()
	//return rune(d), err

	var data [4]byte
	_, err := io.ReadFull(g.reader, data[:])
	if err != nil {
		return 0, err
	}
	return rune(binary.LittleEndian.Uint32(data[:])), nil
}

// ReadLong 读取一个长整形 8字节
func (g *LittleEndianReader) ReadLong() (int64, error) {
	var data [8]byte
	_, err := io.ReadFull(g.reader, data[:])
	if err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(data[:])), nil
}

// ReadFloat 读取一个浮点型 4字节
func (g *LittleEndianReader) ReadFloat() (float32, error) {
	d, err := g.ReadInt()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(uint32(d)), nil
}

// ReadDouble 读取双精度浮点型 8字节
func (g *LittleEndianReader) ReadDouble() (float64, error) {
	d, err := g.ReadLong()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(uint64(d)), nil
}

// ReadAsciiString 读取一串字符串
func (g *LittleEndianReader) ReadAsciiString(n int) (string, error) {
	data := make([]byte, n)
	_, err := io.ReadFull(g.reader, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadMapleAsciiString 读取枫叶传说专用字符串
// 会自动读取长度并返回字符串
func (g *LittleEndianReader) ReadMapleAsciiString() (string, error) {
	length, err := g.ReadShort()
	if err != nil {
		return "", err
	}
	return g.ReadAsciiString(int(length))
}

// ReadPos 读取坐标点
func (g *LittleEndianReader) ReadPos() (int, int, error) {
	x, err := g.ReadShort()
	if err != nil {
		return 0, 0, err
	}
	y, err := g.ReadShort()
	if err != nil {
		return 0, 0, err
	}
	return int(x), int(y), nil
}

// Read 读取指定num个字节
func (g *LittleEndianReader) Read(num int) ([]byte, error) {
	data := make([]byte, num)
	_, err := io.ReadFull(g.reader, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Seek 将读取位置移动到num的位置上
func (g *LittleEndianReader) Seek(num int) error {
	_, err := g.reader.Seek(int64(num), io.SeekStart)
	return err
}

// Skip 跳过num个字节不读取
func (g *LittleEndianReader) Skip(num int) error {
	_, err := g.reader.Seek(int64(num), io.SeekCurrent)
	return err
}
