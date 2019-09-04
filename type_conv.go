package gocef

/*
#include "type_conv.h"
*/
import "C"
import "image/color"

func gocefToBool(v C.int) bool {
	return v != 0
}

func gocefFromBool(v bool) C.int {
	if v {
		return 1
	}
	return 0
}

func gocefToARGB(c color.Color) uint32 {
	if c == nil {
		return 0
	}
	r, g, b, a := c.RGBA()
	r &= 0xff00
	g &= 0xff00
	b &= 0xff00
	a &= 0xff00
	return (a<<24 | r<<16 | g<<8 | b<<0)
}

func gocefToUtf8(s string) *C.cef_string_utf8_t {
	return C.gocef_to_utf8_impl(C.CString(s), C.size_t(len(s)))
}

func gocefToUtf16(s string) *C.cef_string_utf16_t {
	return C.gocef_to_utf16_impl(C.CString(s), C.size_t(len(s)))
}
