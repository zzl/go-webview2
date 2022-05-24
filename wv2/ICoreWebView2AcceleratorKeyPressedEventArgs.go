package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 9F760F8A-FB79-42BE-9990-7B56900FA9C7
var IID_ICoreWebView2AcceleratorKeyPressedEventArgs = syscall.GUID{0x9F760F8A, 0xFB79, 0x42BE, 
	[8]byte{0x99, 0x90, 0x7B, 0x56, 0x90, 0x0F, 0xA9, 0xC7}}

type ICoreWebView2AcceleratorKeyPressedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2AcceleratorKeyPressedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2AcceleratorKeyPressedEventArgs {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2AcceleratorKeyPressedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2AcceleratorKeyPressedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2AcceleratorKeyPressedEventArgs
}

func (this *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventKind(keyEventKind *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(keyEventKind)))
	return com.Error(ret)
}

func (this *ICoreWebView2AcceleratorKeyPressedEventArgs) GetVirtualKey(virtualKey *uint32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(virtualKey)))
	return com.Error(ret)
}

func (this *ICoreWebView2AcceleratorKeyPressedEventArgs) GetKeyEventLParam(lParam *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(lParam)))
	return com.Error(ret)
}

func (this *ICoreWebView2AcceleratorKeyPressedEventArgs) GetPhysicalKeyStatus(physicalKeyStatus *COREWEBVIEW2_PHYSICAL_KEY_STATUS) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(physicalKeyStatus)))
	return com.Error(ret)
}

func (this *ICoreWebView2AcceleratorKeyPressedEventArgs) GetHandled(handled *int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handled)))
	return com.Error(ret)
}

func (this *ICoreWebView2AcceleratorKeyPressedEventArgs) SetHandled(handled int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(handled))
	return com.Error(ret)
}

