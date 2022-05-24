package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// AD26D6BE-1486-43E6-BF87-A2034006CA21
var IID_ICoreWebView2Cookie = syscall.GUID{0xAD26D6BE, 0x1486, 0x43E6, 
	[8]byte{0xBF, 0x87, 0xA2, 0x03, 0x40, 0x06, 0xCA, 0x21}}

type ICoreWebView2Cookie struct {
	win32.IUnknown
}

func NewICoreWebView2Cookie(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Cookie {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2Cookie)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Cookie) IID() *syscall.GUID {
	return &IID_ICoreWebView2Cookie
}

func (this *ICoreWebView2Cookie) GetName(name *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) GetValue(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) SetValue(value string) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) GetDomain(domain *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(domain)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) GetPath(path *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(path)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) GetExpires(expires *float64) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(expires)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) SetExpires(expires float64) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(expires))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) GetIsHttpOnly(isHttpOnly *int32) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isHttpOnly)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) SetIsHttpOnly(isHttpOnly int32) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(isHttpOnly))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) GetSameSite(sameSite *int32) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(sameSite)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) SetSameSite(sameSite int32) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(sameSite))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) GetIsSecure(isSecure *int32) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isSecure)))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) SetIsSecure(isSecure int32) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(isSecure))
	return com.Error(ret)
}

func (this *ICoreWebView2Cookie) GetIsSession(isSession *int32) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isSession)))
	return com.Error(ret)
}

