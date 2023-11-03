package server

import (
	"context"
	"log"

	"github.com/jerbe/goms/client"
	"github.com/jerbe/goms/config"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/22 21:02
  @describe :
*/

// IChannelServer 频道服务器接口
type IChannelServer interface {
	IServer

	// LoginMessage 登录时给客户端发送的消息
	LoginMessage() string

	// AdminOnly 只允许管理员登录
	AdminOnly() bool

	// Channel 返回频道编号
	Channel() int

	// RegisterCharacter 注册角色
	RegisterCharacter(cli *client.Client, character *client.Character)

	// UnregisterCharacter 取消注册角色
	UnregisterCharacter(character *client.Character)
}

var _ IChannelServer = new(ChannelServer)

// NewChannelServer 返回频道服务器
func NewChannelServer(ctx context.Context, channel int) *ChannelServer {
	cfg := config.Default.Server
	svr := &ChannelServer{
		BaseServer: *NewBaseServer(ctx, cfg.Channel.Port+channel-1, cfg.Name),
		channel:    channel,
	}

	svr.SetFlag(cfg.World.Flags)
	svr.adminOnly = cfg.Login.Admin
	return svr
}

// ChannelServer 频道服务器
type ChannelServer struct {
	BaseServer
	// 基础配置

	// channel 当前频道数值
	channel int

	// expRate 经验倍率. 这个参数影响玩家角色从击败怪物或完成任务等活动中获得的经验值。如果服务器的经验值倍率为2x，那么玩家获得的经验值将翻倍，加快升级速度。
	expRate int

	// mesoRate 金币倍率.金币倍率会影响玩家从击败怪物或其他金币获取途径获得的金币数量。如果服务器的金币倍率为2x，那么玩家将获得两倍的金币。
	mesoRate int

	// dropRate 物品掉落倍率.这个参数影响怪物掉落物品的几率和数量。如果服务器的物品掉落倍率为2x，那么怪物掉落的物品将翻倍。
	dropRate int

	// cashRate 现金卡券商城倍率.现金商城倍率通常影响购买商城物品所需的现金点数。如果服务器的现金商城倍率为2x，那么购买商城物品所需的现金点数将减半。
	cashRate int

	// bossDropRate BOSS怪物掉落倍率.这个参数通常影响从BOSS怪物身上获得的掉落物品的几率和数量。如果服务器的BOSS掉落倍率为2x，那么BOSS怪物的掉落物品将翻倍。
	bossDropRate int

	// loginMessage 登陆时给客户端发送的消息
	loginMessage string

	// adminOnly 只允许管理员登录
	adminOnly bool

	// megaphoneMute 扩音器广播是否关闭
	megaphoneMute bool

	// mapFactory 地图
	mapFactory any

	// characterContainer 角色容器
	characterContainer *client.CharacterContainer
}

// Channel 返回通道
func (s *ChannelServer) Channel() int {
	return s.channel
}

// RegisterCharacter 注册角色
func (s *ChannelServer) RegisterCharacter(cli *client.Client, character *client.Character) {
	s.characterContainer.Register(cli, character)
}

// UnregisterCharacter 取消注册角色
func (s *ChannelServer) UnregisterCharacter(character *client.Character) {
	s.characterContainer.Unregister(character)
}

// LoginMessage 登录时发送的消息
func (s *ChannelServer) LoginMessage() string {
	return s.loginMessage
}

// AdminOnly 是否只允许管理员登录
func (s *ChannelServer) AdminOnly() bool {
	return s.adminOnly
}

// Startup 配置并启动服务,但未设置上线
func (s *ChannelServer) Startup() error {
	err := s.BaseServer.Startup()
	if err == nil {
		log.Printf("ChannelServer startup and listen on %d", s.port)
	}
	return err
}
