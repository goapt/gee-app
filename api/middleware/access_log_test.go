package middleware

import (
	"bytes"
	"testing"

	"github.com/goapt/gee"
	"github.com/goapt/logger"
	"github.com/goapt/test"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestAccessLog(t *testing.T) {
	var testHandler gee.HandlerFunc = func(c *gee.Context) gee.Response {
		return c.JSON(gee.H{
			"code": 10000,
			"msg":  "success",
		})
	}

	byteBuf := new(bytes.Buffer)
	lg := logger.NewLogger(func(c *logger.Config) {
		c.LogMode = "custom"
		c.LogWriter = byteBuf
	})

	req := test.NewRequest("/dummy/impl", gee.HandlerFunc(NewAccessLog(lg)), testHandler)
	resp, err := req.JSON(map[string]interface{}{
		"user_id": 123,
	})
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.Code)

	l := gjson.Parse(byteBuf.String())
	assert.Equal(t, `{"user_id":123}`, l.Get("request_body").String())
	assert.Equal(t, `{"code":10000,"msg":"success"}`, l.Get("response_body").String())
}
