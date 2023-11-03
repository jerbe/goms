package wz

import (
	"os"

	"github.com/jucardi/go-streams/streams"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/23 10:26
  @describe :
*/

func NewWzDirectory(dirPath *os.File) *WzDirectory {
	return &WzDirectory{
		dirs:  make(map[string]*WzDirectory),
		files: make(map[string]*WzFile),
		path:  dirPath,
	}
}

type WzDirectory struct {
	dirs map[string]*WzDirectory

	files map[string]*WzFile

	path *os.File
}

// Parse 解析当前目录文件下的所有内容，包括子目录和子文件。
func (d *WzDirectory) Parse() {

}

func (d *WzDirectory) FindDir(name string) {

}

func (d *WzDirectory) FindFile(name string) {

}

func (d *WzDirectory) DirStream() *streams.IStream {
	return nil
}

func (d *WzDirectory) FileStream() {

}
