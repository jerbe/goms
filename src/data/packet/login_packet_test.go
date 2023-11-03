package packet

import (
	"log"
	"testing"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/27 13:45
  @describe :
*/

func TestSayHello(t *testing.T) {
	s := []byte{82, 48, 120, 7}
	r := []byte{70, 114, 122, 241}
	packet := SayHello(79, s, r)
	log.Println(packet.data)
}

func TestLoginSuccess(t *testing.T) {
	packet := LoginSuccess(1, 1, true, "admin")
	log.Println(packet.Bytes())
}
