package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// EF05516F-D897-4F9E-B672-D8E2307A3FB0
var IID_ICoreWebView2BasicAuthenticationRequestedEventArgs = syscall.GUID{0xEF05516F, 0xD897, 0x4F9E, 
	[8]byte{0xB6, 0x72, 0xD8, 0xE2, 0x30, 0x7A, 0x3F, 0xB0}}

type ICoreWebView2BasicAuthenticationRequestedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2BasicAuthenticationRequestedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2BasicAuthenticationRequestedEventArgs {
	p := (*ICoreWebView2BasicAuthenticationRequestedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2BasicAuthenticationRequestedEventArgs
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetUri(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetChallenge(challenge *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(challenge)))
	return com.Error(ret)
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetResponse(response **ICoreWebView2BasicAuthenticationResponse) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(response)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*response).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetCancel(cancel *int32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cancel)))
	return com.Error(ret)
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventArgs) SetCancel(cancel int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(cancel))
	return com.Error(ret)
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventArgs) GetDeferral(deferral **ICoreWebView2Deferral) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deferral)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*deferral).IUnknown))
	}
	return com.Error(ret)
}

