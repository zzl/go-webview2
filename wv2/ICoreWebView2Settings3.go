package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// FDB5AB74-AF33-4854-84F0-0A631DEB5EBA
var IID_ICoreWebView2Settings3 = syscall.GUID{0xFDB5AB74, 0xAF33, 0x4854, 
	[8]byte{0x84, 0xF0, 0x0A, 0x63, 0x1D, 0xEB, 0x5E, 0xBA}}

type ICoreWebView2Settings3 struct {
	ICoreWebView2Settings2
}

func NewICoreWebView2Settings3(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Settings3 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2Settings3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Settings3) IID() *syscall.GUID {
	return &IID_ICoreWebView2Settings3
}

func (this *ICoreWebView2Settings3) GetAreBrowserAcceleratorKeysEnabled(areBrowserAcceleratorKeysEnabled *int32) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(areBrowserAcceleratorKeysEnabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings3) SetAreBrowserAcceleratorKeysEnabled(areBrowserAcceleratorKeysEnabled int32) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(areBrowserAcceleratorKeysEnabled))
	return com.Error(ret)
}

