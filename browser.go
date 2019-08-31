package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_client_capi.h"
#include "base_object.h"
#include "string_utils.h"
*/
import "C"

type browser *C.cef_browser_t

func CreateBrowser() bool {
	wndInfo := (*C.cef_window_info_t)(C.calloc(1, C.sizeof_cef_window_info_t))
	// defer C.free(unsafe.Pointer(wndInfo))
	client := (*C.cef_client_t)(C.gocef_new(C.sizeof_cef_client_t))
	url := (*C.cef_string_t)(C.calloc(1, C.sizeof_cef_string_t))
	// defer C.free(unsafe.Pointer(url))
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_cef_browser_settings_t))
	// defer C.free(unsafe.Pointer(settings))
	// init data
	wndInfo.width = 500
	wndInfo.height = 300
	C.gocef_set_string_utf16("https://www.baidu.com", url)
	retval := C.cef_browser_host_create_browser(wndInfo, client, url, settings, nil)
	return retval != 0
}

func CreateBrowserSync() browser {
	return C.cef_browser_host_create_browser_sync(nil, nil, nil, nil, nil)
}
