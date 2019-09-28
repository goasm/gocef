package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_client_capi.h"
*/
import "C"

type Browser struct {
	Nativer
	cref *C.cef_browser_t
}

func CreateBrowser(windowInfo *WindowInfo, client *ClientDelegate, url string) bool {
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_cef_browser_settings_t))
	// defer C.free(unsafe.Pointer(settings))
	tmp1 := C.cef_string_userfree_utf16_alloc()
	gocefSetUtf16(tmp1, url)
	retval := C.cef_browser_host_create_browser(
		(*C.cef_window_info_t)(windowInfo.toNative()),
		(*C.cef_client_t)(client.toNative()),
		tmp1,
		settings,
		nil,
	)
	C.cef_string_userfree_utf16_free(tmp1)
	return gocefToBool(retval)
}

func CreateBrowserSync(windowInfo *WindowInfo, client *ClientDelegate, url string) *Browser {
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_cef_browser_settings_t))
	// defer C.free(unsafe.Pointer(settings))
	tmp1 := C.cef_string_userfree_utf16_alloc()
	gocefSetUtf16(tmp1, url)
	cref := C.cef_browser_host_create_browser_sync(
		(*C.cef_window_info_t)(windowInfo.toNative()),
		(*C.cef_client_t)(client.toNative()),
		tmp1,
		settings,
		nil,
	)
	C.cef_string_userfree_utf16_free(tmp1)
	return &Browser{cref: cref}
}
