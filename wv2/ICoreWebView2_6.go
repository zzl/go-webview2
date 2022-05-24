package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 499AADAC-D92C-4589-8A75-111BFC167795
var IID_ICoreWebView2_6 = syscall.GUID{0x499AADAC, 0xD92C, 0x4589, 
	[8]byte{0x8A, 0x75, 0x11, 0x1B, 0xFC, 0x16, 0x77, 0x95}}

type ICoreWebView2_6 struct {
	ICoreWebView2_5
}

func NewICoreWebView2_6(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_6 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2_6)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_6) IID() *syscall.GUID {
	return &IID_ICoreWebView2_6
}

func (this *ICoreWebView2_6) OpenTaskManagerWindow() com.Error {
	addr := (*this.LpVtbl)[79]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

