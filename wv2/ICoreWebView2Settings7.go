package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 488DC902-35EF-42D2-BC7D-94B65C4BC49C
var IID_ICoreWebView2Settings7 = syscall.GUID{0x488DC902, 0x35EF, 0x42D2,
	[8]byte{0xBC, 0x7D, 0x94, 0xB6, 0x5C, 0x4B, 0xC4, 0x9C}}

type ICoreWebView2Settings7 struct {
	ICoreWebView2Settings6
}

func NewICoreWebView2Settings7(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Settings7 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Settings7)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Settings7) IID() *syscall.GUID {
	return &IID_ICoreWebView2Settings7
}

func (this *ICoreWebView2Settings7) GetHiddenPdfToolbarItems(hidden_pdf_toolbar_items *int32) com.Error {
	addr := (*this.LpVtbl)[33]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(hidden_pdf_toolbar_items)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings7) SetHiddenPdfToolbarItems(hidden_pdf_toolbar_items int32) com.Error {
	addr := (*this.LpVtbl)[34]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(hidden_pdf_toolbar_items))
	return com.Error(ret)
}

