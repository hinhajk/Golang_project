package _case

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var DB *gorm.DB
var dsn = "root:19681974ll@tcp(127.0.0.1)/mydb?charset=utf8&parseTime=True"

// 初始化
func init() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
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
	setPool(DB)
}

// 设置数据库连接池
func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
		return
	}
	sqlDB.SetMaxIdleConns(5)            // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //设置最长活跃时间
	sqlDB.SetMaxOpenConns(10)
}
