package client

import (
	"database/sql"
	"errors"
	"github.com/jerbe/goms/config"
	"github.com/jerbe/goms/crypt"
	"github.com/jerbe/goms/database"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 14:33
  @describe :
*/

// GetWorld 获取当前用户所选择的世界
func (c *Client) GetWorld() int {
	return c.world
}

// GetChannel 获取当前用户所选择的频道
func (c *Client) GetChannel() int {
	return c.channel
}

// GetAccountID 获取账户ID
func (c *Client) GetAccountID() int64 {
	return c.accountID
}

// GetAccountName 获取账户名
func (c *Client) GetAccountName() string {
	return c.accountName
}

// GetMacAddress 获取网卡的mac地址
func (c *Client) GetMacAddress() string {
	return ""
}

// GetLocalAddress 获取本地连接地址
func (c *Client) GetLocalAddress() string {
	return c.conn.LocalAddr().String()
}

// GetRemoteAddress 获取远程地址
func (c *Client) GetRemoteAddress() string {
	return c.conn.RemoteAddr().String()
}

// GetCharacterSlots 获取允许的角色卡槽数量
func (c *Client) GetCharacterSlots() int {
	if c.isGM {
		return 15
	}

	if c.characterSlots != 0 {
		return c.characterSlots
	}
	defaultCharSlots := config.Default.Server.Limit.Characters

	slots, err := database.CharactersSlots(c.accountID, c.world)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0
	}

	if errors.Is(err, sql.ErrNoRows) {
		_, err := database.AddCharactersSlots(c.accountID, c.world, defaultCharSlots)
		if err != nil {
			return 0
		}
		slots = defaultCharSlots
	}
	c.characterSlots = slots
	return slots

}

// GetEncodeCypher 返回编码加密器
func (c *Client) GetEncodeCypher() *crypt.MapleCypher {
	return c.encodeCypher
}

// GetDecodeCypher 返回解码加密器
func (c *Client) GetDecodeCypher() *crypt.MapleCypher {
	return c.decodeCypher
}
