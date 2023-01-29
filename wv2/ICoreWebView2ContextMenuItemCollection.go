package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F562A2F5-C415-45CF-B909-D4B7C1E276D3
var IID_ICoreWebView2ContextMenuItemCollection = syscall.GUID{0xF562A2F5, 0xC415, 0x45CF,
	[8]byte{0xB9, 0x09, 0xD4, 0xB7, 0xC1, 0xE2, 0x76, 0xD3}}

type ICoreWebView2ContextMenuItemCollection struct {
	win32.IUnknown
}

func NewICoreWebView2ContextMenuItemCollection(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ContextMenuItemCollection {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2ContextMenuItemCollection)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ContextMenuItemCollection) IID() *syscall.GUID {
	return &IID_ICoreWebView2ContextMenuItemCollection
}

func (this *ICoreWebView2ContextMenuItemCollection) GetCount(value *uint32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItemCollection) GetValueAtIndex(index uint32, value **ICoreWebView2ContextMenuItem) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(value)))
	com.AddToScope(value)
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItemCollection) RemoveValueAtIndex(index uint32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(index))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItemCollection) InsertValueAtIndex(index uint32, value *ICoreWebView2ContextMenuItem) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}
