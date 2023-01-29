package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 1BF89E2D-1B2B-4629-B28F-05099B41BB03
var IID_ICoreWebView2FrameInfoCollectionIterator = syscall.GUID{0x1BF89E2D, 0x1B2B, 0x4629,
	[8]byte{0xB2, 0x8F, 0x05, 0x09, 0x9B, 0x41, 0xBB, 0x03}}

type ICoreWebView2FrameInfoCollectionIterator struct {
	win32.IUnknown
}

func NewICoreWebView2FrameInfoCollectionIterator(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2FrameInfoCollectionIterator {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2FrameInfoCollectionIterator)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2FrameInfoCollectionIterator) IID() *syscall.GUID {
	return &IID_ICoreWebView2FrameInfoCollectionIterator
}

func (this *ICoreWebView2FrameInfoCollectionIterator) GetHasCurrent(hasCurrent *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hasCurrent)))
	return com.Error(ret)
}

func (this *ICoreWebView2FrameInfoCollectionIterator) GetCurrent(frameInfo **ICoreWebView2FrameInfo) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frameInfo)))
	com.AddToScope(frameInfo)
	return com.Error(ret)
}

func (this *ICoreWebView2FrameInfoCollectionIterator) MoveNext(hasNext *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hasNext)))
	return com.Error(ret)
}

