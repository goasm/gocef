package gocef

/*
#include <stdlib.h>
#include "include/internal/cef_types.h"
*/
import "C"
import "image/color"

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

type Settings struct {
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

func (s *Settings) toNative() *C.cef_settings_t {
	p := (*C.cef_settings_t)(C.calloc(1, C.sizeof_cef_settings_t))
	p.size = C.sizeof_cef_settings_t
	p.no_sandbox = gocefFromBool(s.NoSandbox)
	p.browser_subprocess_path = *gocefToUtf16(s.BrowserSubprocessPath)
	p.framework_dir_path = *gocefToUtf16(s.FrameworkDirPath)
	p.multi_threaded_message_loop = gocefFromBool(s.MultiThreadedMessageLoop)
	p.external_message_pump = gocefFromBool(s.ExternalMessagePump)
	p.windowless_rendering_enabled = gocefFromBool(s.WindowlessRenderingEnabled)
	p.command_line_args_disabled = gocefFromBool(s.CommandLineArgsDisabled)
	p.cache_path = *gocefToUtf16(s.CachePath)
	p.user_data_path = *gocefToUtf16(s.UserDataPath)
	p.persist_session_cookies = gocefFromBool(s.PersistSessionCookies)
	p.persist_user_preferences = gocefFromBool(s.PersistUserPreferences)
	p.user_agent = *gocefToUtf16(s.UserAgent)
	p.product_version = *gocefToUtf16(s.ProductVersion)
	p.locale = *gocefToUtf16(s.Locale)
	p.log_file = *gocefToUtf16(s.LogFile)
	p.log_severity = C.cef_log_severity_t(s.LogSeverity)
	p.javascript_flags = *gocefToUtf16(s.JavascriptFlags)
	p.resources_dir_path = *gocefToUtf16(s.ResourcesDirPath)
	p.locales_dir_path = *gocefToUtf16(s.LocalesDirPath)
	p.pack_loading_disabled = gocefFromBool(s.PackLoadingDisabled)
	p.remote_debugging_port = C.int(s.RemoteDebuggingPort)
	p.uncaught_exception_stack_size = C.int(s.UncaughtExceptionStackSize)
	p.ignore_certificate_errors = gocefFromBool(s.IgnoreCertificateErrors)
	p.enable_net_security_expiration = gocefFromBool(s.EnableNetSecurityExpiration)
	p.background_color = C.cef_color_t(gocefToARGB(s.BackgroundColor))
	p.accept_language_list = *gocefToUtf16(s.AcceptLanguageList)
	return p
}
