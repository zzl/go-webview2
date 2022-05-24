package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 11CB3ACD-9BC8-43B8-83BF-F40753714F87
var IID_ICoreWebView2Settings6 = syscall.GUID{0x11CB3ACD, 0x9BC8, 0x43B8, 
	[8]byte{0x83, 0xBF, 0xF4, 0x07, 0x53, 0x71, 0x4F, 0x87}}

type ICoreWebView2Settings6 struct {
	ICoreWebView2Settings5
}

func NewICoreWebView2Settings6(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Settings6 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2Settings6)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Settings6) IID() *syscall.GUID {
	return &IID_ICoreWebView2Settings6
}

func (this *ICoreWebView2Settings6) GetIsSwipeNavigationEnabled(enabled *int32) com.Error {
	addr := (*this.LpVtbl)[31]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(enabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings6) SetIsSwipeNavigationEnabled(enabled int32) com.Error {
	addr := (*this.LpVtbl)[32]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(enabled))
	return com.Error(ret)
}

