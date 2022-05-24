package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// B1690564-6F5A-4983-8E48-31D1143FECDB
var IID_ICoreWebView2_10 = syscall.GUID{0xB1690564, 0x6F5A, 0x4983, 
	[8]byte{0x8E, 0x48, 0x31, 0xD1, 0x14, 0x3F, 0xEC, 0xDB}}

type ICoreWebView2_10 struct {
	ICoreWebView2_9
}

func NewICoreWebView2_10(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_10 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2_10)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_10) IID() *syscall.GUID {
	return &IID_ICoreWebView2_10
}

func (this *ICoreWebView2_10) Add_BasicAuthenticationRequested(eventHandler *ICoreWebView2BasicAuthenticationRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[97]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_10) Remove_BasicAuthenticationRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[98]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

