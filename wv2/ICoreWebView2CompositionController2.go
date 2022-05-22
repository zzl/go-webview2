package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 0B6A3D24-49CB-4806-BA20-B5E0734A7B26
var IID_ICoreWebView2CompositionController2 = syscall.GUID{0x0B6A3D24, 0x49CB, 0x4806, 
	[8]byte{0xBA, 0x20, 0xB5, 0xE0, 0x73, 0x4A, 0x7B, 0x26}}

type ICoreWebView2CompositionController2 struct {
	ICoreWebView2CompositionController
}

func NewICoreWebView2CompositionController2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2CompositionController2 {
	p := (*ICoreWebView2CompositionController2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2CompositionController2) IID() *syscall.GUID {
	return &IID_ICoreWebView2CompositionController2
}

func (this *ICoreWebView2CompositionController2) GetAutomationProvider(provider **com.UnknownClass) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(provider)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*provider).IUnknown))
	}
	return com.Error(ret)
}

