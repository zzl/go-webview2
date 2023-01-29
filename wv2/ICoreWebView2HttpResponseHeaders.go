package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 03C5FF5A-9B45-4A88-881C-89A9F328619C
var IID_ICoreWebView2HttpResponseHeaders = syscall.GUID{0x03C5FF5A, 0x9B45, 0x4A88,
	[8]byte{0x88, 0x1C, 0x89, 0xA9, 0xF3, 0x28, 0x61, 0x9C}}

type ICoreWebView2HttpResponseHeaders struct {
	win32.IUnknown
}

func NewICoreWebView2HttpResponseHeaders(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2HttpResponseHeaders {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2HttpResponseHeaders)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2HttpResponseHeaders) IID() *syscall.GUID {
	return &IID_ICoreWebView2HttpResponseHeaders
}

func (this *ICoreWebView2HttpResponseHeaders) AppendHeader(name string, value string) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpResponseHeaders) Contains(name string, contains *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(contains)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpResponseHeaders) GetHeader(name string, value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpResponseHeaders) GetHeaders(name string, iterator **ICoreWebView2HttpHeadersCollectionIterator) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(iterator)))
	com.AddToScope(iterator)
	return com.Error(ret)
}

func (this *ICoreWebView2HttpResponseHeaders) GetIterator(iterator **ICoreWebView2HttpHeadersCollectionIterator) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(iterator)))
	com.AddToScope(iterator)
	return com.Error(ret)
}
