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

// ExecuteProcess executes a secondary process using the command-line arguments,
// it will block until the process should exit and then return the exit code.
func ExecuteProcess() int {
	retval := C.cef_execute_process(mainArgs, nil, nil)
	return int(retval)
}

// Initialize initializes the CEF browser process
func Initialize() bool {
	settings := Settings{}
	settings.NoSandbox = true
	app := (*C.cef_app_t)(gocefNew(C.sizeof_cef_app_t))
	retval := C.cef_initialize(mainArgs, settings.toNative(), app, nil)
	return gocefToBool(retval)
}

// RunMessageLoop runs the CEF message loop
func RunMessageLoop() {
	C.cef_run_message_loop()
}

// QuitMessageLoop quits the CEF message loop that was started by RunMessageLoop
func QuitMessageLoop() {
	C.cef_quit_message_loop()
}

// Shutdown shuts down the CEF browser process
func Shutdown() {
	C.cef_shutdown()
}
