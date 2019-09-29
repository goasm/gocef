package gocef

/*
#cgo pkg-config: gtk+-3.0
#include <gtk/gtk.h>
#include <gdk/gdkx.h>

Window gocef_get_window_xid(GtkWidget* window) {
  return GDK_WINDOW_XID(gtk_widget_get_window(window));
}

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
	xid := uintptr(C.gocef_get_window_xid(hwnd))
	return unsafe.Pointer(xid)
}
