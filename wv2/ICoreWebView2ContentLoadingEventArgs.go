package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 0C8A1275-9B6B-4901-87AD-70DF25BAFA6E
var IID_ICoreWebView2ContentLoadingEventArgs = syscall.GUID{0x0C8A1275, 0x9B6B, 0x4901,
	[8]byte{0x87, 0xAD, 0x70, 0xDF, 0x25, 0xBA, 0xFA, 0x6E}}

type ICoreWebView2ContentLoadingEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2ContentLoadingEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ContentLoadingEventArgs {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2ContentLoadingEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ContentLoadingEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2ContentLoadingEventArgs
}

func (this *ICoreWebView2ContentLoadingEventArgs) GetIsErrorPage(isErrorPage *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isErrorPage)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContentLoadingEventArgs) GetNavigationId(navigationId *uint64) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(navigationId)))
	return com.Error(ret)
}
