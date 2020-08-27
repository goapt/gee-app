package errorx

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"net"
)

type Errorx struct {
	err     error
	Code    string
	Message string
}

func New(code string, msg interface{}) *Errorx {
	ex := &Errorx{
		Code: code,
	}

	var m string
	switch e := msg.(type) {
	case string:
		m = e
	case error:
		m = e.Error()
		ex.err = e
	default:
		m = fmt.Sprintf("[%s]%v", code, e)
	}

	if m == "" {
		m = fmt.Sprintf("[%s]unknow error", code)
	}

	ex.Message = m

	return ex
}

// Error 这里必须是指针，否则两个相同的错误信息的不同变量是相等的，这不是我们希望的结果，因此需要和官方保持一致
// err1 := New("EOF","EOF")
// err2 := New("EOF","EOF")
// err1 != err2
func (e *Errorx) Error() string {
	return e.Message
}

func (e *Errorx) Unwrap() error {
	return e.err
}

func (e *Errorx) Is(err error) bool {
	if er, ok := err.(*Errorx); ok {
		return e.Code == er.Code
	}
	return false
}

// PrettyNoRows 优化查无记录的错误 ，调用方无需额外判断是否为空记录，只需要指定遇到空记录时，替换的错误提示
func PrettyNoRows(err error, msg interface{}) error {
	if err == nil {
		return nil
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	var newerr error
	switch v := msg.(type) {
	case error:
		newerr = v
	case string:
		newerr = errors.New(v)
	default:
		newerr = errors.New(fmt.Sprint(v))
	}

	return newerr
}

// FilterNoRows 过滤数据不存在的错误
func FilterNoRows(err error) error {
	if err == nil || errors.Is(err, sql.ErrNoRows) {
		return nil
	}

	return err
}

// 判断是否是系统错误
func IsSystemError(err error) bool {
	// 数据库错误
	if errors.Is(err, sql.ErrConnDone) || errors.Is(err, sql.ErrTxDone) || errors.Is(err, driver.ErrBadConn) {
		return true
	}

	// 网络错误
	if _, ok := err.(net.Error); ok {
		return true
	}

	return false
}
