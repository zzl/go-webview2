package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 30D68B7D-20D9-4752-A9CA-EC8448FBB5C1
var IID_ICoreWebView2NavigationCompletedEventArgs = syscall.GUID{0x30D68B7D, 0x20D9, 0x4752, 
	[8]byte{0xA9, 0xCA, 0xEC, 0x84, 0x48, 0xFB, 0xB5, 0xC1}}

type ICoreWebView2NavigationCompletedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2NavigationCompletedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2NavigationCompletedEventArgs {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2NavigationCompletedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2NavigationCompletedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2NavigationCompletedEventArgs
}

func (this *ICoreWebView2NavigationCompletedEventArgs) GetIsSuccess(isSuccess *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isSuccess)))
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationCompletedEventArgs) GetWebErrorStatus(webErrorStatus *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webErrorStatus)))
	return com.Error(ret)
}

func (this *ICoreWebView2NavigationCompletedEventArgs) GetNavigationId(navigationId *uint64) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(navigationId)))
	return com.Error(ret)
}

