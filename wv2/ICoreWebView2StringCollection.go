package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// F41F3F8A-BCC3-11EB-8529-0242AC130003
var IID_ICoreWebView2StringCollection = syscall.GUID{0xF41F3F8A, 0xBCC3, 0x11EB, 
	[8]byte{0x85, 0x29, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2StringCollection struct {
	win32.IUnknown
}

func NewICoreWebView2StringCollection(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2StringCollection {
	p := (*ICoreWebView2StringCollection)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2StringCollection) IID() *syscall.GUID {
	return &IID_ICoreWebView2StringCollection
}

func (this *ICoreWebView2StringCollection) GetCount(value *uint32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2StringCollection) GetValueAtIndex(index uint32, value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

