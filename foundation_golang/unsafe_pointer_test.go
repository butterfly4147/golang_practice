package foundation_golang

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestUnsafePointer(t *testing.T) {
	bs := []byte("Hello, world!")
	s := toString(bs)

	fmt.Printf("bs: %x\n", &bs)
	fmt.Printf("s: %x\n", &s)
}

func toString(bs []byte) string {
	//此特定情况（[]byte和string头结构“部分相同”）下，通过非安全指针类型转化，实现类型“变更”
	//避免了底层数组复制，可以有效改善执行性能
	return *(*string)(unsafe.Pointer(&bs))
}
