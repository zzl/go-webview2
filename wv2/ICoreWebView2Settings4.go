package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// CB56846C-4168-4D53-B04F-03B6D6796FF2
var IID_ICoreWebView2Settings4 = syscall.GUID{0xCB56846C, 0x4168, 0x4D53, 
	[8]byte{0xB0, 0x4F, 0x03, 0xB6, 0xD6, 0x79, 0x6F, 0xF2}}

type ICoreWebView2Settings4 struct {
	ICoreWebView2Settings3
}

func NewICoreWebView2Settings4(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Settings4 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2Settings4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Settings4) IID() *syscall.GUID {
	return &IID_ICoreWebView2Settings4
}

func (this *ICoreWebView2Settings4) GetIsPasswordAutosaveEnabled(value *int32) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings4) SetIsPasswordAutosaveEnabled(value int32) com.Error {
	addr := (*this.LpVtbl)[26]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings4) GetIsGeneralAutofillEnabled(value *int32) com.Error {
	addr := (*this.LpVtbl)[27]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings4) SetIsGeneralAutofillEnabled(value int32) com.Error {
	addr := (*this.LpVtbl)[28]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

