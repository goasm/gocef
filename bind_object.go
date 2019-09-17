package gocef

import (
	"errors"
	"unsafe"
)

var (
	objResolveMap = make(map[uintptr]Nativer)
	errResolveObj = errors.New("failed to resolve object")
)

func gocefBind(ptr unsafe.Pointer, obj Nativer) {
	objResolveMap[uintptr(ptr)] = obj
}

func gocefResolve(ptr unsafe.Pointer) Nativer {
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
