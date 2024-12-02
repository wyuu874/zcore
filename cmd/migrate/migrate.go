package migrate

import (
	"fmt"
	"github.com/wyuu874/zcore/internal/goosex"
	"github.com/wyuu874/zcore/internal/viperx"
	"github.com/wyuu874/zcore/pkg/config"
	"log"
	"os"
)

// Migrate 迁移
func Migrate(args []string) {
	// 获取数据库配置
	viperx.InitConfig()
	dbConfig := config.Database{}
	config.GetConfig("database", &dbConfig)

	// 判断如果目录不存在则创建
	dir := dbConfig.MigrationsDir
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	// 连接数据库
	dbstring := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database, dbConfig.Charset)

	// 创建迁移器
	m, err1 := goosex.New(dbstring, dir)
	if err1 != nil {
		log.Fatal(err1)
	}
	defer m.Close()

	// 根据命令行参数执行相应操作
	if len(args) < 2 {
		log.Fatal("请指定操作: up, down, status, create")
	}

	var err error
	switch args[1] {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	case "status":
		err = m.Status()
	case "create":
		if len(args) < 3 {
			log.Fatal("请指定迁移文件名")
		}
		err = m.Create(args[2])
	default:
		log.Fatal("未知的命令")
	}

	if err != nil {
		log.Fatal(err)
	}
}
