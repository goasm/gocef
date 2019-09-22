package gocef

import (
	"unsafe"
)

// Nativer represents an abstraction for CEF native object
type Nativer interface {
	allocate() unsafe.Pointer
	toNative() unsafe.Pointer
	Destroy() bool
}
