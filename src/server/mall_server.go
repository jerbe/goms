package server

import (
	"context"
	"github.com/jerbe/goms/client"
	"github.com/jerbe/goms/config"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/22 21:02
  @describe :
*/

// ICashShopServer 现金商城服务器接口
type ICashShopServer interface {
	IServer

	// EnterCashShop 进入的现金商城
	EnterCashShop(cli *client.Client, characterId int)
}

type IMallServer interface {
	IServer
}

//var _ base.IMallServer = &MallServer{}

// NewMallServer 返回商城服务器
func NewMallServer(ctx context.Context, config config.ServerConfig) *MallServer {
	svr := &MallServer{
		BaseServer: *NewBaseServer(ctx, config.Channel.Port, config.Name),
	}
	svr.SetFlag(config.Flag)

	return svr
}

type MallServer struct {
	BaseServer
}
