package gocef

/*
#include "include/capi/cef_life_span_handler_capi.h"

extern void gocef_life_span_handler_on_before_close(cef_life_span_handler_t*, cef_browser_t*);
*/
import "C"
import (
	"unsafe"
)

// LifeSpanHandler is the interface of browser life span event handler
type LifeSpanHandler interface {
	OnBeforeClose(b *Browser)
}

//export gocef_life_span_handler_on_before_close
func gocef_life_span_handler_on_before_close(h *C.cef_life_span_handler_t, b *C.cef_browser_t) {
	self := gocefResolve(unsafe.Pointer(h)).(*LifeSpanHandlerDelegate).self
	self.OnBeforeClose((*Browser)(b))
}

// LifeSpanHandlerDelegate delegates all methods of LifeSpanHandler
type LifeSpanHandlerDelegate struct {
	Nativer
	cref *C.cef_life_span_handler_t
	self LifeSpanHandler
}

// NewLifeSpanHandler creates a delegate of LifeSpanHandler
func NewLifeSpanHandler(h LifeSpanHandler) *LifeSpanHandlerDelegate {
	d := &LifeSpanHandlerDelegate{self: h}
	d.cref = (*C.cef_life_span_handler_t)(d.allocate())
	return d
}

func (h *LifeSpanHandlerDelegate) allocate() unsafe.Pointer {
	p := gocefNew(C.sizeof_cef_life_span_handler_t)
	gocefBind(p, h)
	return p
}

func (h *LifeSpanHandlerDelegate) copyToNative(p *C.cef_life_span_handler_t) {
	p.on_before_close = gocefFuncPtr(C.gocef_life_span_handler_on_before_close)
}

func (h *LifeSpanHandlerDelegate) toNative() unsafe.Pointer {
	h.copyToNative(h.cref)
	return unsafe.Pointer(h.cref)
}

// Destroy destroys the delegate and free the C object
func (h *LifeSpanHandlerDelegate) Destroy() bool {
	result := gocefRelease(unsafe.Pointer(h.cref))
	h.cref = nil
	return result
}
