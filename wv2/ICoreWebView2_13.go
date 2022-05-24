package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// F75F09A8-667E-4983-88D6-C8773F315E84
var IID_ICoreWebView2_13 = syscall.GUID{0xF75F09A8, 0x667E, 0x4983, 
	[8]byte{0x88, 0xD6, 0xC8, 0x77, 0x3F, 0x31, 0x5E, 0x84}}

type ICoreWebView2_13 struct {
	ICoreWebView2_12
}

func NewICoreWebView2_13(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_13 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2_13)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_13) IID() *syscall.GUID {
	return &IID_ICoreWebView2_13
}

func (this *ICoreWebView2_13) GetProfile(value **ICoreWebView2Profile) com.Error {
	addr := (*this.LpVtbl)[105]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
		com.AddToScope(value)
	return com.Error(ret)
}

