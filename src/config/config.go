package config

import (
	"flag"
	"log"
	"strings"

	"github.com/go-ini/ini"
	"github.com/jerbe/goms/utils"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/22 01:29
  @describe :
*/

var configFile = flag.String("config_file", "setting.ini", "配置文件")

func init() {
	log.Println("[config] 开始初始")
	utils.InitFuncQueue.Push(initConfig)
}

func initConfig() {
	log.Println("[config] 初始中...")
	file, err := ini.Load(*configFile)
	utils.PanicError(err)
	section := file.Section("")

	Default.DataSource.Driver = section.Key("datasource.driver").String()
	Default.DataSource.Url = section.Key("datasource.url").String()
	Default.DataSource.Username = section.Key("datasource.username").String()
	Default.DataSource.Password = section.Key("datasource.password").String()

	Default.Server.Name = section.Key("server.name").MustString("")
	Default.Server.Flag = section.Key("server.flag").MustInt(3)
	Default.Server.Limit.Online = section.Key("server.limit.online").MustInt(100)
	Default.Server.Limit.Characters = section.Key("server.limit.characters").MustInt(3)
	Default.Server.Address = section.Key("server.address").MustString("127.0.0.1")
	Default.Server.Register.Auto = section.Key("server.register.auto").MustBool(true)
	Default.Server.Rand.Drop = section.Key("server.rand.drop").MustBool(false)
	Default.Server.Login.Port = section.Key("server.login.port").MustInt(9595)
	Default.Server.Login.Admin = section.Key("server.login.admin").MustBool(false)
	Default.Server.Login.Message.Content = section.Key("server.login.message").String()
	Default.Server.Login.Message.Event = section.Key("server.login.message.event").String()
	Default.Server.Channel.Port = section.Key("server.channel.port").MustInt(7575)
	Default.Server.Channel.Count = section.Key("server.channel.count").MustInt(3)
	Default.Server.World.Rate.Exp = section.Key("server.world.rate.exp").MustInt(1)
	Default.Server.World.Rate.Gold = section.Key("server.world.rate.gold").MustInt(1)
	Default.Server.World.Rate.Drop.Normal = section.Key("server.world.rate.drop").MustInt(1)
	Default.Server.World.Rate.Drop.Boss = section.Key("server.world.rate.drop.boss").MustInt(1)
	Default.Server.World.Rate.Cash = section.Key("server.world.rate.cash").MustInt(1)
	Default.Server.World.Flags = section.Key("server.world.flags").MustInt()
	Default.Server.Mall.Port = section.Key("server.mall.port").MustInt(8600)
	Default.Server.Debug.Enabled = section.Key("server.debug.enabled").MustBool(false)
	Default.Server.Logger.Packet.Display = section.Key("server.logger.packet").MustBool(false)
	Default.Server.Logger.Packet.Debug = section.Key("server.logger.packet.debug").MustBool(false)
	Default.Server.Job.Adventurer = section.Key("server.job.adventurer").MustBool(true)
	Default.Server.Job.Knights = section.Key("server.job.knights").MustBool()
	Default.Server.Job.WarGod = section.Key("server.job.war-god").MustBool()
	Default.Server.Events = strings.Split(section.Key("server.events").String(), ",")
	Default.Server.Mall.Disabled = strings.Split(section.Key("server.mall.disabled").String(), ",")
	Default.Server.CashJy = strings.Split(section.Key("server.cashjy").String(), ",")
	Default.Server.Qysj = strings.Split(section.Key("server.gysj").String(), ",")
}

var Default = &config{}

type DataSourceConfig struct {
	Driver   string
	Url      string
	Username string
	Password string
}

type ServerConfig struct {
	// Name 服务器名称
	Name string

	// Flag 游戏标致
	Flag int

	// Limit 游戏限制
	Limit struct {
		// Online 在线人数
		Online int

		// Characters 角色数量
		Characters int
	}

	// Address 游戏服务IP地址
	Address string

	// Register 注册设置
	Register struct {
		// Auto 开启自动注册
		Auto bool
	}

	// Rand 随机
	Rand struct {
		// Drop 是否随机掉落物品
		Drop bool
	}

	// Login 登陆服务设置
	Login struct {
		// Port 服务端口
		Port int
		// Admin 是否仅允许管理员登录
		Admin bool

		// Message 上线后发给玩家的消息
		Message struct {
			// Content 消息内容
			Content string

			// Event 消息事件
			Event string
		}
	}

	// Channel 频道设置
	Channel struct {
		// Port 频道起始端口
		Port int

		// Count 频道数量，遍历生成以起始端口为基础的多个频道端口
		Count int
	}

	// World 世界设置
	World struct {
		// Rate 倍数
		Rate struct {
			// Exp 经验倍数
			Exp int

			// Gold 金币倍数
			Gold int

			// Drop 物品掉落倍数
			Drop struct {
				// Normal 普通物品掉落倍数
				Normal int

				// Boss 物品掉落倍数
				Boss int
			}

			// Cash 点卷掉落倍数
			Cash int
		}

		// Flags (WFlags) 世界标志,??我也不知道是什么
		Flags int
	}

	// Mall 商城服务配置
	Mall struct {
		// Port 服务端口
		Port int

		// Disabled  商城物品屏蔽ID
		Disabled []string
	}

	// Debug 调试配置
	Debug struct {
		// Enabled 是否启用调试
		Enabled bool
	}

	// Logger 日志配置
	Logger struct {
		// Packet 封包配置
		Packet struct {
			// Display 显示封包日志
			Display bool

			// Debug 显示调试封包日志
			Debug bool
		}
	}

	// Job 职业群配置
	Job struct {
		// Adventurer 冒险家职业群开关
		Adventurer bool

		// Knights 骑士职业群开关
		Knights bool

		// WarGod  战神职业群开关
		WarGod bool
	}

	// Events 启用的事件列表
	Events []string

	// CashJy 点券交易?
	CashJy []string

	// Qysj 未知?
	Qysj []string
}

// config 配置
type config struct {
	// DataSource 数据源配置
	DataSource DataSourceConfig

	// Server 服务配置
	Server ServerConfig
}
