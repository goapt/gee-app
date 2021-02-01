package middleware

import (
	"testing"

	"github.com/goapt/gee"
	"github.com/goapt/test"
	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	var testHandler gee.HandlerFunc = func(c *gee.Context) gee.Response {
		panic("dummy")

		return c.JSON(gee.H{
			"code": 10000,
			"msg":  "success",
		})
	}

	assert.NotPanics(t, func() {
		req := test.NewRequest("/dummy/impl", gee.HandlerFunc(NewRecover()), testHandler)
		resp, err := req.Get()
		assert.NoError(t, err)
		assert.Equal(t, 500, resp.Code)
	})
}
