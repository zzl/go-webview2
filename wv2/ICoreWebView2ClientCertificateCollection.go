package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// EF5674D2-BCC3-11EB-8529-0242AC130003
var IID_ICoreWebView2ClientCertificateCollection = syscall.GUID{0xEF5674D2, 0xBCC3, 0x11EB, 
	[8]byte{0x85, 0x29, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2ClientCertificateCollection struct {
	win32.IUnknown
}

func NewICoreWebView2ClientCertificateCollection(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ClientCertificateCollection {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2ClientCertificateCollection)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ClientCertificateCollection) IID() *syscall.GUID {
	return &IID_ICoreWebView2ClientCertificateCollection
}

func (this *ICoreWebView2ClientCertificateCollection) GetCount(value *uint32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ClientCertificateCollection) GetValueAtIndex(index uint32, certificate **ICoreWebView2ClientCertificate) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(certificate)))
		com.AddToScope(certificate)
	return com.Error(ret)
}

