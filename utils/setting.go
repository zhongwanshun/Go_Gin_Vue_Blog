package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

//变量引入
var (
	//服务器相关变量
	AppMode  string
	HttpPort string
	JwtKey   string
	//数据库相关变量
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

//作为一个函数初始化的一个接口
func init() {
	//引入文件，路径为绝对路径
	file, err := ini.Load("config/config.ini") //返回的是一个file指针,是一个结构体
	//处理错误数据
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	//实现file指针里面的结构体
	LoadServer(file)
	LoadData(file)
}

//实现server
func LoadServer(file *ini.File) {
	//注释:当在server里面无法访问到Appmode的时候,就是用默认值
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

//实现数据库操作
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("1234")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
