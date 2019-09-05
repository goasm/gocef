package gocef

/*
#include "include/capi/cef_client_capi.h"
#include "base_object.h"
*/
import "C"

type Client struct {
}

func (s *Client) toNative() *C.cef_client_t {
	p := (*C.cef_client_t)(C.gocef_new(C.sizeof_cef_client_t))
	return p
}
