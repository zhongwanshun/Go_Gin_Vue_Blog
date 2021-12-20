package errmsg

const (
	//通过状态码反馈相关信息
	SUCCSE = 200
	ERROR  = 500
	//分配状态码
	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED  = 1001
	ERROR_PASSWORD_WRONG = 1002
	ERROR_USER_NOT_EXIST = 1003
	ERROR_TOKEN_EXIST    = 1004
	ERROR_TOKEN_RUNTIME  = 1005
	ERROR_TOKEN_WRONG    = 1006
	//TOKEN格式不正确
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code= 2000... 文章模块的错误

	ERROR_ART_NOT_EXIST = 2001
	// code= 3000... 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

//字典
var codeMsg = map[int]string{
	SUCCSE: "OK",
	ERROR:  "FAIL",
	//用户模块
	ERROR_USERNAME_USED:    "用户名已存在!",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	//文章模块
	ERROR_ART_NOT_EXIST: "文章不存在",
	//分类模块
	ERROR_CATENAME_USED:  "该分类已存在",
	ERROR_CATE_NOT_EXIST: "该分类不存在",
}

//大写其他才可以引用,返回一个字符串信息
func GetErrMsg(code int) string {
	return codeMsg[code]
}
