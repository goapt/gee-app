package handler

import (
	"testing"

	"app/test"
)

func TestUserHandle(t *testing.T) {
	req := test.NewRequest("/api/user", "")
	req.Header.Set("Access-Token", "4xgCqZpNHGyEwSHshM6ocg==")
	resp := test.Run(t, req)
	test.IsSuccess(t, resp)
}
