package gocef

/*
#include <stdlib.h>
#include "include/internal/cef_string.h"

cef_string_utf8_t* gocef_to_utf8_impl(char* str, size_t len) {
  cef_string_utf8_t* p = calloc(1, sizeof(cef_string_utf8_t));
  p->str = str;
  p->length = len;
  return p;
}

cef_string_utf16_t* gocef_to_utf16_impl(char* str, size_t len) {
  cef_string_utf16_t* p = calloc(1, sizeof(cef_string_utf16_t));
  cef_string_utf8_to_utf16(str, len, p);
  return p;
}

*/
import "C"
import (
	"image/color"
	"unsafe"
)

func gocefToBool(v C.int) bool {
	return v != 0
}

func gocefFromBool(v bool) C.int {
	if v {
		return 1
	}
	return 0
}

func gocefToFuncPtr(p unsafe.Pointer) *[0]byte {
	return (*[0]byte)(p)
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
