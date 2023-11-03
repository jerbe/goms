package wz

import "os"

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/23 10:26
  @describe :
*/

type WzData struct {
	directory *WzDirectory
}

func NewWzData(file *os.File) *WzData {
	return nil
}

func (d *WzData) Directory() *WzDirectory {
	return d.directory
}
