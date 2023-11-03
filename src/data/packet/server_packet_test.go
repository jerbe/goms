package packet

import (
	"fmt"
	"testing"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 17:50
  @describe :
*/

func TestServerList(t *testing.T) {
	load := map[int]int{1: 110, 2: 120, 3: 130}
	packet := ServerList(0, 3, "你好世界", "", load)
	bytes := packet.Bytes()
	fmt.Println(bytes)
	fmt.Println(string(bytes))
}
