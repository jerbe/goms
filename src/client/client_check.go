package client

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 14:45
  @describe :
*/

// IsClose 判断是否已经关闭
func (c *Client) IsClose() (bool, error) {
	select {
	case <-c.ctx.Done():
		return true, c.ctx.Err()
	default:
		return false, nil
	}
}

// IsBanedIP 是否是被禁用IP
func (c *Client) IsBanedIP() bool {
	return false
}

// IsBanedMAC 是否是被禁用mac
func (c *Client) IsBanedMAC() bool {
	return false
}

// IsGM 判断该账户是否是GM账户
func (c *Client) IsGM() bool {
	return c.isGM
}
