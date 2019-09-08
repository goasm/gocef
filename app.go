package gocef

/*
#include <stdlib.h>
#include "include/capi/cef_app_capi.h"
*/
import "C"
import (
	"os"
)

var mainArgs *C.cef_main_args_t

func init() {
	mainArgs = (*C.cef_main_args_t)(C.calloc(1, C.sizeof_cef_main_args_t))
	fillMainArgs(mainArgs)
}

func fillMainArgs(args *C.cef_main_args_t) {
	argv := make([]*C.char, len(os.Args))
	for i := 0; i < len(os.Args); i++ {
		argv[i] = C.CString(os.Args[i])
	}
	args.argc = C.int(len(os.Args))
	args.argv = &argv[0]
}

func ExecuteProcess() int {
	retval := C.cef_execute_process(mainArgs, nil, nil)
	return int(retval)
}

func Initialize() bool {
	settings := Settings{}
	settings.NoSandbox = true
	app := (*C.cef_app_t)(gocefNew(C.sizeof_cef_app_t))
	retval := C.cef_initialize(mainArgs, settings.toNative(), app, nil)
	return gocefToBool(retval)
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
