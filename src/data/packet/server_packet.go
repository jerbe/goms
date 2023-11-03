package packet

import (
	"net"

	"github.com/jerbe/goms/data"
	"github.com/jerbe/goms/data/packet/code"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 15:58
  @describe :
*/

// SendServerIP 发送服务器IP
// serverIP 服务IP
// port 服务端端口
// characterID 角色id
func SendServerIP(serverIP string, port int, characterID int) *MaplePacket {
	writer := data.NewLittleEndianWriter().
		WriteShort(code.Send.SERVER_IP.Int16()).
		WriteShort(0).
		WriteBytes(net.ParseIP(serverIP).To4()). // 写入IP
		WriteShort(int16(port)).
		WriteInt(characterID).
		WriteBytes([]byte{1, 0, 0, 0, 0})
	return NewMaplePacket(code.Send.SERVER_IP, writer.Bytes())
}
