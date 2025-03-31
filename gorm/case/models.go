package _case

import (
	"gorm.io/gorm"
)

//建造模型
//gorm提供的tag
//float精度问题：
//check约束：
//序列化：
//结构内嵌：

// 建表
func init() {
	err := DB.Migrator().AutoMigrate(Teacher{})
	//建立一张teacher表
	if err != nil {
		return
	}
	//建造多个表时在后面加即可
	//例如：DB.Migrator().AutoMigrate(Teacher{}, Course{})
}

type Roles []string

type Teacher struct {
	gorm.Model         //内嵌gorm.Model结构体， 保活ID、Created At、Updated At、Deleted At
	Name       string  `gorm:"size:255"` //使用的时候去gorm官网上查看文档即可
	Email      string  `gorm:"size:255"`
	Salary     float64 `gorm:"scale:2;precision:6"`
	Age        uint8   `gorm:"check:age > 30"`
	Roles      Roles   `gorm:"serializer:json"`
	Birthday   int64   `gorm:"serializer:unixtime;type:time"`
	//创建时间
	//更新时间
	//删除时间
}
