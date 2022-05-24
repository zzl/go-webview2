package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// EE0EB9DF-6F12-46CE-B53F-3F47B9C928E0
var IID_ICoreWebView2Environment10 = syscall.GUID{0xEE0EB9DF, 0x6F12, 0x46CE, 
	[8]byte{0xB5, 0x3F, 0x3F, 0x47, 0xB9, 0xC9, 0x28, 0xE0}}

type ICoreWebView2Environment10 struct {
	ICoreWebView2Environment9
}

func NewICoreWebView2Environment10(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment10 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2Environment10)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment10) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment10
}

func (this *ICoreWebView2Environment10) CreateCoreWebView2ControllerOptions(options **ICoreWebView2ControllerOptions) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(options)))
		com.AddToScope(options)
	return com.Error(ret)
}

func (this *ICoreWebView2Environment10) CreateCoreWebView2ControllerWithOptions(parentWindow win32.HWND, options *ICoreWebView2ControllerOptions, handler *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(parentWindow), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2Environment10) CreateCoreWebView2CompositionControllerWithOptions(parentWindow win32.HWND, options *ICoreWebView2ControllerOptions, handler *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(parentWindow), uintptr(unsafe.Pointer(options)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

