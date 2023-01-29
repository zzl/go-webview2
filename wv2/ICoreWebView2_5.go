package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// BEDB11B8-D63C-11EB-B8BC-0242AC130003
var IID_ICoreWebView2_5 = syscall.GUID{0xBEDB11B8, 0xD63C, 0x11EB,
	[8]byte{0xB8, 0xBC, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2_5 struct {
	ICoreWebView2_4
}

func NewICoreWebView2_5(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_5 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2_5)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_5) IID() *syscall.GUID {
	return &IID_ICoreWebView2_5
}

func (this *ICoreWebView2_5) Add_ClientCertificateRequested(eventHandler *ICoreWebView2ClientCertificateRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[77]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_5) Remove_ClientCertificateRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[78]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

