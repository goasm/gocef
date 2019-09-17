package gocef

import (
	"unsafe"
)

// Nativer represents an abstraction for CEF native object
type Nativer interface {
	toNative() unsafe.Pointer
}
