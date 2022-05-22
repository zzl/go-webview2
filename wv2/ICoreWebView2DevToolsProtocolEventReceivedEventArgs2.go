package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 2DC4959D-1494-4393-95BA-BEA4CB9EBD1B
var IID_ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 = syscall.GUID{0x2DC4959D, 0x1494, 0x4393, 
	[8]byte{0x95, 0xBA, 0xBE, 0xA4, 0xCB, 0x9E, 0xBD, 0x1B}}

type ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 struct {
	ICoreWebView2DevToolsProtocolEventReceivedEventArgs
}

func NewICoreWebView2DevToolsProtocolEventReceivedEventArgs2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2 {
	p := (*ICoreWebView2DevToolsProtocolEventReceivedEventArgs2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2) IID() *syscall.GUID {
	return &IID_ICoreWebView2DevToolsProtocolEventReceivedEventArgs2
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventArgs2) GetSessionId(sessionId *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sessionId)))
	return com.Error(ret)
}

