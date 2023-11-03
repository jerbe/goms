package server

import (
	"context"
	"log"

	"github.com/jerbe/goms/config"
	"github.com/jerbe/goms/utils"
	"github.com/jerbe/goms/utils/container"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/22 21:02
  @describe :
*/

type ILoginServer interface {
	IServer

	// AddChannel 添加频道
	AddChannel(int)

	// RemoveChannel 移除频道
	RemoveChannel(int)

	// EventMessage 事件消息
	EventMessage() string

	// MaxCharacters 最多角色
	MaxCharacters() int

	// UserLogin 累加用户登录次数
	UserLogin(channelId int)

	// UserLimit 用户数量限制
	UserLimit() int

	// SetUserLimit 设置用户数量限制
	SetUserLimit(int)

	// AdminOnly 只允许管理员登录
	AdminOnly() bool
}

var LoginIPAuth = container.NewSet[string]()
var LoginAuth = make(map[int64]utils.Triple[string, string, int])

var _ ILoginServer = &LoginServer{}

// NewLoginServer 返回登陆服务器
func NewLoginServer(ctx context.Context, config config.ServerConfig) *LoginServer {
	svr := &LoginServer{
		BaseServer:    *NewBaseServer(ctx, config.Login.Port, config.Name),
		maxCharacters: config.Limit.Characters,
		adminOnly:     config.Login.Admin,
		userLimit:     config.Limit.Online,
		eventMessage:  config.Login.Message.Event,
		usersOn:       0,
	}

	svr.SetFlag(config.Flag)
	return svr
}

// LoginServer 登录服务器
type LoginServer struct {
	BaseServer

	load map[int]int

	// maxCharacters 最多允许角色数量
	maxCharacters int

	// adminOnly 只允许管理员登陆
	adminOnly bool

	// userLimit 用户数量限制
	userLimit int

	// eventMessage 事件消息,登录后的事件消息?
	eventMessage string

	// usersOn 已登录用户数
	usersOn int
}

// Startup 配置并启动
func (s *LoginServer) Startup() error {
	err := s.BaseServer.Startup()
	if err == nil {
		log.Printf("LoginServer startup and listen on %d", s.port)
	}
	return err
}

// AddChannel 添加频道
func (s *LoginServer) AddChannel(i int) {
	//TODO implement me
	panic("implement me")
}

// RemoveChannel 移除频道
func (s *LoginServer) RemoveChannel(i int) {
	//TODO implement me
	panic("implement me")
}

// EventMessage 事件消息
func (s *LoginServer) EventMessage() string {
	//TODO implement me
	panic("implement me")
}

// MaxCharacters 最多角色
func (s *LoginServer) MaxCharacters() int {
	//TODO implement me
	panic("implement me")
}

// UserLogin 累加用户登录次数
func (s *LoginServer) UserLogin(channelId int) {
	//TODO implement me
	panic("implement me")
}

// UserLimit 用户数量限制
func (s *LoginServer) UserLimit() int {
	//TODO implement me
	panic("implement me")
}

// SetUserLimit 设置用户数量限制
func (s *LoginServer) SetUserLimit(i int) {
	//TODO implement me
	panic("implement me")
}

// AdminOnly 只允许管理员登录
func (s *LoginServer) AdminOnly() bool {
	//TODO implement me
	panic("implement me")
}
