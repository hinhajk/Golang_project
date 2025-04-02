package models

// 该文件中进行数据库连接

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

// 这个dsn与config.go中的path本质是一致的，连接mysql的路径
//var dsn = "root:19681974ll@tcp(127.0.0.1)/mydb?charset=utf8&parseTime=True"

func DataBase(connstring string) {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               connstring,
		DefaultStringSize: 256,
	}), &gorm.Config{
		//gorm配置
		Logger:      logger.Default.LogMode(logger.Info), //每次操作数据库打印日志
		PrepareStmt: true,                                //全局预编译
		//每次编译的时候都会将SQL语句进行预编译缓存
		//预编译不支持嵌套事务
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("数据库连接成功")
	setPool(DB)
	Migration()
}

// 设置数据库连接池
func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	sqlDB.SetMaxIdleConns(20)           // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //设置最长活跃时间
	sqlDB.SetMaxOpenConns(100)
}
