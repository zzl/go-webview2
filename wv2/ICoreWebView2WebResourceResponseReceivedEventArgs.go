package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// D1DB483D-6796-4B8B-80FC-13712BB716F4
var IID_ICoreWebView2WebResourceResponseReceivedEventArgs = syscall.GUID{0xD1DB483D, 0x6796, 0x4B8B, 
	[8]byte{0x80, 0xFC, 0x13, 0x71, 0x2B, 0xB7, 0x16, 0xF4}}

type ICoreWebView2WebResourceResponseReceivedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2WebResourceResponseReceivedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2WebResourceResponseReceivedEventArgs {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2WebResourceResponseReceivedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2WebResourceResponseReceivedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2WebResourceResponseReceivedEventArgs
}

func (this *ICoreWebView2WebResourceResponseReceivedEventArgs) GetRequest(request **ICoreWebView2WebResourceRequest) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(request)))
		com.AddToScope(request)
	return com.Error(ret)
}

func (this *ICoreWebView2WebResourceResponseReceivedEventArgs) GetResponse(response **ICoreWebView2WebResourceResponseView) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(response)))
		com.AddToScope(response)
	return com.Error(ret)
}

