package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

//引入变量
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

//一个包初始化的一个函数接口
func init() {
	//file获得的是一个结构体,Load返回的是一个file的指针
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	//实现结构体
	LoadServer(file)
	LoadData(file)

}
func LoadServer(file *ini.File) {
	//读取相关变量(去server中取AppMode,取不到就用默认debug)
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")

}
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("1234")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
