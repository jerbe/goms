package handler

import (
	"github.com/jerbe/goms/server"
	"log"

	"github.com/jerbe/goms/client"
	"github.com/jerbe/goms/data"
	"github.com/jerbe/goms/data/packet"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/11/2 09:15
  @describe :
*/

// handlePlayerLoggedin 处理角色选中并登入游戏
func (h *Handler) handlePlayerLoggedin(cli *client.Client, reader *data.LittleEndianReader) {
	characterID, err := reader.ReadInt()
	if err != nil {
		log.Printf("Handler.handlePlayerLoggedin read character id error. reson:%v", err)
		return
	}

	switch h.server.(type) {
	case server.IChannelServer:
		h.handlePlayerLoggedinChannelServer(cli, reader, characterID)
	case server.ICashShopServer:
		h.handlePlayerLoggedinCashShopServer(cli, reader, characterID)
	default:
		log.Printf("Handler.handlePlayerLoggedin not able loggined")
		cli.Close()
	}
}

// handlePlayerLoggedinChannelServer 处理玩家登录频道服务器
func (h *Handler) handlePlayerLoggedinChannelServer(cli *client.Client, reader *data.LittleEndianReader, characterID int) {
	server := h.server.(server.IChannelServer)
	character := client.NewCharacter(int64(characterID))
	err := character.LoadFromDB(false)
	if err != nil {
		log.Printf("Handler.handlePlayerLoggedinChannelServer error. reason:%v", err)
		return
	}

	cli.SetAccountID(int64(character.GetAccountID()))
	cli.SetChannel(server.Channel())

	// 移除所有频道内的该账户信息,因为有可能是账户意外退出

	// 判断用户是否有在其他地方登录,有的话直接断卡连接

	// 更新账户登录时间与地点. 为什么不在输入密码时就更新?

	// 往这个服务器注册角色
	server.RegisterCharacter(cli, character)

	// 往客户端发送角色详细信息
	err = cli.SendMessage(packet.CharacterInfo(cli, character))
	if err != nil {
		log.Printf("handlePlayerLoggedinChannelServer error. reason:%v ", err)
	}

	// 如果是GM,给角色附加某个技能效果

	// 临时能力值重置

	// 将角色加入到地图上

	// 沉默增益效果

	// 将将技能冷却处理

	// 沉默减益效果

	// 通知好友我上线了

	// 更新团队信息

	// 获取在线状态的好友,并给客户端发送上线好友的数据包

	// 聊天系统更新

	// 工会系统更新

	// 家族系统更新

	// 给角色发送技能宏

	// 显示交易消息 MTS(Maple Trading System)

	// 发送队伍里面其他队员的血条

	// 发送本人血条给其他队员

	// 开始Fairy计时

	// 修复玩家技能?

	// 发送键盘映射

	// 开始启动任务

	// 取得并移除最后一个好友请求

	// 判断角色是否是黑暗骑士

	//
}

// handlePlayerLoggedinCashShopServer 处理玩家登录现金商城服务器
func (h *Handler) handlePlayerLoggedinCashShopServer(cli *client.Client, reader *data.LittleEndianReader, characterID int) {

}
