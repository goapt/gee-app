package response

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"app/errorx"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestErrorResponse_Render(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gee.CreateTestContext(w)
	ctx.HttpStatus = http.StatusInternalServerError

	resp := Error(ctx, ErrSystemError)
	resp.Render()
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, `{"code":"SystemError","msg":"系统错误"}`, w.Body.String())
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

func TestContext_ErrorFail(t *testing.T) {
	type args struct {
		err error
		msg string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "biz error", args: args{newStatusError(http.StatusBadRequest, errorx.New("CustomError", "this is constom error")), ""}, want: `{"code":"CustomError","msg":"this is constom error"}`},
		{name: "errorx exists", args: args{ErrInvalidParameter, "params error"}, want: `{"code":"InvalidParameter","msg":"params error"}`},
		{name: "errorx invalid", args: args{ErrInvalidSignature, ""}, want: `{"code":"InvalidSignature","msg":"签名错误"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gee.CreateTestContext(w)
			resp := Error(ctx, tt.args.err, tt.args.msg)
			resp.Render()
			assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

			got := w.Body.String()
			require.Equal(t, got, tt.want)
		})
	}
}
