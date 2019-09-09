package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_browser_capi.h"
*/
import "C"

type Browser C.cef_browser_t

func CreateBrowser(client Client) bool {
	wndInfo := (*C.cef_window_info_t)(C.calloc(1, C.sizeof_cef_window_info_t))
	// defer C.free(unsafe.Pointer(wndInfo))
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_cef_browser_settings_t))
	// defer C.free(unsafe.Pointer(settings))
	// init data
	// wndInfo.parent_window = C.cef_window_handle_t(uintptr(hwnd))
	wndInfo.width = 500
	wndInfo.height = 300
	url := gocefToUtf16("https://www.baidu.com")
	retval := C.cef_browser_host_create_browser(wndInfo, clientDelegate{client}.toNative(), url, settings, nil)
	return gocefToBool(retval)
}

func CreateBrowserSync(client Client) *Browser {
	C.cef_browser_host_create_browser_sync(nil, nil, nil, nil, nil)
	return nil
}
