package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 97055CD4-512C-4264-8B5F-E3F446CEA6A5
var IID_ICoreWebView2WebResourceRequest = syscall.GUID{0x97055CD4, 0x512C, 0x4264, 
	[8]byte{0x8B, 0x5F, 0xE3, 0xF4, 0x46, 0xCE, 0xA6, 0xA5}}

type ICoreWebView2WebResourceRequest struct {
	win32.IUnknown
}

func NewICoreWebView2WebResourceRequest(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2WebResourceRequest {
	p := (*ICoreWebView2WebResourceRequest)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2WebResourceRequest) IID() *syscall.GUID {
	return &IID_ICoreWebView2WebResourceRequest
}

func (this *ICoreWebView2WebResourceRequest) GetUri(uri *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequest) SetUri(uri string) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequest) GetMethod(method *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(method)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequest) SetMethod(method string) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(method)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequest) GetContent(content **com.UnknownClass) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*content).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequest) SetContent(content *win32.IStream) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequest) GetHeaders(headers **ICoreWebView2HttpRequestHeaders) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(headers)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*headers).IUnknown))
	}
	return com.Error(ret)
}

