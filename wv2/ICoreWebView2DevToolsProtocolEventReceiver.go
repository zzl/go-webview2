package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// B32CA51A-8371-45E9-9317-AF021D080367
var IID_ICoreWebView2DevToolsProtocolEventReceiver = syscall.GUID{0xB32CA51A, 0x8371, 0x45E9, 
	[8]byte{0x93, 0x17, 0xAF, 0x02, 0x1D, 0x08, 0x03, 0x67}}

type ICoreWebView2DevToolsProtocolEventReceiver struct {
	win32.IUnknown
}

func NewICoreWebView2DevToolsProtocolEventReceiver(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2DevToolsProtocolEventReceiver {
	p := (*ICoreWebView2DevToolsProtocolEventReceiver)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2DevToolsProtocolEventReceiver) IID() *syscall.GUID {
	return &IID_ICoreWebView2DevToolsProtocolEventReceiver
}

func (this *ICoreWebView2DevToolsProtocolEventReceiver) Add_DevToolsProtocolEventReceived(handler *ICoreWebView2DevToolsProtocolEventReceivedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2DevToolsProtocolEventReceiver) Remove_DevToolsProtocolEventReceived(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

