package handler

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/jerbe/goms/client"
	"github.com/jerbe/goms/config"
	"github.com/jerbe/goms/data"
	"github.com/jerbe/goms/data/packet"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 14:19
  @describe :
*/

// handleLoginPassword 登录密码方法
func (h *Handler) handleLoginPassword(cli *client.Client, reader *data.LittleEndianReader) {
	username, err := reader.ReadMapleAsciiString()
	if err != nil {
		return
	}

	password, err := reader.ReadMapleAsciiString()
	if err != nil {
		return
	}

	// 读出mac地址
	macBytes, err := reader.Read(6)
	if err != nil {
		return
	}

	macBytesLen := len(macBytes)
	macAddr := ""
	for i := 0; i < macBytesLen; i++ {
		macAddr += fmt.Sprintf("%02X", macBytes[i])
		if i != macBytesLen-1 {
			macAddr += "-"
		}
	}

	// 设置该连接的Mac地址
	cli.SetMacAddress(macAddr)

	// 验证是否被IP限制
	//ipBaned := cli.IsBanedIP()
	// 验证是否被Mac地址限制
	//macBaned := cli.IsBanedMAC()

	//baned := ipBaned || macBaned

	// 判断服务是否打开全局自动注册
	if config.Default.Server.Register.Auto {
		// 验证是否有自动注册开关
	}

	// 验证账户密码是否正确

	log.Printf("IP:%v, MAC:%v, Username:%v, Password:%v", cli.GetRemoteAddress(), macAddr, username, password)

	cli.SetAccountID(1)
	cli.SetAccountName("admin")

	// 发送登录成功信息
	err = cli.SendMessage(packet.LoginSuccess(int(cli.GetAccountID()), 1, true, cli.GetAccountName()))
	if err != nil {
		log.Printf("handleLoginPassword error. reason:%v ", err)
	}

	servername := config.Default.Server.Name
	flag := config.Default.Server.Flag
	loginEventMessage := config.Default.Server.Login.Message.Event

	channelMap := make(map[int]int)
	for i, _ := range ChannelServerPool {
		channelMap[i] = 0
	}

	// 发送
	err = cli.SendMessage(packet.ServerList(0, flag, servername, loginEventMessage, channelMap))
	if err != nil {
		log.Printf("handleLoginPassword error. reason:%v ", err)
	}

	err = cli.SendMessage(packet.EndOfServerList())
	if err != nil {
		log.Printf("handleLoginPassword error. reason:%v ", err)
	}

}

// handleHelloLogin 登录时对服务端打招呼
func (h *Handler) handleHelloLogin(cli *client.Client, reader *data.LittleEndianReader) {

}

// handleServerStatusRequest 请求获取服务器状态列表
func (h *Handler) handleServerStatusRequest(cli *client.Client, reader *data.LittleEndianReader) {
	// 0 = 正常选择世界
	// 1 = "由于用户较多，您可能会遇到一些..."
	// 2 = "该世界的并发用户数已达到最大值"
	userOn := rand.Intn(100)
	userLimit := 100
	var err error
	if userOn >= userLimit {
		err = cli.SendMessage(packet.ServerStatus(2))
	} else if userOn*2 >= userLimit {
		err = cli.SendMessage(packet.ServerStatus(1))
	} else {
		err = cli.SendMessage(packet.ServerStatus(0))
	}

	if err != nil {
		log.Printf("handler.handleServerStatusRequest error. reason:%v", err)
	}
}

// handleCharListRequest 获取角色列表请求
func (h *Handler) handleCharListRequest(cli *client.Client, reader *data.LittleEndianReader) {
	worldID, err := reader.ReadByte()
	if err != nil {
		cli.Close()
		return
	}

	channelIdx, err := reader.ReadByte()
	if err != nil {
		cli.Close()
		return
	}
	channelIdx += 1
	if _, err = reader.ReadInt(); err != nil {
		cli.Close()
		return
	}

	cli.SetWorld(int(worldID))
	cli.SetChannel(int(channelIdx))
	charList := client.LoadCharacterList(int(worldID), cli.GetAccountID())

	err = cli.SendMessage(packet.CharacterList(charList, false, cli.GetCharacterSlots()))
	if err != nil {
		cli.Close()
		return
	}
}

// handleSelectCharacter 处理选中角色事件
func (h *Handler) handleSelectCharacter(cli *client.Client, reader *data.LittleEndianReader) {

	charID, err := reader.ReadInt()
	if err != nil {
		return
	}

	chanSvrIP := config.Default.Server.Address
	chanSvr := ChannelServerPool[cli.GetChannel()]
	port := chanSvr.Port()

	err = cli.SendMessage(packet.SendServerIP(chanSvrIP, port, charID))
	if err != nil {
		cli.Close()
		return
	}

}
