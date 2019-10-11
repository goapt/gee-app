package handler

import (
	"testing"

	"app/testutil"
)

func TestUserHandle(t *testing.T)  {
	req := testutil.NewRequest(t, "")
	req.Header.Set("Access-Token", "4xgCqZpNHGyEwSHshM6ocg==")
	testutil.Run(t, LoginHandle, req)
}