package errorx

import (
	"fmt"
)

type Errorx struct {
	err     error
	Code    int
	Message string
}

func New(code int, msg interface{}) Errorx {
	ex := Errorx{
		Code: code,
	}

	var m string
	switch e := msg.(type) {
	case Errorx:
		m = e.Error()
		ex.err = e
	case string:
		m = e
	case error:
		m = e.Error()
		ex.err = e
	default:
		m = fmt.Sprint(e)
	}

	if m == "" {
		m = "Undefind Error"
	}

	ex.Message = m

	return ex
}

func (e Errorx) Error() string {
	return fmt.Sprintf("%s(code=%d)", e.Message, e.Code)
}

func (e Errorx) Unwrap() error {
	return e.err
}

func (e Errorx) Is(err error) bool {
	if er, ok := err.(Errorx); ok {
		return e.Code == er.Code
	}
	return false
}

func Database(err interface{}) Errorx {
	return New(ErrDatabase.Code, err)
}

func InvalidParam(err interface{}) Errorx {
	return New(ErrInvalidParam.Code, err)
}

func Business(err interface{}) Errorx {
	return New(ErrBusiness.Code, err)
}
