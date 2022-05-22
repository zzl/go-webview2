package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 183E7052-1D03-43A0-AB99-98E043B66B39
var IID_ICoreWebView2Settings5 = syscall.GUID{0x183E7052, 0x1D03, 0x43A0, 
	[8]byte{0xAB, 0x99, 0x98, 0xE0, 0x43, 0xB6, 0x6B, 0x39}}

type ICoreWebView2Settings5 struct {
	ICoreWebView2Settings4
}

func NewICoreWebView2Settings5(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Settings5 {
	p := (*ICoreWebView2Settings5)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Settings5) IID() *syscall.GUID {
	return &IID_ICoreWebView2Settings5
}

func (this *ICoreWebView2Settings5) GetIsPinchZoomEnabled(enabled *int32) com.Error {
	addr := (*this.LpVtbl)[29]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(enabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings5) SetIsPinchZoomEnabled(enabled int32) com.Error {
	addr := (*this.LpVtbl)[30]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(enabled))
	return com.Error(ret)
}

