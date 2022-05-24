package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 0BE78E56-C193-4051-B943-23B460C08BDB
var IID_ICoreWebView2_11 = syscall.GUID{0x0BE78E56, 0xC193, 0x4051, 
	[8]byte{0xB9, 0x43, 0x23, 0xB4, 0x60, 0xC0, 0x8B, 0xDB}}

type ICoreWebView2_11 struct {
	ICoreWebView2_10
}

func NewICoreWebView2_11(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_11 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2_11)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_11) IID() *syscall.GUID {
	return &IID_ICoreWebView2_11
}

func (this *ICoreWebView2_11) CallDevToolsProtocolMethodForSession(sessionId string, methodName string, parametersAsJson string, handler *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[99]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(sessionId)), uintptr(win32.StrToPointer(methodName)), uintptr(win32.StrToPointer(parametersAsJson)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2_11) Add_ContextMenuRequested(eventHandler *ICoreWebView2ContextMenuRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[100]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_11) Remove_ContextMenuRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[101]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

