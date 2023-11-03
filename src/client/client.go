package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"time"

	"github.com/jerbe/goms/codec"
	"github.com/jerbe/goms/crypt"
	"github.com/jerbe/goms/database"
	"github.com/jerbe/goms/utils"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/26 12:47
  @describe :
*/

const (
	// ReadIdleDuration 读空闲判断时间间隔
	ReadIdleDuration = time.Second * 10

	// WriteIdleDuration 写空闲判断时间间隔
	WriteIdleDuration = time.Second * 10
)

type IdleStatus uint8

const (
	// IdleStatusBothIdle 读写都空闲
	IdleStatusBothIdle IdleStatus = iota

	// IdleStatusWriteIdle 写空闲
	IdleStatusWriteIdle

	// IdleStatusReadIdle 读空闲
	IdleStatusReadIdle
)

// NewClient 返回枫叶传说客户端
func NewClient(ctx context.Context, conn net.Conn, ofbKey []byte, version uint16) *Client {
	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithCancel(ctx)

	// 编码协议
	encodeCypher, err := crypt.NewMapleCypher(ofbKey, []byte{1, 0x5F, 4, 0x3F}, 0xFFFF-version)
	if err != nil {
		panic(err)
	}

	// 解码协议
	decodeCypher, err := crypt.NewMapleCypher(ofbKey, []byte{9, 0, 0x5, 0x5F}, version)
	if err != nil {
		panic(err)
	}

	return &Client{
		conn:            conn,
		codec:           codec.NewMapleCodec(encodeCypher, decodeCypher),
		ctx:             ctx,
		cancelFunc:      cancel,
		idleStatus:      IdleStatusBothIdle,
		writeIdleTicker: time.NewTimer(WriteIdleDuration),
		handled:         false,
		encodeCypher:    encodeCypher,
		decodeCypher:    decodeCypher,
	}
}

// Client 枫叶传说客户端
type Client struct {
	// conn 网络连接
	conn net.Conn

	// codec 协议编码解码器
	codec codec.ProtocolCodecer

	// ctx 上下文
	ctx context.Context

	// cancelFunc 取消上下文方法
	cancelFunc context.CancelFunc

	// idleStatus 空闲状态
	idleStatus IdleStatus

	// writeIdleTicker 写空闲定时器
	writeIdleTicker *time.Timer

	// handled 已经在处理
	handled bool

	// encodeCypher 编码暗号器
	encodeCypher *crypt.MapleCypher

	// decodeCypher 解码暗号器
	decodeCypher *crypt.MapleCypher

	// OnConnect 当连接时调度
	OnConnect func(*Client) error

	// OnClosed 关闭后执行方法
	OnClosed func(*Client)

	// OnIdle 连接空闲时执行的方法
	OnIdle func(*Client, IdleStatus)

	// OnReceiveMessage 收到消息的执行方法
	OnReceiveMessage func(*Client, io.Reader)

	// OnWrittenMessage 写消息事件
	OnWrittenMessage func(*Client, any)

	/// ------------------
	// accountID 账户ID
	accountID int64

	// accountName 账户名
	accountName string

	// world 登录的世界
	world int

	// channel 登录的频道
	channel int

	// isGM 是否是GM账户
	isGM bool

	// characterSlots 当前账户允许的总卡槽数量,卡槽数就是总的可以容纳的角色数量
	characterSlots int
}

// Handle 开始处理
func (c *Client) Handle() {
	if c.conn == nil || c.handled {
		return
	}

	c.handled = true

	// 进行写空闲判断
	if c.writeIdleTicker == nil {
		c.writeIdleTicker = time.NewTimer(WriteIdleDuration)
	}

	// 写空闲计时
	go c.writeIdleTick()

	// 获取数据
	go c.receive()

	if c.OnConnect != nil {
		err := c.OnConnect(c)
		// 发给客户端数据Hello数据包
		//err := c.SendMessageWithoutEncode(packet.SayHello(int16(constants.MapleVersion), c.encodeCypher.OriginalIv(), c.decodeCypher.OriginalIv()))
		if err != nil {
			if errors.Is(err, io.EOF) {
				c.Close()
			}
			log.Printf("Client OnConnect error. reason:%v", err)
		}
	}
}

// writeIdleTick 写空闲间隔计数
func (c *Client) writeIdleTick() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-c.writeIdleTicker.C:
			c.idleStatus &^= IdleStatusReadIdle
			go c.handleIdle(IdleStatusReadIdle)
			c.resetWriteIdleTick()
		}
	}
}

// receive 处理
func (c *Client) receive() {
	defer func() {
		if obj := recover(); obj != nil {
			log.Println("Client receive panic")
		}
	}()
	defer func() {
		c.cancelFunc()
		if c.OnClosed != nil {
			c.OnClosed(c)
		}
	}()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		//err := c.conn.SetReadDeadline(time.Now().Add(ReadIdleDuration))
		buffer := &bytes.Buffer{}
		err := c.codec.Decode(c.conn, buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				log.Printf("Client: receive() codec.Decode error. reason: %v ", err)
				return
			}

			var netErr net.Error
			if errors.As(err, &netErr) && netErr.Timeout() {
				c.idleStatus &^= IdleStatusReadIdle
				go c.handleIdle(c.idleStatus)
				continue
			}

			log.Printf("Client: receive() codec.Decode error. reason: %v ", err)
			continue
		}

		// 接收消息处理
		go c.receiveHandle(buffer)
	}
}

// resetWriteIdleTick 重置写空间
func (c *Client) resetWriteIdleTick() {
	c.writeIdleTicker.Reset(WriteIdleDuration)
}

// handleIdle 处理空闲状态
func (c *Client) handleIdle(status IdleStatus) {
	defer func() {
		if obj := recover(); obj != nil {
			log.Printf("Client handleIdle panic. %v", obj)
		}
	}()

	select {
	case <-c.ctx.Done():
		return
	default:
	}

	if c.OnIdle != nil {
		c.OnIdle(c, status)
	}
}

// Close 关闭客户端
func (c *Client) Close() error {
	err := c.conn.Close()
	c.cancelFunc()
	return err
}

// SendMessage 发送消息
func (c *Client) SendMessage(obj any) error {
	if ok, err := c.IsClose(); ok {
		return err
	}
	c.idleStatus |= IdleStatusWriteIdle
	c.resetWriteIdleTick()

	err := c.codec.Encode(obj, c.conn)
	if errors.Is(err, io.EOF) {
		c.Close()
		return err
	}

	if err == nil && c.OnWrittenMessage != nil {
		c.OnWrittenMessage(c, obj)
	}
	return err
}

// SendMessageWithoutEncode 发送不编码的数据
func (c *Client) SendMessageWithoutEncode(obj any) error {
	if ok, err := c.IsClose(); ok {
		return err
	}
	c.idleStatus |= IdleStatusWriteIdle
	c.resetWriteIdleTick()

	data, err := utils.InputToBytes(obj)
	if err != nil {
		return err
	}

	_, err = c.conn.Write(data)
	if errors.Is(err, io.EOF) {
		c.Close()
		return err
	}

	if err == nil && c.OnWrittenMessage != nil {
		c.OnWrittenMessage(c, obj)
	}
	return err
}

// receiveHandle 接收到消息的处理方法
func (c *Client) receiveHandle(input io.Reader) {
	defer func() {
		if obj := recover(); obj != nil {
			buf := make([]byte, 4096)
			n := runtime.Stack(buf, false)
			fmt.Printf("Recovered: %v\n", obj)
			fmt.Printf("Stack trace:\n%s", buf[:n])
		}
	}()

	if c.OnReceiveMessage != nil {
		c.OnReceiveMessage(c, input)
	}
}

// LoadCharacterList 加载并获取角色列表
// 当客户端登录时有用
func LoadCharacterList(worldId int, accountId int64) []*Character {
	characterList := make([]*Character, 0)

	modelList, err := database.CharactersList(accountId, worldId)
	if err != nil {
		log.Printf("client.LoadCharacterList error. reason:%v", err)
		return nil
	}

	for i := 0; i < len(modelList); i++ {
		char := NewCharacterFromModel(modelList[i])
		err := char.LoadFromDB(false)
		if err != nil {
			log.Printf("client.LoadCharacterList char.LoadFromDB() error. reason: %v", err)
			return nil
		}
		characterList = append(characterList, char)
	}

	return characterList
}
