package gocef

/*
#include <gtk/gtk.h>
#cgo pkg-config: gtk+-3.0
*/
import "C"
import (
	"unsafe"
)

func init() {
	C.gtk_init(nil, nil)
}

func CreateWindow() unsafe.Pointer {
	hwnd := C.gtk_window_new(C.GTK_WINDOW_TOPLEVEL)
	return unsafe.Pointer(hwnd)
}
