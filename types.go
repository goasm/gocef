package gocef

/*
#include <stdlib.h>
#include "include/internal/cef_types.h"
*/
import "C"
import (
	"image/color"
	"unsafe"
)

type mainArgs []string

func (m mainArgs) toNative() *C.cef_main_args_t {
	p := (*C.cef_main_args_t)(C.calloc(1, C.sizeof_cef_main_args_t))
	p.argc = C.int(len(m))
	args := make([]*C.char, len(m))
	for i := 0; i < len(m); i++ {
		args[i] = C.CString(m[i])
	}
	p.argv = &args[0]
	return p
}

type LogSeverity int

const (
	LogSeverityDefault LogSeverity = iota
	LogSeverityVerbose
	LogSeverityInfo
	LogSeverityWarning
	LogSeverityError
	LogSeverityFatal
	LogSeverityDebug   LogSeverity = LogSeverityVerbose
	LogSeverityDisable LogSeverity = 99
)

// Settings represents the initialization settings of process
type Settings struct {
	Nativer
	cref                        *C.cef_settings_t
	NoSandbox                   bool
	BrowserSubprocessPath       string
	FrameworkDirPath            string
	MultiThreadedMessageLoop    bool
	ExternalMessagePump         bool
	WindowlessRenderingEnabled  bool
	CommandLineArgsDisabled     bool
	CachePath                   string
	UserDataPath                string
	PersistSessionCookies       bool
	PersistUserPreferences      bool
	UserAgent                   string
	ProductVersion              string
	Locale                      string
	LogFile                     string
	LogSeverity                 LogSeverity
	JavascriptFlags             string
	ResourcesDirPath            string
	LocalesDirPath              string
	PackLoadingDisabled         bool
	RemoteDebuggingPort         int
	UncaughtExceptionStackSize  int
	IgnoreCertificateErrors     bool
	EnableNetSecurityExpiration bool
	BackgroundColor             color.Color
	AcceptLanguageList          string
}

func (s *Settings) copyToNative(p *C.cef_settings_t) {
	p.no_sandbox = gocefToInt(s.NoSandbox)
	gocefSetUtf16(&p.browser_subprocess_path, s.BrowserSubprocessPath)
	gocefSetUtf16(&p.framework_dir_path, s.FrameworkDirPath)
	p.multi_threaded_message_loop = gocefToInt(s.MultiThreadedMessageLoop)
	p.external_message_pump = gocefToInt(s.ExternalMessagePump)
	p.windowless_rendering_enabled = gocefToInt(s.WindowlessRenderingEnabled)
	p.command_line_args_disabled = gocefToInt(s.CommandLineArgsDisabled)
	gocefSetUtf16(&p.cache_path, s.CachePath)
	gocefSetUtf16(&p.user_data_path, s.UserDataPath)
	p.persist_session_cookies = gocefToInt(s.PersistSessionCookies)
	p.persist_user_preferences = gocefToInt(s.PersistUserPreferences)
	gocefSetUtf16(&p.user_agent, s.UserAgent)
	gocefSetUtf16(&p.product_version, s.ProductVersion)
	gocefSetUtf16(&p.locale, s.Locale)
	gocefSetUtf16(&p.log_file, s.LogFile)
	p.log_severity = C.cef_log_severity_t(s.LogSeverity)
	gocefSetUtf16(&p.javascript_flags, s.JavascriptFlags)
	gocefSetUtf16(&p.resources_dir_path, s.ResourcesDirPath)
	gocefSetUtf16(&p.locales_dir_path, s.LocalesDirPath)
	p.pack_loading_disabled = gocefToInt(s.PackLoadingDisabled)
	p.remote_debugging_port = C.int(s.RemoteDebuggingPort)
	p.uncaught_exception_stack_size = C.int(s.UncaughtExceptionStackSize)
	p.ignore_certificate_errors = gocefToInt(s.IgnoreCertificateErrors)
	p.enable_net_security_expiration = gocefToInt(s.EnableNetSecurityExpiration)
	p.background_color = C.cef_color_t(gocefToARGB(s.BackgroundColor))
	gocefSetUtf16(&p.accept_language_list, s.AcceptLanguageList)
}

func (s *Settings) toNative() unsafe.Pointer {
	if s.cref == nil {
		s.cref = (*C.cef_settings_t)(C.calloc(1, C.sizeof_cef_settings_t))
		s.cref.size = C.sizeof_cef_settings_t
		s.copyToNative(s.cref)
	}
	return unsafe.Pointer(s.cref)
}

type WindowInfo struct {
	cref                       *C.cef_window_info_t
	WindowName                 string
	X                          int
	Y                          int
	Width                      int
	Height                     int
	ParentWindow               unsafe.Pointer
	WindowlessRenderingEnabled bool
	SharedTextureEnabled       bool
	ExternalBeginFrameEnabled  bool
	Window                     unsafe.Pointer
}

func (w *WindowInfo) copyToNative(p *C.cef_window_info_t) {
	gocefSetUtf16(&p.window_name, w.WindowName)
	p.x = C.uint(w.X)
	p.y = C.uint(w.Y)
	p.width = C.uint(w.Width)
	p.height = C.uint(w.Height)
	p.parent_window = C.ulong(uintptr(w.ParentWindow))
	p.windowless_rendering_enabled = gocefToInt(w.WindowlessRenderingEnabled)
	p.shared_texture_enabled = gocefToInt(w.SharedTextureEnabled)
	p.external_begin_frame_enabled = gocefToInt(w.ExternalBeginFrameEnabled)
	p.window = C.ulong(uintptr(w.Window))
}

func (w *WindowInfo) toNative() unsafe.Pointer {
	if w.cref == nil {
		w.cref = (*C.cef_window_info_t)(C.calloc(1, C.sizeof_cef_window_info_t))
		w.copyToNative(w.cref)
	}
	return unsafe.Pointer(w.cref)
}
