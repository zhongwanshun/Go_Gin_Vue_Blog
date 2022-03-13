package v1

//用户
import (
	"ginweibo/model"
	"ginweibo/utils/errmsg"
	"ginweibo/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	//引入结构体
	var data model.ZwsUser
	var msg string
	var validCode int
	//利用绑定模型的验证进行赋值，返回的是一个err
	_ = c.ShouldBindJSON(&data)

	msg, validCode = validator.Validate(&data)
	if validCode != errmsg.SUCCESS {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  validCode,
				"message": msg,
				"data":    data,
			},
		)
		c.Abort()
		return
	}
	//进行赋值,直接调用mode里面的ZwsCheckUser的方法就行
	code := model.CheckUser(c, data.Username)
	if code == errmsg.SUCCESS {
		//写入数据库
		model.CreateUser(c, &data)
	}
	//返回
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,                   //状态码(跟我们设置的不一样，是网络状态)
			"message": errmsg.GetErrMsg(code), //调用函数
		},
	)
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data, code := model.GetUser(c, id)
	maps["username"] = data.Username
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    maps,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	//使用query接收,query接收的是一个字符串，使用内置方法转换为Int类型
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username") //用户名不用转换,就是字符串

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		//如果页数等于=0的话，就只显示一页
		pageNum = 1
	}

	data, total := model.GetUsers(c, username, pageSize, pageNum)

	code := errmsg.SUCCESS
	//结果返回
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.ZwsUser
	//拿到id
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	//查询用户是否存在

	//执行修改
	code := model.EditUser(c, id, &data)
	//返回结果，只需要返回状态码
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.ZwsUser
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.ChangePassword(c, id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DeleteUser 删除用户(软删除)
func DeleteUser(c *gin.Context) {
	//类型转换
	id, _ := strconv.Atoi(c.Param("id"))
	//
	code := model.DeleteUser(c, id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
