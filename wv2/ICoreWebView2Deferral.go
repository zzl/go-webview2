package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// C10E7F7B-B585-46F0-A623-8BEFBF3E4EE0
var IID_ICoreWebView2Deferral = syscall.GUID{0xC10E7F7B, 0xB585, 0x46F0, 
	[8]byte{0xA6, 0x23, 0x8B, 0xEF, 0xBF, 0x3E, 0x4E, 0xE0}}

type ICoreWebView2Deferral struct {
	win32.IUnknown
}

func NewICoreWebView2Deferral(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Deferral {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2Deferral)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Deferral) IID() *syscall.GUID {
	return &IID_ICoreWebView2Deferral
}

func (this *ICoreWebView2Deferral) Complete() com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

