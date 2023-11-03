package utils

import (
	"fmt"
	"os"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/23 11:38
  @describe :
*/

// ChildFile 获取子文件
func ChildFile(parent *os.File, name string) (*os.File, error) {
	return os.Open(fmt.Sprintf("%s%s%s", parent.Name(), string(os.PathSeparator), name))
}
