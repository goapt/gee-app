package handler

import (
	"testing"

	"app/testutil"
)

func TestLoginHandle(t *testing.T) {
	req := testutil.NewRequest(t, `{"user_name":"test","password":"123456"}`)
	testutil.Run(t, LoginHandle, req)
}
