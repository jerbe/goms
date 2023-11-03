package optional

import (
	"log"
	"testing"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/23 21:42
  @describe :
*/

func TestOptionalString(t *testing.T) {
	opt := Of("hello world")
	opt = Of(opt)
	log.Println(opt)
}

func TestMap(t *testing.T) {

}
