package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 12AAE616-8CCB-44EC-BCB3-EB1831881635
var IID_ICoreWebView2ControllerOptions = syscall.GUID{0x12AAE616, 0x8CCB, 0x44EC,
	[8]byte{0xBC, 0xB3, 0xEB, 0x18, 0x31, 0x88, 0x16, 0x35}}

type ICoreWebView2ControllerOptions struct {
	win32.IUnknown
}

func NewICoreWebView2ControllerOptions(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ControllerOptions {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2ControllerOptions)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ControllerOptions) IID() *syscall.GUID {
	return &IID_ICoreWebView2ControllerOptions
}

func (this *ICoreWebView2ControllerOptions) GetProfileName(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ControllerOptions) SetProfileName(value string) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ControllerOptions) GetIsInPrivateModeEnabled(value *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ControllerOptions) SetIsInPrivateModeEnabled(value int32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

