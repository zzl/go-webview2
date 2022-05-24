package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 74D7127F-9DE6-4200-8734-42D6FB4FF741
var IID_ICoreWebView2PermissionRequestedEventArgs2 = syscall.GUID{0x74D7127F, 0x9DE6, 0x4200, 
	[8]byte{0x87, 0x34, 0x42, 0xD6, 0xFB, 0x4F, 0xF7, 0x41}}

type ICoreWebView2PermissionRequestedEventArgs2 struct {
	ICoreWebView2PermissionRequestedEventArgs
}

func NewICoreWebView2PermissionRequestedEventArgs2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2PermissionRequestedEventArgs2 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2PermissionRequestedEventArgs2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2PermissionRequestedEventArgs2) IID() *syscall.GUID {
	return &IID_ICoreWebView2PermissionRequestedEventArgs2
}

func (this *ICoreWebView2PermissionRequestedEventArgs2) GetHandled(handled *int32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handled)))
	return com.Error(ret)
}

func (this *ICoreWebView2PermissionRequestedEventArgs2) SetHandled(handled int32) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(handled))
	return com.Error(ret)
}

