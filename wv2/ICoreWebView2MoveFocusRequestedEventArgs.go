package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 2D6AA13B-3839-4A15-92FC-D88B3C0D9C9D
var IID_ICoreWebView2MoveFocusRequestedEventArgs = syscall.GUID{0x2D6AA13B, 0x3839, 0x4A15,
	[8]byte{0x92, 0xFC, 0xD8, 0x8B, 0x3C, 0x0D, 0x9C, 0x9D}}

type ICoreWebView2MoveFocusRequestedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2MoveFocusRequestedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2MoveFocusRequestedEventArgs {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2MoveFocusRequestedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2MoveFocusRequestedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2MoveFocusRequestedEventArgs
}

func (this *ICoreWebView2MoveFocusRequestedEventArgs) GetReason(reason *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(reason)))
	return com.Error(ret)
}

func (this *ICoreWebView2MoveFocusRequestedEventArgs) GetHandled(value *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2MoveFocusRequestedEventArgs) SetHandled(value int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

