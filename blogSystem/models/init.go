package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB

func DataBase(path string) {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               path,
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
