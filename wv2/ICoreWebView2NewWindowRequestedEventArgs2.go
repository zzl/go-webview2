package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// BBC7BAED-74C6-4C92-B63A-7F5AEAE03DE3
var IID_ICoreWebView2NewWindowRequestedEventArgs2 = syscall.GUID{0xBBC7BAED, 0x74C6, 0x4C92, 
	[8]byte{0xB6, 0x3A, 0x7F, 0x5A, 0xEA, 0xE0, 0x3D, 0xE3}}

type ICoreWebView2NewWindowRequestedEventArgs2 struct {
	ICoreWebView2NewWindowRequestedEventArgs
}

func NewICoreWebView2NewWindowRequestedEventArgs2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2NewWindowRequestedEventArgs2 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2NewWindowRequestedEventArgs2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2NewWindowRequestedEventArgs2) IID() *syscall.GUID {
	return &IID_ICoreWebView2NewWindowRequestedEventArgs2
}

func (this *ICoreWebView2NewWindowRequestedEventArgs2) GetName(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

