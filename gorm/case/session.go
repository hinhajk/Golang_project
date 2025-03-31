package _case

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func Session() {
	tx := DB.Session(&gorm.Session{
		PrepareStmt:              true, //是否预编译
		SkipHooks:                true,
		DisableNestedTransaction: true,
		Logger:                   DB.Logger.LogMode(logger.Error), //设置是否打印日志
	})

	t := Teacher{
		Name:     "nick",
		Age:      48,
		Roles:    []string{"普通用户", "讲师"},
		Birthday: time.Now().Unix(),
		Salary:   1234.1234,
		Email:    "nick@gmail.com",
	}
	tx.Create(&t)

}
