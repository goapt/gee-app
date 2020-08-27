package response

import (
	"net/http"

	"app/errorx"
)

var (
	// 全局错误
	ErrInvalidSignature = newStatusError(http.StatusUnauthorized, errorx.New("InvalidSignature", "签名错误"))
	ErrRateLimited      = newStatusError(http.StatusTooManyRequests, errorx.New("RateLimited", "接口访问太频繁"))
	ErrAccessForbidden  = newStatusError(http.StatusForbidden, errorx.New("AccessForbidden", "无权访问该接口"))
	// 系统错误
	ErrSystemError = newStatusError(http.StatusInternalServerError, errorx.New("SystemError", "系统错误"))

	// 参数错误
	ErrInvalidParameter = errorx.New("InvalidParameter", "参数错误")
	ErrAuthFailure      = errorx.New("AuthFailure", "非法访问")
)
