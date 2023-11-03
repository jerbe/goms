package packet

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jerbe/goms/client"
	"github.com/jerbe/goms/data"
	"github.com/jerbe/goms/data/packet/code"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 10:08
  @describe :
*/

func SayHello(mapleVersion int16, sendIV, receiveIv []byte) *MaplePacket {
	writer := data.NewLittleEndianWriter().
		WriteShort(code.Opcode_HELLO.Int16()). // 国区? 13 = MSEA, 14 = GlobalMS, 15 = EMS
		WriteShort(mapleVersion).
		WriteZeroBytes(2).
		WriteBytes(receiveIv).
		WriteBytes(sendIV).
		WriteByte(4) // 7 = MSEA, 8 = GlobalMS, 5 = Test Server
	return NewMaplePacket(code.Opcode_HELLO, writer.Bytes())
}

// LoginSuccess 登录成功数据包
func LoginSuccess(accountID int, gender byte, isGM bool, accountName string) *MaplePacket {
	hexStr, _ := hex.DecodeString(strings.ReplaceAll("00 00 00 03 01 00 00 00 E2 ED A3 7A FA C9 01", " ", ""))

	gm := int16(0)
	if isGM {
		gm = 1
	}

	writer := data.NewLittleEndianWriter()

	writer.WriteShort(int16(code.Send.LOGIN_STATUS)).
		WriteByte(0).
		WriteInt(accountID).
		WriteByte(gender).
		WriteShort(gm).
		WriteMapleAsciiString(accountName).
		WriteBytes(hexStr).
		WriteInt(0).
		WriteLong(0).
		WriteMapleAsciiString(strconv.Itoa(accountID)).
		WriteMapleAsciiString(accountName).
		WriteByte(1)
	return NewMaplePacket(code.Send.LOGIN_STATUS, writer.Bytes())
}

// ServerStatus 服务状态数据包
func ServerStatus(status int) *MaplePacket {
	/*	 * 0 - Normal
	 * 1 - Highly populated
	 * 2 - Full*/
	writer := data.NewLittleEndianWriter().WriteShort(code.Send.SERVERSTATUS.Int16()).WriteShort(int16(status))
	return NewMaplePacket(code.Send.SERVERSTATUS, writer.Bytes())
}

// ServerList 获取服务列表的信息
// serverId 服务ID
// serverName 服务名
// channelLoad 频道负载
func ServerList(serverId, serverFlag int, serverName, longinEventMessage string, channelLoad map[int]int) *MaplePacket {
	writer := data.NewLittleEndianWriter()
	writer.
		WriteShort(code.Send.SERVERLIST.Int16()).
		WriteIntAsByte(serverId). // 0 = Aquilla, 1 = bootes, 2 = cass, 3 = delphinus
		WriteMapleAsciiString(serverName).
		WriteIntAsByte(serverFlag).
		WriteMapleAsciiString(longinEventMessage).
		WriteShort(100).
		WriteShort(100)

	lastChannel := 1
	maxChannel := 30

	for i := maxChannel; i > 0; i-- {
		if _, ok := channelLoad[i]; ok {
			lastChannel = i
			break
		}
	}

	writer.WriteIntAsByte(lastChannel).WriteInt(500)

	for i := 1; i <= lastChannel; i++ {
		load := 1200
		if l, ok := channelLoad[i]; ok {
			load = l
		}
		writer.WriteMapleAsciiString(fmt.Sprintf("%s-%d", serverName, i)).
			WriteIntAsByte(load).
			WriteInt(serverId).
			WriteShort(int16(i - 1))
	}
	writer.WriteShort(0)

	return NewMaplePacket(code.Send.SERVERLIST, writer.Bytes())
}

// EndOfServerList 获取服务列表停止信号
// serverId 服务ID
// serverName 服务名
// channelLoad 频道负载
func EndOfServerList() *MaplePacket {
	writer := data.NewLittleEndianWriter().
		WriteShort(code.Send.SERVERLIST.Int16()).
		WriteByte(0xFF)
	return NewMaplePacket(code.Send.SERVERLIST, writer.Bytes())
}

// CharacterList 返回角色列表
// secondPwd 是否启用二次密码验证
// slots 所有的角色卡槽数量
func CharacterList(charList []*client.Character, secondPwd bool, slots int) *MaplePacket {
	writer := data.NewLittleEndianWriter().
		WriteShort(code.Send.CHARLIST.Int16()).
		WriteByte(0).
		WriteInt(0).                  // 40 42 0F 00
		WriteIntAsByte(len(charList)) // 1?

	// 追加角色列表信息
	appendCharacterEntryList(writer, charList, false).
		WriteShort(3).
		WriteInt(slots)

	return NewMaplePacket(code.Send.CHARLIST, writer.Bytes())
}

func CharacterInfo(cli *client.Client, char *client.Character) *MaplePacket {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	writer := data.NewLittleEndianWriter()
	writer.WriteShort(code.Send.WARP_TO_MAP.Int16()).
		WriteInt(cli.GetChannel() - 1).
		Write(0).Write(1).Write(1).WriteShort(0).
		WriteInt(rnd.Int()).WriteInt(rnd.Int()).WriteInt(rnd.Int()) // 随机写入3个数字
	appendCharacterInfo(writer, char)

	return NewMaplePacket(code.Send.WARP_TO_MAP, writer.Bytes())
}
