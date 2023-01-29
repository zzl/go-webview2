package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// BC59DB28-BCC3-11EB-8529-0242AC130003
var IID_ICoreWebView2ClientCertificateRequestedEventArgs = syscall.GUID{0xBC59DB28, 0xBCC3, 0x11EB,
	[8]byte{0x85, 0x29, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2ClientCertificateRequestedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2ClientCertificateRequestedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ClientCertificateRequestedEventArgs {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2ClientCertificateRequestedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2ClientCertificateRequestedEventArgs
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetHost(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetPort(value *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetIsProxy(value *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetAllowedCertificateAuthorities(value **ICoreWebView2StringCollection) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	com.AddToScope(value)
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetMutuallyTrustedCertificates(value **ICoreWebView2ClientCertificateCollection) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	com.AddToScope(value)
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetSelectedCertificate(value **ICoreWebView2ClientCertificate) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	com.AddToScope(value)
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) SetSelectedCertificate(value *ICoreWebView2ClientCertificate) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetCancel(value *int32) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) SetCancel(value int32) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetHandled(value *int32) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) SetHandled(value int32) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateRequestedEventArgs) GetDeferral(deferral **ICoreWebView2Deferral) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deferral)))
	com.AddToScope(deferral)
	return com.Error(ret)
}

