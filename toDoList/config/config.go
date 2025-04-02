package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
	"toDoList/models"
)

// 将config.ini中的配置文件读取出来
var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

// 初始化函数
// 利用init依赖读取ini文件内容传给file变量
func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径")
	}
	LoadServer(file)
	LoadMysql(file)

	// MySQL数据库配置的读写
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ")/", DbName, "?charset=utf8&parseTime=True"}, "")
	models.DataBase(path)
}

// 将服务器配置信息给加载出来
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
	// 将file 文件中server模块中键值为"AppMode"或者其他的部分读取出来兵器转换为对应的string类型
}

// 将MySQL配置信息加载出来
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	//DbPort = file.Section("mysql").Key("DbPOrt").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
