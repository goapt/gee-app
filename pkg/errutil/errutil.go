package errutil

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"net"
)

// IsSystemError 判断错误是否为系统错误
// 系统错误的定义通常指的是网络通信类的错误
func IsSystemError(err error) bool {

	errs := make([]error, 0)
	// database
	errs = append(errs, sql.ErrConnDone)
	errs = append(errs, sql.ErrTxDone)
	errs = append(errs, driver.ErrBadConn)

	for _, v := range errs {
		if errors.Is(err, v) {
			return true
		}
	}

	// network
	if _, ok := err.(net.Error); ok {
		return true
	}

	return false
}
