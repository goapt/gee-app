package testutil

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/gin-gonic/gin"
)

func TestCallWebHandle(t *testing.T) {

	ctx := &gin.Context{}

	var p **gin.Engine = (**gin.Engine)(unsafe.Pointer(uintptr(unsafe.Pointer(&ctx.Keys)) - 8))
	fmt.Println("p", reflect.ValueOf(*p).IsNil())
	*p = gin.Default()
	fmt.Println("x", *p)

	v := reflect.ValueOf(ctx).Elem()
	field := v.FieldByName("engine")
	fmt.Println("field.UnsafeAddr()", unsafe.Pointer(field.UnsafeAddr()))
	newVal := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
	fmt.Println("enginPointer", newVal.Interface(), p)
}
