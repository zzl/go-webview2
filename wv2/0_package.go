package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"os"
)

func EnsureRuntimeAvailable() {
	var ver string
	GetAvailableCoreWebView2BrowserVersionString("", &ver)
	if ver == "" {
		msg := win32.StrToPwstr("Webview2 runtime not found. \n" +
			"Please install it before using this application.")
		win32.MessageBox(0, msg, win32.StrToPwstr("go-webview2"), win32.MB_ICONEXCLAMATION)
		url := win32.StrToPwstr("https://developer.microsoft.com/en-us/microsoft-edge/webview2/#download-section")
		win32.ShellExecute(0, win32.StrToPwstr("Open"), url, nil, nil, int32(win32.SW_SHOWNORMAL))
		os.Exit(1)
	}
}
