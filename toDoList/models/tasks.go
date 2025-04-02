package models

import "gorm.io/gorm"

// 建立任务表格
type Tasks struct {
	gorm.Model
	TaskName  string `gorm:"unique"`
	User      User   `gorm:"ForeignKey:Uid;references:id"` // 其Uid为外键，参考User表中的id键
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"` // 不能为空且为索引，加快查询
	Status    int    `gorm:"default:0"`      //任务状态：默认值为0代表未完成，1代表已完成
	Content   string `gorm:"type:longtext"`  //任务内容，longtext代表内容字段较长
	StartTime int64  // 备忘录开始时间
	EndTime   int64  //备忘录完成时间
}
