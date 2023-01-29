package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 402B99CD-A0CC-4FA5-B7A5-51D86A1D2339
var IID_ICoreWebView2ProcessInfoCollection = syscall.GUID{0x402B99CD, 0xA0CC, 0x4FA5,
	[8]byte{0xB7, 0xA5, 0x51, 0xD8, 0x6A, 0x1D, 0x23, 0x39}}

type ICoreWebView2ProcessInfoCollection struct {
	win32.IUnknown
}

func NewICoreWebView2ProcessInfoCollection(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ProcessInfoCollection {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2ProcessInfoCollection)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ProcessInfoCollection) IID() *syscall.GUID {
	return &IID_ICoreWebView2ProcessInfoCollection
}

func (this *ICoreWebView2ProcessInfoCollection) GetCount(count *uint32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return com.Error(ret)
}

func (this *ICoreWebView2ProcessInfoCollection) GetValueAtIndex(index uint32, processInfo **ICoreWebView2ProcessInfo) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(processInfo)))
	com.AddToScope(processInfo)
	return com.Error(ret)
}
