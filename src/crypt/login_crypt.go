package crypt

import (
	"crypto/sha1"
	"encoding/hex"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/28 01:28
  @describe :
*/

// LoginCrypt 登录用的加密
func LoginCrypt(data string) string {
	s := sha1.New()
	s.Write([]byte(data))
	return hex.EncodeToString(s.Sum(nil))
}
