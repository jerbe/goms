package bit

import "math"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/25 10:19
  @describe :
*/

func GetShort(array []byte, index int) int {
	ret := int(array[index])
	ret &= 0xFF
	ret |= (int(array[index+1]) << 8) & 0xFF00
	return ret
}

func GetString(array []byte, index int, length int) string {
	cret := make([]byte, length)
	for x := 0; x < length; x++ {
		cret[x] = array[x+index]
	}
	return string(cret)
}

func GetMapleString(array []byte, index int) string {
	length := (int(array[index]) & 0xFF) | (int(array[index+1]) << 8 & 0xFF00)
	return GetString(array, index+2, length)
}

func RollLeft(in byte, count int) byte {
	tmp := int(in) & 0xFF
	tmp = tmp << (count % 8)
	return byte((tmp & 0xFF) | (tmp >> 8))
}

func RollRight(in byte, count int) byte {
	tmp := int(in) & 0xFF
	tmp = (tmp << 8) >> (count % 8)
	return byte((tmp & 0xFF) | (tmp >> 8))
}

func MultiplyBytes(in []byte, count, mul int) []byte {
	ret := make([]byte, count*mul)
	for x := 0; x < count*mul; x++ {
		ret[x] = in[x%count]
	}
	return ret
}

func DoubleToShortBits(d float64) int {
	l := int64(math.Float64bits(d))
	return int(l >> 48)
}
