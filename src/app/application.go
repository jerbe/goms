package app

import (
	"context"
	"github.com/jerbe/goms/handler"
	"github.com/jerbe/goms/server"
	"log"
	"time"

	"github.com/jerbe/goms/config"
	"github.com/jerbe/goms/database"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/22 20:50
  @describe :
*/

func NewApplication(ctx context.Context) *Application {
	return &Application{
		maxOnline:    0,
		serverConfig: config.Default.Server,
		ctx:          ctx,
	}
}

// Application 应用
type Application struct {
	// maxOnline 最大在线数量
	maxOnline int

	// serverConfig 服务器配置
	serverConfig config.ServerConfig

	// cashShopServer 现金商城服务器
	cashShopServer interface{}

	// loginServer 登陆服务器
	loginServer *server.ILoginServer

	// ctx 上下文数据
	ctx context.Context

	// cancel 取消方法
	cancelFunc context.CancelFunc
}

// Run 启动服务
func (app *Application) Run() {
	if app.serverConfig.Login.Admin {
		log.Println("[登陆模式] 只允许管理员登录")
	} else {
		log.Println("[登陆模式] 允许所有人登录")
	}

	if app.serverConfig.Register.Auto {
		log.Println("[注册模式] 自动生成账号")
	} else {
		log.Println("[注册模式] 手动生成账户,包括[网页注册/GM工具注册]")
	}

	ctx, cancel := context.WithCancel(app.ctx)
	app.cancelFunc = cancel

	log.Println("[账户信息] 重置中...")
	start := time.Now()
	affectLSRows, err := database.Default.ResetAllLoginStatus()
	if err != nil {
		log.Fatal(err)
	}
	affectLGHMRows, err := database.Default.ResetAllLastGainHiredMerchant()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[账户信息] 重置完毕，耗时：%v，影响行数：登录状态=%d 最后获得雇佣商人=%d", time.Now().Sub(start), affectLSRows, affectLGHMRows)

	start = time.Now()
	log.Println("[NPC数据] 初始化...")
	// 加载 WZdata
	log.Printf("[NPC数据] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[世界服务器] 初始化...")
	// 启动世界服务器
	log.Printf("[世界服务器] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[定时器] 初始化...")
	// WorldTimer
	//EtcTimer
	//MapTimer
	//MobTimer
	//CloneTimer
	//EventTimer // 事件
	//BuffTimer	// 附加效果
	//TimerManager // 点卷赌博?
	log.Printf("[定时器] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[任务数据] 初始化...")
	// MapleQuest
	log.Printf("[任务数据] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[道具数据] 初始化...")
	//
	log.Printf("[道具数据] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[脏话检测] 初始化...")
	//
	log.Printf("[脏话检测] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[随机奖励] 初始化...")
	//
	log.Printf("[随机奖励] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[OX题目] 初始化...")
	//
	log.Printf("[OX题目] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[技能数据] 初始化...")
	//
	log.Printf("[技能数据] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[排名系统] 初始化...")
	//
	log.Printf("[排名系统] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[商城道具] 初始化...")
	//
	log.Printf("[商城道具] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[登录服务器] 初始化...")
	loginServer := server.NewLoginServer(ctx, config.Default.Server)
	loginServerHandler := handler.NewHandler(loginServer)
	loginServer.OnClientConnect = loginServerHandler.OnClientConnect
	loginServer.OnClientReceiveMessage = loginServerHandler.OnClientReceiveMessage
	loginServer.Startup()

	log.Printf("[登录服务器] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[频道服务器] 初始化...")
	startupChannelServers(ctx)
	log.Printf("[频道服务器] 初始化完毕，耗时：%v", time.Now().Sub(start))

	start = time.Now()
	log.Println("[商城服务器] 初始化...")
	mallServer := server.NewMallServer(ctx, config.Default.Server)
	mallServer.Startup()
	log.Printf("[商城服务器] 初始化完毕，耗时：%v", time.Now().Sub(start))

	worldRate := config.Default.Server.World.Rate
	log.Printf("[经验倍率]:%d [物品倍率]:正常:%d,Boss:%d，[金币倍率]%d", worldRate.Exp, worldRate.Drop.Normal, worldRate.Drop.Boss, worldRate.Cash)

	job := config.Default.Server.Job
	log.Printf("[开放职业] 冒险家=>%v, 骑士团=>%v, 战神=>%v", job.Adventurer, job.Knights, job.WarGod)

	loginServer.SetOn()
	setonChannelServers()
	mallServer.SetOn()
}

func (app *Application) Shutdown() {
	app.cancelFunc()
}

// startupChannelServers 启动频道服务器
func startupChannelServers(ctx context.Context) {
	if len(handler.ChannelServerPool) > 0 {
		return
	}

	handler.ChannelServerPool = make(map[int]*server.ChannelServer)
	cfg := config.Default.Server.Channel

	for i := 1; i <= cfg.Count; i++ {
		svr := server.NewChannelServer(ctx, i)
		handler.ChannelServerPool[i] = svr
		h := handler.NewHandler(svr)
		svr.OnClientConnect = h.OnClientConnect
		svr.OnClientReceiveMessage = h.OnClientReceiveMessage
		err := svr.Startup()
		if err != nil {
			log.Panicln(err)
		}
	}
}

// setonChannelServers 将服务设置成开启
func setonChannelServers() {
	for _, svr := range handler.ChannelServerPool {
		err := svr.SetOn()
		if err != nil {
			log.Panicln(err)
		}
	}
}
