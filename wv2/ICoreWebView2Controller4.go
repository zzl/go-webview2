package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 97D418D5-A426-4E49-A151-E1A10F327D9E
var IID_ICoreWebView2Controller4 = syscall.GUID{0x97D418D5, 0xA426, 0x4E49,
	[8]byte{0xA1, 0x51, 0xE1, 0xA1, 0x0F, 0x32, 0x7D, 0x9E}}

type ICoreWebView2Controller4 struct {
	ICoreWebView2Controller3
}

func NewICoreWebView2Controller4(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Controller4 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Controller4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Controller4) IID() *syscall.GUID {
	return &IID_ICoreWebView2Controller4
}

func (this *ICoreWebView2Controller4) GetAllowExternalDrop(value *int32) com.Error {
	addr := (*this.LpVtbl)[36]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller4) SetAllowExternalDrop(value int32) com.Error {
	addr := (*this.LpVtbl)[37]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}
