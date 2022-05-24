package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 8F834154-D38E-4D90-AFFB-6800A7272839
var IID_ICoreWebView2FrameInfoCollection = syscall.GUID{0x8F834154, 0xD38E, 0x4D90, 
	[8]byte{0xAF, 0xFB, 0x68, 0x00, 0xA7, 0x27, 0x28, 0x39}}

type ICoreWebView2FrameInfoCollection struct {
	win32.IUnknown
}

func NewICoreWebView2FrameInfoCollection(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2FrameInfoCollection {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2FrameInfoCollection)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2FrameInfoCollection) IID() *syscall.GUID {
	return &IID_ICoreWebView2FrameInfoCollection
}

func (this *ICoreWebView2FrameInfoCollection) GetIterator(iterator **ICoreWebView2FrameInfoCollectionIterator) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(iterator)))
		com.AddToScope(iterator)
	return com.Error(ret)
}

