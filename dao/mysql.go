package dao

import (
	"fmt"
	"ginweb/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

// 连接数据库
func InitMySQL(cfg *setting.MySQLConfig) (err error) {
	//dsn := "root:root@tcp(127.0.0.1:3306)/app?charset=utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return DB.Error
	}
	return
}

// 关闭数据库 V2 不用关闭
//func Close() {
//	DB.Close()
//}
