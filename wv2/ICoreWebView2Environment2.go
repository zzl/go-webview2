package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 41F3632B-5EF4-404F-AD82-2D606C5A9A21
var IID_ICoreWebView2Environment2 = syscall.GUID{0x41F3632B, 0x5EF4, 0x404F,
	[8]byte{0xAD, 0x82, 0x2D, 0x60, 0x6C, 0x5A, 0x9A, 0x21}}

type ICoreWebView2Environment2 struct {
	ICoreWebView2Environment
}

func NewICoreWebView2Environment2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment2 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Environment2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment2) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment2
}

func (this *ICoreWebView2Environment2) CreateWebResourceRequest(uri string, method string, postData *win32.IStream, headers string, request **ICoreWebView2WebResourceRequest) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(uri)), uintptr(win32.StrToPointer(method)), uintptr(unsafe.Pointer(postData)), uintptr(win32.StrToPointer(headers)), uintptr(unsafe.Pointer(request)))
	com.AddToScope(request)
	return com.Error(ret)
}
