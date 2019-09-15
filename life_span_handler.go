package gocef

/*
#include "include/capi/cef_life_span_handler_capi.h"

extern void gocef_life_span_handler_on_before_close(cef_life_span_handler_t*, cef_browser_t*);
*/
import "C"
import (
	"unsafe"
)

type LifeSpanHandler interface {
	OnBeforeClose(b *Browser)
}

//export gocef_life_span_handler_on_before_close
func gocef_life_span_handler_on_before_close(h *C.cef_life_span_handler_t, b *C.cef_browser_t) {
	self := gocefResolve(unsafe.Pointer(h)).(LifeSpanHandler)
	self.OnBeforeClose((*Browser)(b))
}

type lifeSpanHandlerDelegate struct {
	self LifeSpanHandler
}

func (h lifeSpanHandlerDelegate) copyToNative(p *C.cef_life_span_handler_t) {
	p.on_before_close = gocefFuncPtr(C.gocef_life_span_handler_on_before_close)
}

func (h lifeSpanHandlerDelegate) toNative() *C.cef_life_span_handler_t {
	p := (*C.cef_life_span_handler_t)(gocefNew(C.sizeof_cef_life_span_handler_t))
	gocefBind(unsafe.Pointer(p), h.self)
	h.copyToNative(p)
	return p
}
