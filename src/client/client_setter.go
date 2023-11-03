package client

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 14:32
  @describe :
*/

// SetAccountID 设置账户名
func (c *Client) SetAccountID(id int64) {
	c.accountID = id
}

// SetAccountName 设置账户名
func (c *Client) SetAccountName(name string) {
	c.accountName = name
}

// SetMacAddress 设置网卡的mac地址
func (c *Client) SetMacAddress(mac string) {

}

// SetWorld 设置世界
func (c *Client) SetWorld(world int) {
	c.world = world
}

// SetChannel 设置频道
func (c *Client) SetChannel(channel int) {
	c.channel = channel
}
