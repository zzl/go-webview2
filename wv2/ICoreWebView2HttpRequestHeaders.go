package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E86CAC0E-5523-465C-B536-8FB9FC8C8C60
var IID_ICoreWebView2HttpRequestHeaders = syscall.GUID{0xE86CAC0E, 0x5523, 0x465C, 
	[8]byte{0xB5, 0x36, 0x8F, 0xB9, 0xFC, 0x8C, 0x8C, 0x60}}

type ICoreWebView2HttpRequestHeaders struct {
	win32.IUnknown
}

func NewICoreWebView2HttpRequestHeaders(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2HttpRequestHeaders {
	p := (*ICoreWebView2HttpRequestHeaders)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2HttpRequestHeaders) IID() *syscall.GUID {
	return &IID_ICoreWebView2HttpRequestHeaders
}

func (this *ICoreWebView2HttpRequestHeaders) GetHeader(name string, value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpRequestHeaders) GetHeaders(name string, iterator **ICoreWebView2HttpHeadersCollectionIterator) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(iterator)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*iterator).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2HttpRequestHeaders) Contains(name string, contains *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(contains)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpRequestHeaders) SetHeader(name string, value string) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpRequestHeaders) RemoveHeader(name string) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)))
	return com.Error(ret)
}

func (this *ICoreWebView2HttpRequestHeaders) GetIterator(iterator **ICoreWebView2HttpHeadersCollectionIterator) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(iterator)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*iterator).IUnknown))
	}
	return com.Error(ret)
}

