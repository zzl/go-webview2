package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 319E423D-E0D7-4B8D-9254-AE9475DE9B17
var IID_ICoreWebView2Environment5 = syscall.GUID{0x319E423D, 0xE0D7, 0x4B8D, 
	[8]byte{0x92, 0x54, 0xAE, 0x94, 0x75, 0xDE, 0x9B, 0x17}}

type ICoreWebView2Environment5 struct {
	ICoreWebView2Environment4
}

func NewICoreWebView2Environment5(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment5 {
	p := (*ICoreWebView2Environment5)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment5) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment5
}

func (this *ICoreWebView2Environment5) Add_BrowserProcessExited(eventHandler *ICoreWebView2BrowserProcessExitedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Environment5) Remove_BrowserProcessExited(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

