package codec

import (
	"bufio"
	"bytes"
	"encoding"
	"encoding/binary"
	"errors"
	"github.com/jerbe/goms/config"
	"github.com/jerbe/goms/crypt"
	"github.com/jerbe/goms/utils"
	"io"
	"strconv"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/26 11:27
  @describe :
*/

const (
	EncodeCrypt = "encode_crypt"
	DecodeCrypt = "decode_crypt"
)

var (
	InvalidPacket = errors.New("MapleDecoder: invalid packet")
)

//---------------------- MapleEncoder

// NewMapleEncoder 返回一个枫叶传说加密器
func NewMapleEncoder(cypher *crypt.MapleCypher) *MapleEncoder {
	return &MapleEncoder{
		cypher: cypher,
	}
}

// MapleEncoder 枫叶传说专用编码器
type MapleEncoder struct {
	UnimplementedProtocolEncoder
	cypher *crypt.MapleCypher
}

// Encode 编码
func (e *MapleEncoder) Encode(input any, output io.Writer) error {
	toBytes, err := utils.InputToBytes(input)
	if err != nil {
		return InvalidEncodeInput
	}
	// @TODO debug
	unencrypted := make([]byte, len(toBytes))
	copy(unencrypted, toBytes)
	encrypted := make([]byte, len(toBytes)+4)

	header := e.cypher.PacketHeader(len(unencrypted))
	crypt.MapleEncrypt(unencrypted)

	e.cypher.Crypt(unencrypted, unencrypted)
	copy(encrypted, header)
	copy(encrypted[4:], unencrypted)
	_, err = output.Write(encrypted)
	return err
}

//---------------------- MapleDecoder

// NewMapleDecoder 返回一个枫叶传说的解码器
func NewMapleDecoder(cypher *crypt.MapleCypher) *MapleDecoder {
	return &MapleDecoder{
		cypher: cypher,
	}
}

// MapleDecoder 枫叶传说专用编码器
type MapleDecoder struct {
	UnimplementedProtocolDecoder
	cypher *crypt.MapleCypher
}

func bytesToOutput(src []byte, output any) error {
	var (
		err error
	)
	switch data := output.(type) {
	case *[]byte:
		*data = src
	case *string:
		*data = string(src)
	case *int:
		*data, err = strconv.Atoi(string(src))
	case *int8:
		n, err := strconv.ParseInt(string(src), 10, 8)
		if err != nil {
			break
		}
		*data = int8(n)
	case *int16:
		n, err := strconv.ParseInt(string(src), 10, 16)
		if err != nil {
			break
		}
		*data = int16(n)
	case *int32:
		n, err := strconv.ParseInt(string(src), 10, 32)
		if err != nil {
			break
		}
		*data = int32(n)
	case *int64:
		*data, err = strconv.ParseInt(string(src), 10, 64)
	case *uint:
		n, err := strconv.ParseUint(string(src), 10, strconv.IntSize)
		if err != nil {
			break
		}
		*data = uint(n)
	case *uint8:
		n, err := strconv.ParseUint(string(src), 10, 8)
		if err != nil {
			break
		}
		*data = uint8(n)
	case *uint16:
		n, err := strconv.ParseUint(string(src), 10, 16)
		if err != nil {
			break
		}
		*data = uint16(n)
	case *uint32:
		n, err := strconv.ParseUint(string(src), 10, 32)
		if err != nil {
			break
		}
		*data = uint32(n)
	case *uint64:
		*data, err = strconv.ParseUint(string(src), 10, 64)
	case *float32:
		n, err := strconv.ParseFloat(string(src), 32)
		if err != nil {
			break
		}
		*data = float32(n)
	case *float64:
		*data, err = strconv.ParseFloat(string(src), 64)
	case io.Writer:
		_, err = data.Write(src)
	case encoding.BinaryUnmarshaler:
		err = data.UnmarshalBinary(src)
	default:
		err = InvalidDecoderOutput
	}
	return err
}

// Decode 解码
func (e *MapleDecoder) Decode(input io.Reader, output io.Writer) error {

	const headerLength = 4
	reader := bufio.NewReader(input)
	packetLength := -1

	for {
		// 判断是否大于头部信息
		if packetLength == -1 {
			// 读出头部信息
			// 这里不会改变偏移,所以不用担心
			peek, err := reader.Peek(4)
			if err != nil {
				return err
				//if errors.Is(err, io.EOF) {
				//	return err
				//}
				// 如果头部不够长就继续等待
				//continue
			}

			// 读出头部信息
			buffer := bytes.NewBuffer(peek)
			var header uint32
			err = binary.Read(buffer, binary.BigEndian, &header)
			if err != nil {
				continue
			}

			// 不是一个合法的包信息
			if !e.cypher.CheckPacketHeader(int(header)) {
				return InvalidPacket
			}

			packetLength = e.cypher.PacketLength(int(header))
			reader.Discard(4)
		}

		if reader.Buffered() >= packetLength && packetLength != -1 {
			// 根据包长度读取包内容
			decryptedPacket := make([]byte, packetLength)
			//_, err := reader.Read(decryptedPacket)
			err := binary.Read(reader, binary.BigEndian, &decryptedPacket)
			packetLength = -1
			if err != nil {
				return err
			}

			// 解密出网络数据
			err = e.cypher.Crypt(decryptedPacket, decryptedPacket)
			if err != nil {
				return err
			}

			// 解密出内容数据
			crypt.MapleDecrypt(decryptedPacket)

			_, err = output.Write(decryptedPacket)

			// 如果开启了被数据调试并且显示包内容
			if config.Default.Server.Logger.Packet.Debug && config.Default.Server.Logger.Packet.Display {
				//@TODO logging
			}
			return err
		}
	}
}

// ----------------------
var _ ProtocolCodecer = new(MapleCodec)

// NewMapleCodec 返回一个枫叶传说编码解码器
func NewMapleCodec(encodeCypher *crypt.MapleCypher, decodeCypher *crypt.MapleCypher) *MapleCodec {
	encoder := NewMapleEncoder(encodeCypher)
	decoder := NewMapleDecoder(decodeCypher)
	return &MapleCodec{
		encoder: encoder,
		decoder: decoder,
	}
}

// MapleCodec 枫叶传说专用编码解码器
type MapleCodec struct {
	encoder *MapleEncoder
	decoder *MapleDecoder
}

// Encode 编码
func (c *MapleCodec) Encode(input any, output io.Writer) error {
	return c.encoder.Encode(input, output)
}

// Decode 解码
func (c *MapleCodec) Decode(input io.Reader, output io.Writer) error {
	return c.decoder.Decode(input, output)
}
