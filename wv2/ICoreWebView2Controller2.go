package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// C979903E-D4CA-4228-92EB-47EE3FA96EAB
var IID_ICoreWebView2Controller2 = syscall.GUID{0xC979903E, 0xD4CA, 0x4228, 
	[8]byte{0x92, 0xEB, 0x47, 0xEE, 0x3F, 0xA9, 0x6E, 0xAB}}

type ICoreWebView2Controller2 struct {
	ICoreWebView2Controller
}

func NewICoreWebView2Controller2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Controller2 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2Controller2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Controller2) IID() *syscall.GUID {
	return &IID_ICoreWebView2Controller2
}

func (this *ICoreWebView2Controller2) GetDefaultBackgroundColor(backgroundColor *COREWEBVIEW2_COLOR) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(backgroundColor)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller2) SetDefaultBackgroundColor(backgroundColor COREWEBVIEW2_COLOR) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&backgroundColor)))
	return com.Error(ret)
}

