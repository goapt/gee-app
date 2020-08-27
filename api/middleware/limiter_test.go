package middleware

import (
	"testing"

	"github.com/goapt/gee"
	"github.com/goapt/test"
	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	var testHandler gee.HandlerFunc = func(c *gee.Context) gee.Response {
		return c.JSON(gee.H{
			"code": 10000,
			"msg":  "success",
		})
	}

	mid := &Middleware{}
	limiter := mid.Limiter(2)
	for i := 0; i < 5; i++ {
		req := test.NewRequest("/dummy/impl", limiter, testHandler)
		resp, err := req.Get()
		assert.NoError(t, err)
		if i > 1 {
			assert.Equal(t, `{"code":"RateLimited","msg":"接口访问太频繁"}`, resp.GetBodyString())
		} else {
			assert.Equal(t, `{"code":10000,"msg":"success"}`, resp.GetBodyString())
		}
	}
}
