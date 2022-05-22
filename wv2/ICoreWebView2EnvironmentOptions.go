package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 2FDE08A8-1E9A-4766-8C05-95A9CEB9D1C5
var IID_ICoreWebView2EnvironmentOptions = syscall.GUID{0x2FDE08A8, 0x1E9A, 0x4766, 
	[8]byte{0x8C, 0x05, 0x95, 0xA9, 0xCE, 0xB9, 0xD1, 0xC5}}

type ICoreWebView2EnvironmentOptions struct {
	win32.IUnknown
}

func NewICoreWebView2EnvironmentOptions(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2EnvironmentOptions {
	p := (*ICoreWebView2EnvironmentOptions)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2EnvironmentOptions) IID() *syscall.GUID {
	return &IID_ICoreWebView2EnvironmentOptions
}

func (this *ICoreWebView2EnvironmentOptions) GetAdditionalBrowserArguments(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2EnvironmentOptions) SetAdditionalBrowserArguments(value string) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2EnvironmentOptions) GetLanguage(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2EnvironmentOptions) SetLanguage(value string) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2EnvironmentOptions) GetTargetCompatibleBrowserVersion(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2EnvironmentOptions) SetTargetCompatibleBrowserVersion(value string) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2EnvironmentOptions) GetAllowSingleSignOnUsingOSPrimaryAccount(allow *int32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(allow)))
	return com.Error(ret)
}

func (this *ICoreWebView2EnvironmentOptions) SetAllowSingleSignOnUsingOSPrimaryAccount(allow int32) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(allow))
	return com.Error(ret)
}

