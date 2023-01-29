package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// DA86B8A1-BDF3-4F11-9955-528CEFA59727
var IID_ICoreWebView2FrameInfo = syscall.GUID{0xDA86B8A1, 0xBDF3, 0x4F11,
	[8]byte{0x99, 0x55, 0x52, 0x8C, 0xEF, 0xA5, 0x97, 0x27}}

type ICoreWebView2FrameInfo struct {
	win32.IUnknown
}

func NewICoreWebView2FrameInfo(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2FrameInfo {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2FrameInfo)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2FrameInfo) IID() *syscall.GUID {
	return &IID_ICoreWebView2FrameInfo
}

func (this *ICoreWebView2FrameInfo) GetName(name *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)))
	return com.Error(ret)
}

func (this *ICoreWebView2FrameInfo) GetSource(source *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)))
	return com.Error(ret)
}
