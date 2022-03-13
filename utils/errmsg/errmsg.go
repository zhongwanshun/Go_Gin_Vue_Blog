package errmsg

//声明错误代码
const (
	SUCCESS = 200
	ERROR   = 500
	// code= 1000... 用户模块的错误

	ErrorUsernameUsed    = 1001
	ErrorPasswordWrong = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist   = 1004
	ErrorTokenRuntime      = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008
	// code= 2000... 微博模块的错误

	ErrorArtNotExist = 2001
)

//声明错误信息字典(int是错误代码,string就是错误信息)
var codeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "FAIL",
	ErrorUsernameUsed:   "用户名已存在!",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenExist:     "TOKEN不存在,请重新登陆",
	ErrorTokenRuntime:   "TOKEN已过期,请重新登陆",
	ErrorTokenWrong:     "TOKEN不正确,请重新登陆",
	ErrorTokenTypeWrong: "TOKEN格式错误,请重新登陆",
	ErrorUserNoRight:    "该用户无权限",

	ErrorArtNotExist: "微博不存在",
}

// GetErrMsg 输出错误信息的一个函数,直接将错误代码对应的错误信息返回
func GetErrMsg(code int) string {
	return codeMsg[code]
}
