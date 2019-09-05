package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_client_capi.h"
#include "base_object.h"
*/
import "C"
import (
	"unsafe"
)

type browser *C.cef_browser_t

func CreateBrowser(hwnd unsafe.Pointer) bool {
	wndInfo := (*C.cef_window_info_t)(C.calloc(1, C.sizeof_cef_window_info_t))
	// defer C.free(unsafe.Pointer(wndInfo))
	client := (*C.cef_client_t)(C.gocef_new(C.sizeof_cef_client_t))
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_cef_browser_settings_t))
	// defer C.free(unsafe.Pointer(settings))
	// init data
	wndInfo.parent_window = C.cef_window_handle_t(uintptr(hwnd))
	wndInfo.width = 500
	wndInfo.height = 300
	url := gocefToUtf16("https://www.baidu.com")
	retval := C.cef_browser_host_create_browser(wndInfo, client, url, settings, nil)
	return gocefToBool(retval)
}

func CreateBrowserSync(hwnd unsafe.Pointer) browser {
	return C.cef_browser_host_create_browser_sync(nil, nil, nil, nil, nil)
}
