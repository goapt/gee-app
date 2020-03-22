package handler

import (
	"testing"

	"app/testutil"
)

func TestLoginHandle(t *testing.T) {
	req := testutil.NewJsonRequest("/login", map[string]interface{}{"user_name": "test", "password": "123456"})
	resp := testutil.Run(t, req)
	testutil.IsSuccess(t, resp)
}
