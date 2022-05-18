package main

import (
	"fmt"
	"ginweb/dao"
	"ginweb/models"
	"ginweb/routers"
	"ginweb/setting"
	"os"
)

func main() {
	// 加载配置文件
	//conf := "config.ini"
	conf := "debug.ini"
	dir, _ := os.Getwd()
	if err := setting.Init(dir + "/conf/" + conf); err != nil {
		fmt.Printf("load config from file failed, err:%v\n", err)
		return
	}

	// 创建数据库 连接数据库
	err := dao.InitMySQL(setting.Conf.MySQLConfig)
	if err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	// 程序退出关闭数据库连接
	//defer dao.Close()
	// 模型绑定
	DBMigrate()
	// 注册路由
	engine := routers.SetupRouter()

	if err := engine.Run(fmt.Sprintf(":%d", setting.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}

// 模型绑定
func DBMigrate() {
	dao.DB.AutoMigrate(&models.UserInfo{})
	dao.DB.AutoMigrate(&models.FileInfo{})
	dao.DB.AutoMigrate(&models.MarkInfo{})
}