package crypt

import "github.com/jerbe/goms/utils/bit"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/25 10:16
  @describe :
*/

// MapleEncrypt 枫叶传说专用加密方法
func MapleEncrypt(data []byte) {
	for j := 0; j < 6; j++ {
		remember := byte(0)
		dataLength := (byte)(len(data) & 0xFF)
		if j%2 == 0 {
			for i := 0; i < len(data); i++ {
				cur := data[i]
				cur = bit.RollLeft(cur, 3)
				cur += dataLength
				cur ^= remember
				remember = cur
				cur = bit.RollRight(cur, int(dataLength&0xFF))
				cur = (^cur) & 0xFF
				cur += 0x48
				dataLength--
				data[i] = cur
			}
		} else {
			for i := len(data) - 1; i >= 0; i-- {
				cur := data[i]
				cur = bit.RollLeft(cur, 4)
				cur += dataLength
				cur ^= remember
				remember = cur
				cur ^= 0x13
				cur = bit.RollRight(cur, 3)
				dataLength--
				data[i] = cur
			}
		}
	}
}

// MapleDecrypt 枫叶传说专用解密方法
func MapleDecrypt(data []byte) {
	for j := 1; j <= 6; j++ {
		remember := byte(0)
		dataLength := byte(len(data) & 0xFF)
		nextRemember := byte(0)

		if j%2 == 0 {
			for i := 0; i < len(data); i++ {
				cur := data[i]
				cur = cur - 72
				cur = (^cur) & 0xFF
				cur = bit.RollLeft(cur, int(dataLength&0xFF))
				nextRemember = cur
				cur = cur ^ remember
				remember = nextRemember
				cur = cur - dataLength
				cur = bit.RollRight(cur, 3)
				data[i] = cur
				dataLength = dataLength - 1
			}
		} else {
			for i := len(data) - 1; i >= 0; i-- {
				cur := data[i]
				cur = bit.RollLeft(cur, 3)
				cur = cur ^ 0x13
				nextRemember = cur
				cur = cur ^ remember
				remember = nextRemember
				cur = cur - dataLength
				cur = bit.RollRight(cur, 4)
				data[i] = cur
				dataLength = dataLength - 1
			}
		}
	}
}
