package data

import (
	"bytes"
	"encoding/binary"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/25 11:29
  @describe :
*/

type LittleEndianWriter struct {
	writer *bytes.Buffer
	err    error
}

// NewLittleEndianWriter 返回一个小端数据访问器
func NewLittleEndianWriter() *LittleEndianWriter {
	return &LittleEndianWriter{
		writer: &bytes.Buffer{},
	}
}

// WriteZeroBytes 写入count 个空字符
func (w *LittleEndianWriter) WriteZeroBytes(count int) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	buf := make([]byte, count)
	return w.WriteBytes(buf)
}

// WriteBytes 写入指定字节流
func (w *LittleEndianWriter) WriteBytes(in []byte) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	_, w.err = w.writer.Write(in)
	return w
}

// Write 写进一个字节
func (w *LittleEndianWriter) Write(in byte) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, in)
	return w
}

// WriteByte 写进一个字节
func (w *LittleEndianWriter) WriteByte(in byte) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, in)
	return w
}

// WriteInt 写进一个整形,4个字节
func (w *LittleEndianWriter) WriteInt(in int) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, uint32(in))
	return w
}

// WriteIntAsByte 写入一个整形
func (w *LittleEndianWriter) WriteIntAsByte(in int) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, byte(in))
	return w
}

// WriteShort 写入一个短整形 init16
func (w *LittleEndianWriter) WriteShort(in int16) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, in)
	return w
}

// WriteChar 写入一个字符 4字节
func (w *LittleEndianWriter) WriteChar(in rune) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, in)
	return w
}

// WriteLong 写入一个长整形 8字节
func (w *LittleEndianWriter) WriteLong(in int64) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, in)
	return w
}

// WriteFloat 写入一个浮点型 4字节
func (w *LittleEndianWriter) WriteFloat(in float32) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, in)
	return w
}

// WriteDouble 写入双精度浮点型 8字节
func (w *LittleEndianWriter) WriteDouble(in float64) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, in)
	return w
}

// WriteAsciiString 写入一串字符串
func (w *LittleEndianWriter) WriteAsciiString(in string, limit ...int) *LittleEndianWriter {
	if w.err != nil {
		return w
	}
	bgkEncode := simplifiedchinese.GBK.NewEncoder()
	byteData, _, _ := transform.Bytes(bgkEncode, []byte(in))

	if len(limit) > 0 {
		nBytes := make([]byte, limit[0], limit[0])
		copy(nBytes, byteData)
		byteData = nBytes
	}

	w.err = binary.Write(w.writer, binary.LittleEndian, byteData)
	return w
}

// WriteMapleAsciiString 写入枫叶传说专用字符串
// 会自动写入长度
func (w *LittleEndianWriter) WriteMapleAsciiString(in string) *LittleEndianWriter {
	if w.err != nil {
		return w
	}

	bgkEncode := simplifiedchinese.GBK.NewEncoder()
	byteData, _, _ := transform.Bytes(bgkEncode, []byte(in))

	if w.WriteShort(int16(len(byteData))).Err() != nil {
		return w
	}
	w.err = binary.Write(w.writer, binary.LittleEndian, byteData)
	return w
}

// WritePos 写入坐标点
func (w *LittleEndianWriter) WritePos(x, y int) *LittleEndianWriter {
	if w.err != nil {
		return w
	}

	if w.WriteShort(int16(x)).Err() != nil {
		return w
	}

	if w.WriteShort(int16(y)).Err() != nil {
		return w
	}
	return w
}

// Bytes 返回已写入的数据
func (w *LittleEndianWriter) Bytes() []byte {
	return w.writer.Bytes()
}

func (w *LittleEndianWriter) Err() error {
	return w.err
}
