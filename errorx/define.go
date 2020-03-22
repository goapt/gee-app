package errorx

// 10000 成功
// 第一类，服务器错误：
// 10001 服务器异常
//
// 第二类 业务错误
// 20001 参数错误
//
// 第三类：权限类控制等
// 40000 默认错误
// 40003 登录权限错误
var (
	ErrDatabase     = New(10001, "服务器异常")
	ErrInvalidParam = New(20001, "参数错误，请检查参数是否缺失或类型是否匹配")
	ErrBusiness     = New(40000, "业务错误")
	ErrSession      = New(40003, "登录异常")
	ErrRouterFound  = New(40004, "router not found")
	ErrHTTP         = New(40005, "HTTP Error")
)
