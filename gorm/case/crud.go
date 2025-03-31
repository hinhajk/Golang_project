package _case

import (
	"fmt"
	"time"
)

func Crud() {
	t := Teacher{
		Name:     "nick",
		Age:      48,
		Roles:    []string{"普通用户", "讲师"},
		Birthday: time.Now().Unix(),
		Salary:   1234.1234,
		Email:    "nick@gmail.com",
	}

	//添加
	res := DB.Create(&t)
	fmt.Print(res.RowsAffected, res.Error)

	//查找
	t1 := Teacher{}
	DB.First(&t1) //查找第一条记录
	fmt.Println(t1)

	// 更改
	t1.Name = "King"
	t1.Age = 31
	DB.Save(t1) //保存记录
	//DB.Where("name = ?", "King").First(&t1) //返回值也是一个链式结构，可以再次查询
	//DB.Where("name = ?", "King")

	// 使用关联语句就是使用表联查，会比较慢

	// 删除
	DB.Delete(&t1)
}
