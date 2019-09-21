package gocef

/*
#include "include/capi/cef_app_capi.h"
*/
import "C"
import (
	"os"
)

// ExecuteProcess executes a secondary process using the command-line arguments,
// it will block until the process should exit and then return the exit code.
func ExecuteProcess() int {
	args := mainArgs(os.Args)
	retval := C.cef_execute_process(args.toNative(), nil, nil)
	return int(retval)
}

// Initialize initializes the CEF browser process
func Initialize() bool {
	args := mainArgs(os.Args)
	settings := &Settings{}
	settings.NoSandbox = true
	retval := C.cef_initialize(
		args.toNative(),
		(*C.cef_settings_t)(settings.toNative()),
		nil,
		nil,
	)
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
