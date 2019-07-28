package gocef

/*
#include "include/capi/cef_app_capi.h"
*/
import "C"

func Initialize() error {
	C.cef_initialize()
}
