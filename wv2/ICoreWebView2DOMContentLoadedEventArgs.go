package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 16B1E21A-C503-44F2-84C9-70ABA5031283
var IID_ICoreWebView2DOMContentLoadedEventArgs = syscall.GUID{0x16B1E21A, 0xC503, 0x44F2, 
	[8]byte{0x84, 0xC9, 0x70, 0xAB, 0xA5, 0x03, 0x12, 0x83}}

type ICoreWebView2DOMContentLoadedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2DOMContentLoadedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2DOMContentLoadedEventArgs {
	p := (*ICoreWebView2DOMContentLoadedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2DOMContentLoadedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2DOMContentLoadedEventArgs
}

func (this *ICoreWebView2DOMContentLoadedEventArgs) GetNavigationId(navigationId *uint64) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(navigationId)))
	return com.Error(ret)
}

