package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 5B495469-E119-438A-9B18-7604F25F2E49
var IID_ICoreWebView2NavigationStartingEventArgs = syscall.GUID{0x5B495469, 0xE119, 0x438A,
	[8]byte{0x9B, 0x18, 0x76, 0x04, 0xF2, 0x5F, 0x2E, 0x49}}

type ICoreWebView2NavigationStartingEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2NavigationStartingEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2NavigationStartingEventArgs {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2NavigationStartingEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2NavigationStartingEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2NavigationStartingEventArgs
}

func (this *ICoreWebView2NavigationStartingEventArgs) GetUri(uri *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationStartingEventArgs) GetIsUserInitiated(isUserInitiated *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isUserInitiated)))
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationStartingEventArgs) GetIsRedirected(isRedirected *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isRedirected)))
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationStartingEventArgs) GetRequestHeaders(requestHeaders **ICoreWebView2HttpRequestHeaders) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(requestHeaders)))
	com.AddToScope(requestHeaders)
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationStartingEventArgs) GetCancel(cancel *int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cancel)))
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationStartingEventArgs) SetCancel(cancel int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(cancel))
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationStartingEventArgs) GetNavigationId(navigationId *uint64) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(navigationId)))
	return com.Error(ret)
}

