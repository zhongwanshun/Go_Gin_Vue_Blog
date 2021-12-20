package main

import (
	"ZWS_Go/model"
	"ZWS_Go/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
