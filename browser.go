package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_browser_capi.h"
#include "include/capi/cef_client_capi.h"

cef_browser_host_t* gocef_browser_get_host(cef_browser_t* self) {
  return self->get_host(self);
}

int gocef_browser_can_go_back(cef_browser_t* self) {
  return self->can_go_back(self);
}

void gocef_browser_go_back(cef_browser_t* self) {
  self->go_back(self);
}

int gocef_browser_can_go_forward(cef_browser_t* self) {
  return self->can_go_forward(self);
}

void gocef_browser_go_forward(cef_browser_t* self) {
  self->go_forward(self);
}

int gocef_browser_is_loading(cef_browser_t* self) {
  return self->is_loading(self);
}

void gocef_browser_reload(cef_browser_t* self) {
  self->reload(self);
}

void gocef_browser_reload_ignore_cache(cef_browser_t* self) {
  self->reload_ignore_cache(self);
}

void gocef_browser_stop_load(cef_browser_t* self) {
  self->stop_load(self);
}

int gocef_browser_get_identifier(cef_browser_t* self) {
  return self->get_identifier(self);
}

int gocef_browser_is_same(cef_browser_t* self, cef_browser_t* that) {
  return self->is_same(self, that);
}

int gocef_browser_is_popup(cef_browser_t* self) {
  return self->is_popup(self);
}

int gocef_browser_has_document(cef_browser_t* self) {
  return self->has_document(self);
}

cef_frame_t* gocef_browser_get_main_frame(cef_browser_t* self) {
  return self->get_main_frame(self);
}

cef_frame_t* gocef_browser_get_focused_frame(cef_browser_t* self) {
  return self->get_focused_frame(self);
}

cef_frame_t* gocef_browser_get_frame_byident(cef_browser_t* self, int64 identifier) {
  return self->get_frame_byident(self, identifier);
}

cef_frame_t* gocef_browser_get_frame(cef_browser_t* self, const cef_string_t* name) {
  return self->get_frame(self, name);
}

size_t gocef_browser_get_frame_count(cef_browser_t* self) {
  return self->get_frame_count(self);
}

void gocef_browser_get_frame_identifiers(cef_browser_t* self, size_t* identifiersCount, int64* identifiers) {
  self->get_frame_identifiers(self, identifiersCount, identifiers);
}

void gocef_browser_get_frame_names(cef_browser_t* self, cef_string_list_t names) {
  self->get_frame_names(self, names);
}

int gocef_browser_send_process_message(cef_browser_t* self, cef_process_id_t target_process, cef_process_message_t* message) {
  return self->send_process_message(self, target_process, message);
}
*/
import "C"

type BrowserHost C.cef_browser_host_t

type Browser struct {
	Nativer
	cref *C.cef_browser_t
}

func CreateBrowser(windowInfo *WindowInfo, client *ClientDelegate, url string) bool {
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_cef_browser_settings_t))
	// defer C.free(unsafe.Pointer(settings))
	tmp1 := C.cef_string_userfree_utf16_alloc()
	gocefSetUtf16(tmp1, url)
	retval := C.cef_browser_host_create_browser(
		(*C.cef_window_info_t)(windowInfo.toNative()),
		(*C.cef_client_t)(client.toNative()),
		tmp1,
		settings,
		nil,
	)
	C.cef_string_userfree_utf16_free(tmp1)
	return gocefToBool(retval)
}

func CreateBrowserSync(windowInfo *WindowInfo, client *ClientDelegate, url string) *Browser {
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_cef_browser_settings_t))
	// defer C.free(unsafe.Pointer(settings))
	tmp1 := C.cef_string_userfree_utf16_alloc()
	gocefSetUtf16(tmp1, url)
	cref := C.cef_browser_host_create_browser_sync(
		(*C.cef_window_info_t)(windowInfo.toNative()),
		(*C.cef_client_t)(client.toNative()),
		tmp1,
		settings,
		nil,
	)
	C.cef_string_userfree_utf16_free(tmp1)
	return &Browser{cref: cref}
}

func (b *Browser) GetHost() *BrowserHost {
	retval := C.gocef_browser_get_host(b.cref)
	return (*BrowserHost)(retval)
}

func (b *Browser) CanGoBack() bool {
	retval := C.gocef_browser_can_go_back(b.cref)
	return gocefToBool(retval)
}

func (b *Browser) GoBack() {
	C.gocef_browser_go_back(b.cref)
}

func (b *Browser) CanGoForward() bool {
	retval := C.gocef_browser_can_go_forward(b.cref)
	return gocefToBool(retval)
}

func (b *Browser) GoForward() {
	C.gocef_browser_go_forward(b.cref)
}

func (b *Browser) IsLoading() bool {
	retval := C.gocef_browser_is_loading(b.cref)
	return gocefToBool(retval)
}

func (b *Browser) Reload() {
	C.gocef_browser_reload(b.cref)
}

func (b *Browser) ReloadIgnoreCache() {
	C.gocef_browser_reload_ignore_cache(b.cref)
}

func (b *Browser) StopLoad() {
	C.gocef_browser_stop_load(b.cref)
}

func (b *Browser) GetIdentifier() int {
	retval := C.gocef_browser_get_identifier(b.cref)
	return int(retval)
}

func (b *Browser) IsSame(that *Browser) bool {
	retval := C.gocef_browser_is_same(b.cref, that.cref)
	return gocefToBool(retval)
}

func (b *Browser) IsPopup() bool {
	retval := C.gocef_browser_is_popup(b.cref)
	return gocefToBool(retval)
}

func (b *Browser) HasDocument() bool {
	retval := C.gocef_browser_has_document(b.cref)
	return gocefToBool(retval)
}

// func (b *Browser) GetMainFrame() *Frame {
// }

// func (b *Browser) GetFocusedFrame() *Frame {
// }

// func (b *Browser) GetFrameByident(identifier int64) *Frame {
// }

// func (b *Browser) GetFrame(name string) *Frame {
// }

// func (b *Browser) GetFrameCount() int {
// }

// func (b *Browser) GetFrameIdentifiers(identifiersCount *int, identifiers *[]int64) {
// }

// func (b *Browser) GetFrameNames(names *[]string) {
// }

// func (b *Browser) SendProcessMessage(target_process cef_process_id_t, message *cef_process_message_t) bool {
// }
