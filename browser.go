package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_client_capi.h"
*/
import "C"

type Browser C.cef_browser_t

func CreateBrowser(windowInfo *WindowInfo, client *ClientDelegate) bool {
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_cef_browser_settings_t))
	// defer C.free(unsafe.Pointer(settings))
	url := C.cef_string_userfree_utf16_alloc()
	gocefSetUtf16(url, "https://www.baidu.com")
	retval := C.cef_browser_host_create_browser(
		(*C.cef_window_info_t)(windowInfo.toNative()),
		(*C.cef_client_t)(client.toNative()),
		url,
		settings,
		nil,
	)
	return gocefToBool(retval)
}

func CreateBrowserSync(windowInfo *WindowInfo, client Client) *Browser {
	C.cef_browser_host_create_browser_sync(nil, nil, nil, nil, nil)
	return nil
}
