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
	self := gocefGetRef(unsafe.Pointer(c)).(Client)
	handler := self.GetLifeSpanHandler()
	return lifeSpanHandlerDelegate{handler}.toNative()
}

type clientDelegate struct {
	self Client
}

func (c clientDelegate) copyToNative(p *C.cef_client_t) {
	p.get_life_span_handler = gocefToFuncPtr(C.gocef_client_get_life_span_handler)
}

func (c clientDelegate) toNative() *C.cef_client_t {
	p := (*C.cef_client_t)(gocefNew(C.sizeof_cef_client_t))
	gocefSetRef(unsafe.Pointer(p), c.self)
	c.copyToNative(p)
	return p
}
