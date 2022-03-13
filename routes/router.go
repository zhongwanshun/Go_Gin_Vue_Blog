package routes

import (
	v1 "ginweibo/api/v1"
	"ginweibo/middleware"
	"ginweibo/utils"

	"github.com/gin-gonic/gin"
)

// InitRouter 路由的入口文件
func InitRouter() {
	//设置App的mode
	gin.SetMode(utils.AppMode)
	//可以使用New也可以使用Default(默认增加了两个中间件,一个是路由文件)
	r := gin.New()
	r.Use(middleware.Log())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.GET("admin/users", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser) //传入用户的id
		auth.DELETE("user/:id", v1.DeleteUser)
		auth.PUT("admin/changepw/:id", v1.ChangeUserPassword) //修改密码

		// 微博模块的路由接口
		auth.GET("admin/article/info/:id", v1.GetArtInfo)
		auth.GET("admin/article", v1.GetArt)
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)

		// 评论模块
		auth.GET("comment/list", v1.GetCommentList)
		auth.DELETE("delcomment/:id", v1.DeleteComment)
	}

	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 微博模块
		router.GET("article", v1.GetArt)
		router.GET("article/info/:id", v1.GetArtInfo)

		// 登录控制模块
		router.POST("login", v1.Login)

		// 评论模块
		router.POST("addcomment", v1.AddComment)
		router.GET("comment/info/:id", v1.GetComment)
		router.GET("commentfront/:id", v1.GetCommentListFront)
		router.GET("commentcount/:id", v1.GetCommentCount)
	}

	_ = r.Run(utils.HttpPort)

}
