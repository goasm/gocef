package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_app_capi.h"
*/
import "C"
import (
	"os"
	"unsafe"
)

type mainArgs *C.cef_main_args_t
type settings *C.cef_settings_t
type app *C.cef_app_t

type Settings struct {
}

func fillMainArgs(args mainArgs) {
	argv := make([]*C.char, len(os.Args))
	for i := 0; i < len(os.Args); i++ {
		argv[i] = C.CString(os.Args[i])
	}
	args.argc = C.int(len(os.Args))
	args.argv = &argv[0]
}

func Initialize() {
	args := (mainArgs)(C.calloc(1, C.sizeof_cef_main_args_t))
	defer C.free(unsafe.Pointer(args))
	fillMainArgs(args)
	settings := (settings)(C.calloc(1, C.sizeof_cef_settings_t))
	defer C.free(unsafe.Pointer(settings))
	app := (app)(C.calloc(1, C.sizeof_cef_app_t))
	defer C.free(unsafe.Pointer(app))
	C.cef_initialize(args, settings, app, nil)
}

func RunMessageLoop() {
	C.cef_run_message_loop()
}

func QuitMessageLoop() {
	C.cef_quit_message_loop()
}

func Shutdown() {
	C.cef_shutdown()
}
