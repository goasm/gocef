package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_app_capi.h"
*/
import "C"
import "unsafe"

type mainArg *C.cef_main_args_t

func Initialize() error {
	args := (mainArg)(C.malloc(C.size_t(C.sizeof_cef_main_args_t)))
	defer C.free(unsafe.Pointer(args))
	C.cef_initialize(args)
}
