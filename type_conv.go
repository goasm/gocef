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

// gocefSetUtf8 converts a Go string to cef_string_utf8
func gocefSetUtf8(d *C.cef_string_utf8_t, s string) {
	C.gocef_set_utf8_impl(d, C.CString(s), C.size_t(len(s)))
}

// gocefSetUtf16 converts a Go string to cef_string_utf16
func gocefSetUtf16(d *C.cef_string_utf16_t, s string) {
	C.gocef_set_utf16_impl(d, C.CString(s), C.size_t(len(s)))
}
