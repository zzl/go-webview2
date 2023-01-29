package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F06F41BF-4B5A-49D8-B9F6-FA16CD29F274
var IID_ICoreWebView2Environment9 = syscall.GUID{0xF06F41BF, 0x4B5A, 0x49D8,
	[8]byte{0xB9, 0xF6, 0xFA, 0x16, 0xCD, 0x29, 0xF2, 0x74}}

type ICoreWebView2Environment9 struct {
	ICoreWebView2Environment8
}

func NewICoreWebView2Environment9(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment9 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Environment9)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment9) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment9
}

func (this *ICoreWebView2Environment9) CreateContextMenuItem(label string, iconStream *win32.IStream, kind int32, item **ICoreWebView2ContextMenuItem) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(label)), uintptr(unsafe.Pointer(iconStream)), uintptr(kind), uintptr(unsafe.Pointer(item)))
	com.AddToScope(item)
	return com.Error(ret)
}
