package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 8155A9A4-1474-4A86-8CAE-151B0FA6B8CA
var IID_ICoreWebView2ProcessFailedEventArgs = syscall.GUID{0x8155A9A4, 0x1474, 0x4A86,
	[8]byte{0x8C, 0xAE, 0x15, 0x1B, 0x0F, 0xA6, 0xB8, 0xCA}}

type ICoreWebView2ProcessFailedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2ProcessFailedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ProcessFailedEventArgs {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2ProcessFailedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ProcessFailedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2ProcessFailedEventArgs
}

func (this *ICoreWebView2ProcessFailedEventArgs) GetProcessFailedKind(processFailedKind *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(processFailedKind)))
	return com.Error(ret)
}

