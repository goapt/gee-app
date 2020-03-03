package handler

import (
	"testing"

	"app/test"
)

func TestLoginHandle(t *testing.T) {
	req := test.NewJsonRequest("/login", map[string]interface{}{"user_name": "test", "password": "123456"})
	resp := test.Run(t, req)
	test.IsSuccess(t, resp)
}
