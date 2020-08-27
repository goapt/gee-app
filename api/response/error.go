package response

import (
	"net/http"

	"github.com/goapt/gee"
	"github.com/goapt/golib/debug"
	"github.com/goapt/logger"

	"app/errorx"
)

type ErrorResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func newErrorResponseWithMsg(code, msg, changeMsg string) *ErrorResponse {
	if changeMsg != "" {
		msg = changeMsg
	}

	return &ErrorResponse{
		Code: code,
		Msg:  msg,
	}
}

type statusError struct {
	Status int
	*errorx.Errorx
}

func newStatusError(httpStatus int, err *errorx.Errorx) *statusError {
	return &statusError{
		httpStatus,
		err,
	}
}

// 参数错误400
func ParamError(c *gee.Context, msg ...string) gee.Response {
	return Error(c, ErrInvalidParameter, msg...)
}

func Error(c *gee.Context, err error, msg ...string) gee.Response {
	m := ""
	if len(msg) > 0 {
		m = msg[0]
	}
	// 默认都是400
	c.HttpStatus = http.StatusBadRequest

	// 1、定义了http status的Errorx
	if r, ok := err.(*statusError); ok {
		c.HttpStatus = r.Status
		return c.JSON(newErrorResponseWithMsg(r.Code, r.Message, m))
	}

	// 2、errorx.Errorx错误
	if er, ok := err.(*errorx.Errorx); ok {
		return c.JSON(newErrorResponseWithMsg(er.Code, er.Error(), m))
	}

	// 3、500系统错误
	if errorx.IsSystemError(err) {
		c.HttpStatus = http.StatusInternalServerError
		logger.Data(map[string]interface{}{"stack": string(debug.Stack(20))}).Errorf("[SystemError]%s", err.Error())
		return c.JSON(newErrorResponseWithMsg(ErrSystemError.Code, err.Error(), m))
	}

	// 4、非errorx.Errorx的普通错误，理论不应该直接将这样的错误抛出，需要使用errorx定义一个code
	logger.Data(map[string]interface{}{"stack": string(debug.Stack(20))}).Errorf("[UnkonwError]%s", err.Error())
	return c.JSON(newErrorResponseWithMsg("UnkonwError", err.Error(), m))
}
