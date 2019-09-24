package gocef

/*
#include "type_conv.h"
*/
import "C"
import (
	"image/color"
)

type gocefFuncPtr *[0]byte

// gocefToBool converts a value of C int to Go bool
func gocefToBool(v C.int) bool {
	return v != 0
}

// gocefToInt converts a value of Go bool to C int
func gocefToInt(v bool) C.int {
	if v {
		return 1
	}
	return 0
}

// gocefToARGB converts a Go color to cef_color
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

// gocefToUtf8 converts a Go string to cef_string_utf8
func gocefToUtf8(s string) C.cef_string_utf8_t {
	return C.gocef_to_utf8_impl(C.CString(s), C.size_t(len(s)))
}

// gocefToUtf16 converts a Go string to cef_string_utf16
func gocefToUtf16(s string) C.cef_string_utf16_t {
	return C.gocef_to_utf16_impl(C.CString(s), C.size_t(len(s)))
}
