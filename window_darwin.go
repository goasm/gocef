package gocef

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

void* cocoa_window_new() {
  NSRect wndRect = NSMakeRect(100, 100, 200, 100);
  NSUInteger wndStyle = (NSTitledWindowMask | NSClosableWindowMask |
                         NSResizableWindowMask | NSMiniaturizableWindowMask);
  NSWindow* wnd = [[NSWindow alloc] initWithContentRect:wndRect
                                    styleMask:wndStyle
                                    backing:NSBackingStoreBuffered
                                    defer:NO];
  return wnd;
}

*/
import "C"
import "unsafe"

func CreateWindow() unsafe.Pointer {
	return C.cocoa_window_new()
}
