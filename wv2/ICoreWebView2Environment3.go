package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 80A22AE3-BE7C-4CE2-AFE1-5A50056CDEEB
var IID_ICoreWebView2Environment3 = syscall.GUID{0x80A22AE3, 0xBE7C, 0x4CE2,
	[8]byte{0xAF, 0xE1, 0x5A, 0x50, 0x05, 0x6C, 0xDE, 0xEB}}

type ICoreWebView2Environment3 struct {
	ICoreWebView2Environment2
}

func NewICoreWebView2Environment3(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment3 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Environment3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment3) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment3
}

func (this *ICoreWebView2Environment3) CreateCoreWebView2CompositionController(parentWindow win32.HWND, handler *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(parentWindow), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2Environment3) CreateCoreWebView2PointerInfo(pointerInfo **ICoreWebView2PointerInfo) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pointerInfo)))
	com.AddToScope(pointerInfo)
	return com.Error(ret)
}
