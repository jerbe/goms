package packet

import (
	"encoding/binary"
	"github.com/jerbe/goms/data/packet/code"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 09:17
  @describe :
*/

func NewMaplePacket(opcode code.Opcoder, data []byte) *MaplePacket {
	return &MaplePacket{opcode: opcode, data: data}
}

// MaplePacket 枫叶传说专用的数据包
type MaplePacket struct {
	data []byte

	opcode code.Opcoder

	opcodeSet bool
}

// MarshalBinary 编码
func (p *MaplePacket) MarshalBinary() ([]byte, error) {
	return p.Bytes(), nil
}

// UnmarshalBinary 解码
func (p *MaplePacket) UnmarshalBinary(data []byte) error {
	p.SetData(data)
	return nil
}

// SetData 设置数据
func (p *MaplePacket) SetData(data []byte) {
	n := make([]byte, len(data))
	copy(n, data)
	p.data = n

	if !p.opcodeSet {
		p.opcode = code.Opcode(binary.LittleEndian.Uint16(data[0:2]))
	}
}

// SetOpcode 设置操作码
func (p *MaplePacket) SetOpcode(code code.Opcoder) {
	p.opcode = code
}

// Bytes 获取字节数据
func (p *MaplePacket) Bytes() []byte {
	n := make([]byte, len(p.data))
	copy(n, p.data)
	return n
}
