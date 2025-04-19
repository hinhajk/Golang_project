package config

//该文件作用在于读取配置文件信息，并且构造出要连接的数据库路径
import (
	"blogSystem/models"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

// 用于接收配置文件中的值
var (
	Db         string
	DbName     string
	DbUser     string
	DbHost     string
	DbPort     string
	DbPassWord string
	AppMode    string
	HttpPort   string
)

// 初始化函数
func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查配置文件路径")
	} //读取配置文件，用file变量接收
	LoadServer(file) //读取服务器信息
	LoadMysql(file)  //读取mysql数据库配置信息
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ")/", DbName, "?charset=utf8&parseTime=True"}, "")
	models.DataBase(path)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").String()
	HttpPort = file.Section("server").Key("HttpPort").String()
}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
}
