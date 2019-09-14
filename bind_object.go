package gocef

import (
	"errors"
	"unsafe"
)

var (
	objResolveMap = make(map[uintptr]interface{})
	errResolveObj = errors.New("failed to resolve object")
)

func gocefBind(ptr unsafe.Pointer, obj interface{}) {
	objResolveMap[uintptr(ptr)] = obj
}

func gocefResolve(ptr unsafe.Pointer) interface{} {
	obj, ok := objResolveMap[uintptr(ptr)]
	if !ok {
		panic(errResolveObj)
	}
	return obj
}

func gocefUnbind(ptr unsafe.Pointer) {
	_, ok := objResolveMap[uintptr(ptr)]
	if !ok {
		panic(errResolveObj)
	}
	delete(objResolveMap, uintptr(ptr))
}
