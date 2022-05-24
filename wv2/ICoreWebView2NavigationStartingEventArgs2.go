package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 9086BE93-91AA-472D-A7E0-579F2BA006AD
var IID_ICoreWebView2NavigationStartingEventArgs2 = syscall.GUID{0x9086BE93, 0x91AA, 0x472D, 
	[8]byte{0xA7, 0xE0, 0x57, 0x9F, 0x2B, 0xA0, 0x06, 0xAD}}

type ICoreWebView2NavigationStartingEventArgs2 struct {
	ICoreWebView2NavigationStartingEventArgs
}

func NewICoreWebView2NavigationStartingEventArgs2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2NavigationStartingEventArgs2 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2NavigationStartingEventArgs2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2NavigationStartingEventArgs2) IID() *syscall.GUID {
	return &IID_ICoreWebView2NavigationStartingEventArgs2
}

func (this *ICoreWebView2NavigationStartingEventArgs2) GetAdditionalAllowedFrameAncestors(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationStartingEventArgs2) SetAdditionalAllowedFrameAncestors(value string) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

