package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// AAFCC94F-FA27-48FD-97DF-830EF75AAEC9
var IID_ICoreWebView2WebResourceResponse = syscall.GUID{0xAAFCC94F, 0xFA27, 0x48FD, 
	[8]byte{0x97, 0xDF, 0x83, 0x0E, 0xF7, 0x5A, 0xAE, 0xC9}}

type ICoreWebView2WebResourceResponse struct {
	win32.IUnknown
}

func NewICoreWebView2WebResourceResponse(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2WebResourceResponse {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2WebResourceResponse)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2WebResourceResponse) IID() *syscall.GUID {
	return &IID_ICoreWebView2WebResourceResponse
}

func (this *ICoreWebView2WebResourceResponse) GetContent(content **win32.IUnknown) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)))
		com.AddToScope(content)
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponse) SetContent(content *win32.IStream) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponse) GetHeaders(headers **ICoreWebView2HttpResponseHeaders) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(headers)))
		com.AddToScope(headers)
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponse) GetStatusCode(statusCode *int32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(statusCode)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponse) SetStatusCode(statusCode int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(statusCode))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponse) GetReasonPhrase(reasonPhrase *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(reasonPhrase)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponse) SetReasonPhrase(reasonPhrase string) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(reasonPhrase)))
	return com.Error(ret)
}

