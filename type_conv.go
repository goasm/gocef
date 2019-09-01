package gocef

/*
#include "type_conv.h"
*/
import "C"

func gocefToBool(v C.int) bool {
	return v != 0
}

func gocefFromBool(v bool) C.int {
	if v {
		return 1
	}
	return 0
}

func gocefToUtf8(s string) *C.cef_string_utf8_t {
	return C.gocef_to_utf8_impl(C.CString(s), C.size_t(len(s)))
}

func gocefToUtf16(s string) *C.cef_string_utf16_t {
	return C.gocef_to_utf16_impl(C.CString(s), C.size_t(len(s)))
}
