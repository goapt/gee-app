package response

import (
	"net/http"

	"app/api/constant"
	"app/pkg/errutil"
	"github.com/goapt/gee"
	"github.com/goapt/golib/debug"
	"github.com/goapt/logger"
)

type Response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// ParamError 参数错误
func ParamError(c *gee.Context, msg string) gee.Response {
	return Error(c, constant.ErrInvalidParameter, msg)
}

// SystemError 500系统错误
func SystemError(c *gee.Context, msg string) gee.Response {
	logger.Data(map[string]interface{}{"stack": string(debug.Stack(20))}).Errorf("[SystemError]%s", msg)
	return WithStatusError(c, http.StatusInternalServerError, constant.ErrSystemError, msg)
}

// WithSystemError 当被调用方返回的错误可能是系统错误，也可能是业务错误，如果不是系统级别的错误，则返回指定的Code
func WithSystemError(c *gee.Context, code string, err error) gee.Response {
	if errutil.IsSystemError(err) {
		return SystemError(c, err.Error())
	}

	return Error(c, code, err.Error())
}

// Error 业务默认都是400
func Error(c *gee.Context, code, msg string) gee.Response {
	return WithStatusError(c, http.StatusBadRequest, code, msg)
}

// WithStatusError 设置错误的http status
func WithStatusError(c *gee.Context, status int, code, msg string) gee.Response {
	c.Status(status)
	return c.JSON(&Response{
		Code: code,
		Msg:  msg,
	})
}

// Success 成功返回
func Success(c *gee.Context, data interface{}) gee.Response {
	// 如果是对外接口，则直接返回结果
	return c.JSON(data)
	// 如果是前端接口则使用Response返回
	// return c.JSON(&Response{
	// 	Code: "Success",
	// 	Msg:  "ok",
	// 	Data: data,
	// })
}
