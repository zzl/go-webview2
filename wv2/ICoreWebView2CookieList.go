package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// F7F6F714-5D2A-43C6-9503-346ECE02D186
var IID_ICoreWebView2CookieList = syscall.GUID{0xF7F6F714, 0x5D2A, 0x43C6, 
	[8]byte{0x95, 0x03, 0x34, 0x6E, 0xCE, 0x02, 0xD1, 0x86}}

type ICoreWebView2CookieList struct {
	win32.IUnknown
}

func NewICoreWebView2CookieList(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2CookieList {
	p := (*ICoreWebView2CookieList)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2CookieList) IID() *syscall.GUID {
	return &IID_ICoreWebView2CookieList
}

func (this *ICoreWebView2CookieList) GetCount(count *uint32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(count)))
	return com.Error(ret)
}

func (this *ICoreWebView2CookieList) GetValueAtIndex(index uint32, cookie **ICoreWebView2Cookie) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(index), uintptr(unsafe.Pointer(cookie)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*cookie).IUnknown))
	}
	return com.Error(ret)
}

