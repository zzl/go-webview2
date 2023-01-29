package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// D6EB91DD-C3D2-45E5-BD29-6DC2BC4DE9CF
var IID_ICoreWebView2Environment8 = syscall.GUID{0xD6EB91DD, 0xC3D2, 0x45E5,
	[8]byte{0xBD, 0x29, 0x6D, 0xC2, 0xBC, 0x4D, 0xE9, 0xCF}}

type ICoreWebView2Environment8 struct {
	ICoreWebView2Environment7
}

func NewICoreWebView2Environment8(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment8 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Environment8)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment8) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment8
}

func (this *ICoreWebView2Environment8) Add_ProcessInfosChanged(eventHandler *ICoreWebView2ProcessInfosChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Environment8) Remove_ProcessInfosChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Environment8) GetProcessInfos(value **ICoreWebView2ProcessInfoCollection) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	com.AddToScope(value)
	return com.Error(ret)
}
