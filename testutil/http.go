package testutil

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"unsafe"

	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"
)

func NewUrlEncodeRequest(t *testing.T, body string) *http.Request {
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequest("POST", "", buf)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func NewRequest(t *testing.T, body string) *http.Request {
	buf := bytes.NewBufferString(body)
	req, err := http.NewRequest("POST", "", buf)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func Run(t *testing.T, h gee.IHandler, req *http.Request) {
	buf := &bytes.Buffer{}
	ctx := &gin.Context{
		Writer: &BufferWriter{
			buf:    buf,
			header: http.Header{},
		},
		Request: req,
	}
	// 设置私有变量 engine 的值
	v := reflect.ValueOf(ctx).Elem()
	field := v.FieldByName("engine")
	newVal := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
	newVal.Set(reflect.ValueOf(gin.Default()))

	gee.Handle(h)(ctx)
	fmt.Println("API RESPONSE:\n", buf.String())
}