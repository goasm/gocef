package gocef

/*
#include "include/capi/cef_life_span_handler_capi.h"
*/
import "C"

type LifeSpanHandler interface {
	OnBeforeClose(b *Browser)
}

type lifeSpanHandlerDelegate struct {
	self LifeSpanHandler
}

func (h lifeSpanHandlerDelegate) toNative() *C.cef_life_span_handler_t {
	p := (*C.cef_life_span_handler_t)(gocefNew(C.sizeof_cef_life_span_handler_t))
	return p
}
