package main

import (
	"context"
	"flag"
	"github.com/jerbe/goms/app"
	"github.com/jerbe/goms/utils"
	"os"
	"os/signal"
	"syscall"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/21 17:35
  @describe :
*/

/*
flags :
	-config_file 配置文件路径
	-wz.path wz根目录
*/

func init() {
	flag.Parse()

}

func main() {
	utils.InitFuncQueue.Run()

	ctx, cancel := context.WithCancel(context.Background())
	app := app.NewApplication(ctx)
	app.Run()

	// shutdownSig 关闭信号
	shutdownSig := make(chan os.Signal)
	signal.Notify(shutdownSig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	<-shutdownSig
	cancel()
}
