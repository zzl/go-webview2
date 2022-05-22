package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// A1D309EE-C03F-11EB-8529-0242AC130003
var IID_ICoreWebView2ContextMenuRequestedEventArgs = syscall.GUID{0xA1D309EE, 0xC03F, 0x11EB, 
	[8]byte{0x85, 0x29, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2ContextMenuRequestedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2ContextMenuRequestedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ContextMenuRequestedEventArgs {
	p := (*ICoreWebView2ContextMenuRequestedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2ContextMenuRequestedEventArgs
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) GetMenuItems(value **ICoreWebView2ContextMenuItemCollection) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*value).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) GetContextMenuTarget(value **ICoreWebView2ContextMenuTarget) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*value).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) GetLocation(value *TagPOINT) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) SetSelectedCommandId(value int32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) GetSelectedCommandId(value *int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) SetHandled(value int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) GetHandled(value *int32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuRequestedEventArgs) GetDeferral(deferral **ICoreWebView2Deferral) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deferral)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*deferral).IUnknown))
	}
	return com.Error(ret)
}

