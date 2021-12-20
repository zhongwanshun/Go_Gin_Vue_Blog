package routes

import (
	"ZWS_Go/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

//首字母小写的话就是一个私有的方法
func InitRouter() {
	gin.SetMode(utils.AppMode)
	//r:=gin.New(),Default会默认添加两个中间件
	r := gin.Default()
	//初始化路由
	router := r.Group("api/v1")
	{
		router.GET("Hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	r.Run(utils.HttpPort)
}
