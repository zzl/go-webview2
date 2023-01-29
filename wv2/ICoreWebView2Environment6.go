package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E59EE362-ACBD-4857-9A8E-D3644D9459A9
var IID_ICoreWebView2Environment6 = syscall.GUID{0xE59EE362, 0xACBD, 0x4857,
	[8]byte{0x9A, 0x8E, 0xD3, 0x64, 0x4D, 0x94, 0x59, 0xA9}}

type ICoreWebView2Environment6 struct {
	ICoreWebView2Environment5
}

func NewICoreWebView2Environment6(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment6 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Environment6)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment6) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment6
}

func (this *ICoreWebView2Environment6) CreatePrintSettings(printSettings **ICoreWebView2PrintSettings) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(printSettings)))
	com.AddToScope(printSettings)
	return com.Error(ret)
}

