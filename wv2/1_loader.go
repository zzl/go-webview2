package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"syscall"
	"unsafe"
)

var (
	loaderDll                                    = syscall.NewLazyDLL("WebView2Loader.dll")
	createCoreWebView2EnvironmentWithOptions     = loaderDll.NewProc("CreateCoreWebView2EnvironmentWithOptions")
	createCoreWebView2Environment                = loaderDll.NewProc("CreateCoreWebView2Environment")
	getAvailableCoreWebView2BrowserVersionString = loaderDll.NewProc("GetAvailableCoreWebView2BrowserVersionString")
	compareBrowserVersions                       = loaderDll.NewProc("CompareBrowserVersions")
)

func CreateCoreWebView2EnvironmentWithOptions(
	browserExecutableFolder string,
	userDataFolder string,
	environmentOptions *ICoreWebView2EnvironmentOptions,
	environmentCreatedHandler *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler) win32.HRESULT {
	ret, _, _ := syscall.SyscallN(createCoreWebView2EnvironmentWithOptions.Addr(),
		uintptr(win32.StrToPointerOrNil(browserExecutableFolder)),
		uintptr(win32.StrToPointerOrNil(userDataFolder)),
		uintptr(unsafe.Pointer(environmentOptions)),
		uintptr(unsafe.Pointer(environmentCreatedHandler)),
	)
	return win32.HRESULT(ret)
}

func CreateCoreWebView2Environment(
	environmentCreatedHandler *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler) win32.HRESULT {
	ret, _, _ := syscall.SyscallN(createCoreWebView2Environment.Addr(),
		uintptr(unsafe.Pointer(environmentCreatedHandler)),
	)
	return win32.HRESULT(ret)
}

func GetAvailableCoreWebView2BrowserVersionString(
	browserExecutableFolder string,
	versionInfo *string) win32.HRESULT {
	var pwszVersionInfo win32.PWSTR
	ret, _, _ := syscall.SyscallN(getAvailableCoreWebView2BrowserVersionString.Addr(),
		uintptr(win32.StrToPointer(browserExecutableFolder)),
		uintptr(unsafe.Pointer(&pwszVersionInfo)),
	)
	if pwszVersionInfo == nil {
		*versionInfo = ""
	} else {
		*versionInfo = win32.PwstrToStr(pwszVersionInfo)
		win32.CoTaskMemFree(unsafe.Pointer(pwszVersionInfo))
	}
	return win32.HRESULT(ret)
}

func CompareBrowserVersions(
	version1 string, version2 string, result *int32) win32.HRESULT {
	ret, _, _ := syscall.SyscallN(compareBrowserVersions.Addr(),
		uintptr(win32.StrToPointer(version1)),
		uintptr(win32.StrToPointer(version2)),
		uintptr(unsafe.Pointer(result)),
	)
	return win32.HRESULT(ret)
}
