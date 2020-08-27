package errorx

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorx_Is(t *testing.T) {

	var (
		ErrDatabase = errors.New("数据库错误")
	)

	t.Run("wrap database", func(t *testing.T) {
		err := New("InvalidDatabase", ErrDatabase)
		assert.True(t, errors.Is(err, ErrDatabase))
		assert.EqualError(t, err, "数据库错误")
		err2 := New("InvalidDatabase", err)
		assert.True(t, errors.Is(err2, err))

		err3 := New("InvalidDatabase", sql.ErrNoRows)
		assert.True(t, errors.Is(err3, sql.ErrNoRows))
		assert.EqualError(t, err3, "sql: no rows in result set")
	})

	t.Run("coustom", func(t *testing.T) {
		err := New("InvalidEmail", "无效的邮箱")
		err2 := New(err.Code, "邮箱不能为空")
		assert.EqualError(t, err, "无效的邮箱")
		assert.True(t, errors.Is(err, err2))
	})

	t.Run("struct error", func(t *testing.T) {
		p := &struct {
			Id int
		}{
			Id: 1,
		}

		err := New("InvalidEmail", p)
		assert.EqualError(t, err, "[InvalidEmail]&{1}")
	})

	t.Run("unkonw error", func(t *testing.T) {
		err := New("InvalidEmail", "")
		assert.EqualError(t, err, "[InvalidEmail]unknow error")
	})

	t.Run("pointer error", func(t *testing.T) {
		err := New("EOF", "EOF")
		err2 := New("EOF", "EOF")
		// 即便msg相等，由于连个错误指针不同，因此不能直接判断相等
		assert.False(t, err == err2)
		// errorx 可以仅仅通过code来判定两个错误是相同的
		assert.True(t, errors.Is(err, err2))

		err3 := errors.New("EOF")
		err4 := errors.New("EOF")
		assert.False(t, err3 == err4)
		// 官方错误更加严格，必须要指针和错误信息都相等才能Is
		assert.False(t, errors.Is(err4, err3))
	})
}

func TestPrettyNoRows(t *testing.T) {
	err := PrettyNoRows(sql.ErrNoRows, "订单不存在")
	assert.EqualError(t, err, "订单不存在")
	err2 := PrettyNoRows(errors.New("not sql error"), "不应该替换")
	assert.EqualError(t, err2, "not sql error")
	err3 := PrettyNoRows(nil, "is nil")
	assert.Nil(t, err3)
	err4 := PrettyNoRows(sql.ErrNoRows, errors.New("new error"))
	assert.EqualError(t, err4, "new error")
	err5 := PrettyNoRows(sql.ErrNoRows, 123)
	assert.EqualError(t, err5, "123")
}

func TestIsSystemError(t *testing.T) {
	err := FilterNoRows(nil)
	assert.Nil(t, err)
	err2 := FilterNoRows(sql.ErrNoRows)
	assert.Nil(t, err2)
	err3 := FilterNoRows(errors.New("not sql error"))
	assert.EqualError(t, err3, "not sql error")
}

func TestIsSystemError1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "err1",
			args: args{sql.ErrConnDone},
			want: true,
		},
		{
			name: "err1",
			args: args{sql.ErrTxDone},
			want: true,
		},
		{
			name: "err1",
			args: args{driver.ErrBadConn},
			want: true,
		},
		{
			name: "err1",
			args: args{&url.Error{}},
			want: true,
		},
		{
			name: "err1",
			args: args{errors.New("not system error")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSystemError(tt.args.err); got != tt.want {
				t.Errorf("IsSystemError() = %v, want %v", got, tt.want)
			}
		})
	}
}
