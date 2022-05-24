package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 35D69927-BCFA-4566-9349-6B3E0D154CAC
var IID_ICoreWebView2_12 = syscall.GUID{0x35D69927, 0xBCFA, 0x4566, 
	[8]byte{0x93, 0x49, 0x6B, 0x3E, 0x0D, 0x15, 0x4C, 0xAC}}

type ICoreWebView2_12 struct {
	ICoreWebView2_11
}

func NewICoreWebView2_12(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_12 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2_12)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_12) IID() *syscall.GUID {
	return &IID_ICoreWebView2_12
}

func (this *ICoreWebView2_12) Add_StatusBarTextChanged(eventHandler *ICoreWebView2StatusBarTextChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[102]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_12) Remove_StatusBarTextChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[103]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_12) GetStatusBarText(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[104]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

