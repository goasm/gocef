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
	return p
}
