package gocef

/*
#include "include/capi/cef_browser_capi.h"
*/
import "C"

type browser *C.cef_browser_t

func CreateBrowser() browser {
	return C.cef_browser_host_create_browser_sync(nil, nil, nil, nil, nil)
}
