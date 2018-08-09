package tool

import "unsafe"

func ByteTostring(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// 返された []byte は容量の情報が不定です。
// 絶対にappend()の対象にしないでください。
