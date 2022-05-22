package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 4D6E7B5E-9BAA-11EB-A8B3-0242AC130003
var IID_ICoreWebView2FrameCreatedEventArgs = syscall.GUID{0x4D6E7B5E, 0x9BAA, 0x11EB, 
	[8]byte{0xA8, 0xB3, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2FrameCreatedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2FrameCreatedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2FrameCreatedEventArgs {
	p := (*ICoreWebView2FrameCreatedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2FrameCreatedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2FrameCreatedEventArgs
}

func (this *ICoreWebView2FrameCreatedEventArgs) GetFrame(frame **ICoreWebView2Frame) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frame)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*frame).IUnknown))
	}
	return com.Error(ret)
}

