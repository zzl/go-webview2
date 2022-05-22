package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// EE9A0F68-F46C-4E32-AC23-EF8CAC224D2A
var IID_ICoreWebView2Settings2 = syscall.GUID{0xEE9A0F68, 0xF46C, 0x4E32, 
	[8]byte{0xAC, 0x23, 0xEF, 0x8C, 0xAC, 0x22, 0x4D, 0x2A}}

type ICoreWebView2Settings2 struct {
	ICoreWebView2Settings
}

func NewICoreWebView2Settings2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Settings2 {
	p := (*ICoreWebView2Settings2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Settings2) IID() *syscall.GUID {
	return &IID_ICoreWebView2Settings2
}

func (this *ICoreWebView2Settings2) GetUserAgent(userAgent *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(userAgent)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings2) SetUserAgent(userAgent string) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(userAgent)))
	return com.Error(ret)
}

