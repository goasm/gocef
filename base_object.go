package gocef

/*
#include "base_object.h"
*/
import "C"
import (
	"unsafe"
)

func gocefNew(size int64) unsafe.Pointer {
	return C.gocef_new_impl(C.ulong(size))
}

func gocefAddRef(ptr unsafe.Pointer) {
	C.gocef_add_ref_impl(ptr)
}

func gocefRelease(ptr unsafe.Pointer) bool {
	retval := C.gocef_release_impl(ptr)
	return gocefToBool(retval)
}

func gocefHasOneRef(ptr unsafe.Pointer) bool {
	retval := C.gocef_has_one_ref_impl(ptr)
	return gocefToBool(retval)
}

func gocefHasAtLeastOneRef(ptr unsafe.Pointer) bool {
	retval := C.gocef_has_at_least_one_ref_impl(ptr)
	return gocefToBool(retval)
}
