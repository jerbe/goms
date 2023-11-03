package database

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/22 21:26
  @describe :
*/
import (
	"context"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jerbe/goms/config"
	"github.com/jerbe/goms/utils"
	"github.com/jmoiron/sqlx"
)

// init 默认启动方法,只要引用该包就会执行
func init() {
	log.Println("[database] 开始初始")
	utils.InitFuncQueue.Push(initDatabase)
}

func initDatabase() {
	log.Println("[database] 初始中...")
	Default = &database{}

	// 从config那边拿来用
	mysqlCfg := config.Default.DataSource

	timoutCtx, _ := context.WithTimeout(context.Background(), time.Second*10)

	mysqlDB, err := sqlx.ConnectContext(timoutCtx, "mysql", mysqlCfg.Url)
	if err != nil {
		log.Fatalln(err)
	}

	mysqlDB.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	mysqlDB.SetMaxIdleConns(5)

	Default.sqlDB = mysqlDB
}

// Default 默认数据库管理器
var Default *database

// database 数据库管理器
type database struct {
	sqlDB *sqlx.DB
}

func (db *database) SqlDB() *sqlx.DB {
	return db.sqlDB
}
