package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 9E8F0CF8-E670-4B5E-B2BC-73E061E3184C
var IID_ICoreWebView2_2 = syscall.GUID{0x9E8F0CF8, 0xE670, 0x4B5E, 
	[8]byte{0xB2, 0xBC, 0x73, 0xE0, 0x61, 0xE3, 0x18, 0x4C}}

type ICoreWebView2_2 struct {
	ICoreWebView2
}

func NewICoreWebView2_2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_2 {
	p := (*ICoreWebView2_2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_2) IID() *syscall.GUID {
	return &IID_ICoreWebView2_2
}

func (this *ICoreWebView2_2) Add_WebResourceResponseReceived(eventHandler *ICoreWebView2WebResourceResponseReceivedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[61]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_2) Remove_WebResourceResponseReceived(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[62]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_2) NavigateWithWebResourceRequest(request *ICoreWebView2WebResourceRequest) com.Error {
	addr := (*this.LpVtbl)[63]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)))
	return com.Error(ret)
}

func (this *ICoreWebView2_2) Add_DOMContentLoaded(eventHandler *ICoreWebView2DOMContentLoadedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[64]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_2) Remove_DOMContentLoaded(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[65]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_2) GetCookieManager(cookieManager **ICoreWebView2CookieManager) com.Error {
	addr := (*this.LpVtbl)[66]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cookieManager)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*cookieManager).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2_2) GetEnvironment(environment **ICoreWebView2Environment) com.Error {
	addr := (*this.LpVtbl)[67]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(environment)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*environment).IUnknown))
	}
	return com.Error(ret)
}

