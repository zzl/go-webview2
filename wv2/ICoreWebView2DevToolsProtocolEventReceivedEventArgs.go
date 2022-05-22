package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 653C2959-BB3A-4377-8632-B58ADA4E66C4
var IID_ICoreWebView2DevToolsProtocolEventReceivedEventArgs = syscall.GUID{0x653C2959, 0xBB3A, 0x4377, 
	[8]byte{0x86, 0x32, 0xB5, 0x8A, 0xDA, 0x4E, 0x66, 0xC4}}

type ICoreWebView2DevToolsProtocolEventReceivedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2DevToolsProtocolEventReceivedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2DevToolsProtocolEventReceivedEventArgs {
	p := (*ICoreWebView2DevToolsProtocolEventReceivedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2DevToolsProtocolEventReceivedEventArgs
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) GetParameterObjectAsJson(parameterObjectAsJson *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(parameterObjectAsJson)))
	return com.Error(ret)
}

