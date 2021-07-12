package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goapt/gee"
	"github.com/stretchr/testify/assert"
)

func init() {
	gee.SetMode(gee.TestMode)
}

func TestErrorResponse_Render(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gee.CreateTestContext(w)
	resp := WithStatusError(ctx, http.StatusInternalServerError, "SystemError", "系统错误")
	resp.Render()
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"code":"SystemError","msg":"系统错误"}`, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}
