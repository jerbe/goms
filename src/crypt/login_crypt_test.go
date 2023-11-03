package crypt

import (
	"log"
	"testing"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/28 01:31
  @describe :
*/

func TestLoginCrypt(t *testing.T) {
	log.Println(LoginCrypt("admin"))
}
