package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E7188076-BCC3-11EB-8529-0242AC130003
var IID_ICoreWebView2ClientCertificate = syscall.GUID{0xE7188076, 0xBCC3, 0x11EB, 
	[8]byte{0x85, 0x29, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2ClientCertificate struct {
	win32.IUnknown
}

func NewICoreWebView2ClientCertificate(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ClientCertificate {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2ClientCertificate)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ClientCertificate) IID() *syscall.GUID {
	return &IID_ICoreWebView2ClientCertificate
}

func (this *ICoreWebView2ClientCertificate) GetSubject(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificate) GetIssuer(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificate) GetValidFrom(value *float64) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificate) GetValidTo(value *float64) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificate) GetDerEncodedSerialNumber(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificate) GetDisplayName(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificate) ToPemEncoding(pemEncodedData *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(pemEncodedData)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificate) GetPemEncodedIssuerCertificateChain(value **ICoreWebView2StringCollection) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
		com.AddToScope(value)
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificate) GetKind(value *int32) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

