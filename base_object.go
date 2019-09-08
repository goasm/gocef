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
