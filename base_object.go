package gocef

/*
#include "base_object.h"
*/
import "C"
import (
	"errors"
	"unsafe"
)

var (
	objectRefMap  = make(map[uintptr]interface{})
	errResolveRef = errors.New("failed to resolve reference")
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

func gocefSetRef(ptr unsafe.Pointer, obj interface{}) {
	objectRefMap[uintptr(ptr)] = obj
}

func gocefGetRef(ptr unsafe.Pointer) interface{} {
	obj, ok := objectRefMap[uintptr(ptr)]
	if !ok {
		panic(errResolveRef)
	}
	return obj
}

func gocefDelRef(ptr unsafe.Pointer) {
	_, ok := objectRefMap[uintptr(ptr)]
	if !ok {
		panic(errResolveRef)
	}
	delete(objectRefMap, uintptr(ptr))
}
