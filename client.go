package gocef

/*
#include "include/capi/cef_client_capi.h"

extern cef_life_span_handler_t* gocef_client_get_life_span_handler(cef_client_t*);
*/
import "C"
import (
	"unsafe"
)

type Client interface {
	GetLifeSpanHandler() LifeSpanHandler
}

//export gocef_client_get_life_span_handler
func gocef_client_get_life_span_handler(c *C.cef_client_t) *C.cef_life_span_handler_t {
	self := gocefResolve(unsafe.Pointer(c)).(Client)
	handler := self.GetLifeSpanHandler()
	return lifeSpanHandlerDelegate{handler}.toNative()
}

type clientDelegate struct {
	Nativer
	self Client
}

func (c clientDelegate) copyToNative(p *C.cef_client_t) {
	p.get_life_span_handler = gocefFuncPtr(C.gocef_client_get_life_span_handler)
}

func (c clientDelegate) toNative() unsafe.Pointer {
	p := (*C.cef_client_t)(gocefNew(C.sizeof_cef_client_t))
	gocefBind(unsafe.Pointer(p), c.self)
	c.copyToNative(p)
	return p
}
