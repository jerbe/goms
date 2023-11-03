package server

import (
	"context"
	"fmt"
	"github.com/jerbe/goms/constants"
	"io"
	"log"
	"net"
	"sync"

	"github.com/jerbe/goms/client"
	"github.com/jerbe/goms/config"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/25 15:37
  @describe :
*/

type IServer interface {
	// ServerName 服务名称
	ServerName() string

	// Port 服务端口号
	Port() int

	// Startup 启动服务,但是保持服务不在线状态
	Startup() error

	// SetOn 设置服务为上线状态
	SetOn() error

	// Shutdown 关闭
	Shutdown(context.Context) error

	// IsShutdown 判断服务是否已经关闭
	IsShutdown() bool

	// Flag 服务器标识
	Flag() int

	// SetFlag 设置服务器标识
	SetFlag(int)

	// Context 返回服务的上下文
	Context() context.Context

	// onAddClient 当客户端添加时
	onAddClient(*client.Client)

	//onRemoveClient 当客户端被移除时
	onRemoveClient(*client.Client)
}

// ------------ BaseServer
var _ IServer = new(BaseServer)

func NewBaseServer(ctx context.Context, port int, name string) *BaseServer {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)
	return &BaseServer{
		port:       port,
		name:       name,
		ctx:        ctx,
		flag:       0,
		cancelFunc: cancel,
		clients:    make(map[*client.Client]struct{}),
		shutdown:   true,
	}
}

// BaseServer 基础的服务
type BaseServer struct {
	// port 服务监听端口号
	port int

	// name 服务器名称
	name string

	// ctx 上下文
	ctx context.Context

	// cancelFunc 取消方法,也就是关闭
	cancelFunc context.CancelFunc

	// 服务标识
	flag int

	// 监听器
	listener *net.TCPListener

	// 客户列表
	clients map[*client.Client]struct{}

	// 读写锁
	rwMutex sync.RWMutex

	// shutdown 判断服务是否处于关闭状态
	shutdown bool

	// OnClientConnect 当客户端连接到服务端时,调用
	OnClientConnect func(*client.Client) error

	// OnClientReceiveMessage 客户端连接时调用的方法
	OnClientReceiveMessage func(*client.Client, io.Reader)
}

// Accept 监听并启动服务
func (s *BaseServer) accept() {
	defer func() {
		s.shutdown = true
	}()

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
		}

		if s.shutdown {
			continue
		}

		conn, err := s.listener.AcceptTCP()
		if err != nil {
			log.Printf("BaseServer accept 's.listener.Accept' error. reason:%v", err)
			continue
		}

		log.Printf("BaseServer accept. local address:%s, remote address:%s", conn.LocalAddr(), conn.RemoteAddr())

		conn.SetNoDelay(true)
		cli := client.NewClient(s.Context(), conn, constants.OFBKey, constants.MapleVersion)
		s.onAddClient(cli)

	}
}

// onAddClient 增加一个客户端
func (s *BaseServer) onAddClient(cli *client.Client) {
	log.Println("BaseServer onAddClient")
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	s.clients[cli] = struct{}{}
	cli.OnClosed = s.onRemoveClient
	cli.OnConnect = s.OnClientConnect
	cli.OnReceiveMessage = s.OnClientReceiveMessage
	cli.Handle()
}

// onRemoveClient 移除客户端
func (s *BaseServer) onRemoveClient(cli *client.Client) {
	log.Println("BaseServer client closed")
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	delete(s.clients, cli)
}

// Context 返回服务的上下文
func (s *BaseServer) Context() context.Context {
	return s.ctx
}

// ServerName 服务名称
func (s *BaseServer) ServerName() string {
	return s.name
}

// Port 服务端监听端口号
func (s *BaseServer) Port() int {
	return s.port
}

// Startup 配置服务
func (s *BaseServer) Startup() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.Default.Server.Address, s.port))
	if err != nil {
		return err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return err
	}

	s.listener = listener
	go s.accept()
	return nil
}

// SetOn 设置服务为上线状态
func (s *BaseServer) SetOn() error {
	s.shutdown = false
	return nil
}

// Shutdown 关闭服务
func (s *BaseServer) Shutdown(ctx context.Context) error {
	s.cancelFunc()
	s.rwMutex.Lock()
	for cli := range s.clients {
		cli.OnClosed = nil
		cli.Close()
	}

	s.clients = nil
	s.rwMutex.Unlock()
	return s.listener.Close()
}

// IsShutdown 服务是否已经关闭
func (s *BaseServer) IsShutdown() bool {
	select {
	case <-s.ctx.Done():
		return true
	default:
		return s.shutdown
	}
}

// Flag 标识,不知道干嘛用的
func (s *BaseServer) Flag() int {
	return s.flag
}

// SetFlag 设置标识
func (s *BaseServer) SetFlag(i int) {
	s.flag = i
}
