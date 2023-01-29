package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 20944379-6DCF-41D6-A0A0-ABC0FC50DE0D
var IID_ICoreWebView2Environment4 = syscall.GUID{0x20944379, 0x6DCF, 0x41D6,
	[8]byte{0xA0, 0xA0, 0xAB, 0xC0, 0xFC, 0x50, 0xDE, 0x0D}}

type ICoreWebView2Environment4 struct {
	ICoreWebView2Environment3
}

func NewICoreWebView2Environment4(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment4 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Environment4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment4) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment4
}

func (this *ICoreWebView2Environment4) GetAutomationProviderForWindow(hwnd win32.HWND, provider **win32.IUnknown) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(hwnd), uintptr(unsafe.Pointer(provider)))
	com.AddToScope(provider)
	return com.Error(ret)
}

