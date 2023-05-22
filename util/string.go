package util

import "unsafe"

// Str2byte string转bytes
func Str2byte(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s                                                  // 把s的地址付给b
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s) // 修改容量为长度
	return
}

// Byte2str []byte转string
func Byte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
