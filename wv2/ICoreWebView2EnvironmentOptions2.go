package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// FF85C98A-1BA7-4A6B-90C8-2B752C89E9E2
var IID_ICoreWebView2EnvironmentOptions2 = syscall.GUID{0xFF85C98A, 0x1BA7, 0x4A6B,
	[8]byte{0x90, 0xC8, 0x2B, 0x75, 0x2C, 0x89, 0xE9, 0xE2}}

type ICoreWebView2EnvironmentOptions2 struct {
	win32.IUnknown
}

func NewICoreWebView2EnvironmentOptions2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2EnvironmentOptions2 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2EnvironmentOptions2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2EnvironmentOptions2) IID() *syscall.GUID {
	return &IID_ICoreWebView2EnvironmentOptions2
}

func (this *ICoreWebView2EnvironmentOptions2) GetExclusiveUserDataFolderAccess(value *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2EnvironmentOptions2) SetExclusiveUserDataFolderAccess(value int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

