package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 07023F7D-2D77-4D67-9040-6E7D428C6A40
var IID_ICoreWebView2BasicAuthenticationResponse = syscall.GUID{0x07023F7D, 0x2D77, 0x4D67, 
	[8]byte{0x90, 0x40, 0x6E, 0x7D, 0x42, 0x8C, 0x6A, 0x40}}

type ICoreWebView2BasicAuthenticationResponse struct {
	win32.IUnknown
}

func NewICoreWebView2BasicAuthenticationResponse(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2BasicAuthenticationResponse {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2BasicAuthenticationResponse)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2BasicAuthenticationResponse) IID() *syscall.GUID {
	return &IID_ICoreWebView2BasicAuthenticationResponse
}

func (this *ICoreWebView2BasicAuthenticationResponse) GetUserName(userName *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(userName)))
	return com.Error(ret)
}

func (this *ICoreWebView2BasicAuthenticationResponse) SetUserName(userName string) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(userName)))
	return com.Error(ret)
}

func (this *ICoreWebView2BasicAuthenticationResponse) GetPassword(password *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(password)))
	return com.Error(ret)
}

func (this *ICoreWebView2BasicAuthenticationResponse) SetPassword(password string) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(password)))
	return com.Error(ret)
}

