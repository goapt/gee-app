package handler

import (
	"testing"

	"app/testutil"
)

func TestUserHandle(t *testing.T) {
	req := testutil.NewRequest("/api/user", "")
	req.Header.Set("Access-Token", "4xgCqZpNHGyEwSHshM6ocg==")
	resp := testutil.Run(t, req)
	testutil.IsSuccess(t, resp)
}
