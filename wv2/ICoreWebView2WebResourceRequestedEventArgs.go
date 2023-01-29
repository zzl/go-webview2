package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 453E667F-12C7-49D4-BE6D-DDBE7956F57A
var IID_ICoreWebView2WebResourceRequestedEventArgs = syscall.GUID{0x453E667F, 0x12C7, 0x49D4,
	[8]byte{0xBE, 0x6D, 0xDD, 0xBE, 0x79, 0x56, 0xF5, 0x7A}}

type ICoreWebView2WebResourceRequestedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2WebResourceRequestedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2WebResourceRequestedEventArgs {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2WebResourceRequestedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2WebResourceRequestedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2WebResourceRequestedEventArgs
}

func (this *ICoreWebView2WebResourceRequestedEventArgs) GetRequest(request **ICoreWebView2WebResourceRequest) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)))
	com.AddToScope(request)
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequestedEventArgs) GetResponse(response **ICoreWebView2WebResourceResponse) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(response)))
	com.AddToScope(response)
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequestedEventArgs) SetResponse(response *ICoreWebView2WebResourceResponse) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(response)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequestedEventArgs) GetDeferral(deferral **ICoreWebView2Deferral) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deferral)))
	com.AddToScope(deferral)
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceRequestedEventArgs) GetResourceContext(context *int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(context)))
	return com.Error(ret)
}

