package gocef

/*
#include "include/capi/cef_client_capi.h"

extern cef_life_span_handler_t* gocef_client_get_life_span_handler(cef_client_t*);
*/
import "C"
import (
	"unsafe"
)

// Client is the interface of handler implementations
type Client interface {
	GetLifeSpanHandler() LifeSpanHandler
}

//export gocef_client_get_life_span_handler
func gocef_client_get_life_span_handler(c *C.cef_client_t) *C.cef_life_span_handler_t {
	self := gocefResolve(unsafe.Pointer(c)).(*ClientDelegate).self
	handler := self.GetLifeSpanHandler()
	handlerDelegate := lifeSpanHandlerDelegate{self: handler}
	return (*C.cef_life_span_handler_t)(handlerDelegate.toNative())
}

// NewClient creates a delegate of Client
func NewClient(c Client) *ClientDelegate {
	return &ClientDelegate{self: c}
}

// ClientDelegate delegates all methods of Client
type ClientDelegate struct {
	Nativer
	cref *C.cef_client_t
	self Client
}

func (c *ClientDelegate) copyToNative(p *C.cef_client_t) {
	p.get_life_span_handler = gocefFuncPtr(C.gocef_client_get_life_span_handler)
}

func (c *ClientDelegate) toNative() unsafe.Pointer {
	if c.cref == nil {
		c.cref = (*C.cef_client_t)(gocefNew(C.sizeof_cef_client_t))
		gocefBind(unsafe.Pointer(c.cref), c)
		c.copyToNative(c.cref)
	}
	return unsafe.Pointer(c.cref)
}
