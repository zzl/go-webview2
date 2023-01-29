package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 84FA7612-3F3D-4FBF-889D-FAD000492D72
var IID_ICoreWebView2ProcessInfo = syscall.GUID{0x84FA7612, 0x3F3D, 0x4FBF,
	[8]byte{0x88, 0x9D, 0xFA, 0xD0, 0x00, 0x49, 0x2D, 0x72}}

type ICoreWebView2ProcessInfo struct {
	win32.IUnknown
}

func NewICoreWebView2ProcessInfo(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ProcessInfo {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2ProcessInfo)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ProcessInfo) IID() *syscall.GUID {
	return &IID_ICoreWebView2ProcessInfo
}

func (this *ICoreWebView2ProcessInfo) GetProcessId(value *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ProcessInfo) GetKind(kind *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(kind)))
	return com.Error(ret)
}
