package client

import (
	"sync"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/11/2 09:52
  @describe :
*/

type containerElem struct {
	cli  *Client
	char *Character
}

// CharacterContainer 角色容器
type CharacterContainer struct {
	rwMutex sync.RWMutex

	elems map[int64]*containerElem
}

// Register 注册角色
func (c *CharacterContainer) Register(cli *Client, char *Character) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	c.elems[char.GetID()] = &containerElem{cli: cli, char: char}
}

// Unregister 注销角色
func (c *CharacterContainer) Unregister(char *Character) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()
	delete(c.elems, char.GetID())
}
