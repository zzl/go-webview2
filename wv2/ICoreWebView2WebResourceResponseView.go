package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 79701053-7759-4162-8F7D-F1B3F084928D
var IID_ICoreWebView2WebResourceResponseView = syscall.GUID{0x79701053, 0x7759, 0x4162,
	[8]byte{0x8F, 0x7D, 0xF1, 0xB3, 0xF0, 0x84, 0x92, 0x8D}}

type ICoreWebView2WebResourceResponseView struct {
	win32.IUnknown
}

func NewICoreWebView2WebResourceResponseView(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2WebResourceResponseView {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2WebResourceResponseView)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2WebResourceResponseView) IID() *syscall.GUID {
	return &IID_ICoreWebView2WebResourceResponseView
}

func (this *ICoreWebView2WebResourceResponseView) GetHeaders(headers **ICoreWebView2HttpResponseHeaders) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(headers)))
	com.AddToScope(headers)
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponseView) GetStatusCode(statusCode *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statusCode)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponseView) GetReasonPhrase(reasonPhrase *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(reasonPhrase)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponseView) GetContent(handler *ICoreWebView2WebResourceResponseViewGetContentCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

