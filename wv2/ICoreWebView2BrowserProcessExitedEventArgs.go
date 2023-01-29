package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 1F00663F-AF8C-4782-9CDD-DD01C52E34CB
var IID_ICoreWebView2BrowserProcessExitedEventArgs = syscall.GUID{0x1F00663F, 0xAF8C, 0x4782,
	[8]byte{0x9C, 0xDD, 0xDD, 0x01, 0xC5, 0x2E, 0x34, 0xCB}}

type ICoreWebView2BrowserProcessExitedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2BrowserProcessExitedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2BrowserProcessExitedEventArgs {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2BrowserProcessExitedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2BrowserProcessExitedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2BrowserProcessExitedEventArgs
}

func (this *ICoreWebView2BrowserProcessExitedEventArgs) GetBrowserProcessExitKind(browserProcessExitKind *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(browserProcessExitKind)))
	return com.Error(ret)
}

func (this *ICoreWebView2BrowserProcessExitedEventArgs) GetBrowserProcessId(value *uint32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}
