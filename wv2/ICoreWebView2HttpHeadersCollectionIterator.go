package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 0702FC30-F43B-47BB-AB52-A42CB552AD9F
var IID_ICoreWebView2HttpHeadersCollectionIterator = syscall.GUID{0x0702FC30, 0xF43B, 0x47BB,
	[8]byte{0xAB, 0x52, 0xA4, 0x2C, 0xB5, 0x52, 0xAD, 0x9F}}

type ICoreWebView2HttpHeadersCollectionIterator struct {
	win32.IUnknown
}

func NewICoreWebView2HttpHeadersCollectionIterator(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2HttpHeadersCollectionIterator {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2HttpHeadersCollectionIterator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2HttpHeadersCollectionIterator) IID() *syscall.GUID {
	return &IID_ICoreWebView2HttpHeadersCollectionIterator
}

func (this *ICoreWebView2HttpHeadersCollectionIterator) GetCurrentHeader(name *win32.PWSTR, value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpHeadersCollectionIterator) GetHasCurrentHeader(hasCurrent *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hasCurrent)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpHeadersCollectionIterator) MoveNext(hasNext *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hasNext)))
	return com.Error(ret)
}
