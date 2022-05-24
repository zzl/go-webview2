package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 177CD9E7-B6F5-451A-94A0-5D7A3A4C4141
var IID_ICoreWebView2CookieManager = syscall.GUID{0x177CD9E7, 0xB6F5, 0x451A, 
	[8]byte{0x94, 0xA0, 0x5D, 0x7A, 0x3A, 0x4C, 0x41, 0x41}}

type ICoreWebView2CookieManager struct {
	win32.IUnknown
}

func NewICoreWebView2CookieManager(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2CookieManager {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2CookieManager)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2CookieManager) IID() *syscall.GUID {
	return &IID_ICoreWebView2CookieManager
}

func (this *ICoreWebView2CookieManager) CreateCookie(name string, value string, domain string, path string, cookie **ICoreWebView2Cookie) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(win32.StrToPointer(value)), uintptr(win32.StrToPointer(domain)), uintptr(win32.StrToPointer(path)), uintptr(unsafe.Pointer(cookie)))
		com.AddToScope(cookie)
	return com.Error(ret)
}

func (this *ICoreWebView2CookieManager) CopyCookie(cookieParam *ICoreWebView2Cookie, cookie **ICoreWebView2Cookie) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cookieParam)), uintptr(unsafe.Pointer(cookie)))
		com.AddToScope(cookie)
	return com.Error(ret)
}

func (this *ICoreWebView2CookieManager) GetCookies(uri string, handler *ICoreWebView2GetCookiesCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(uri)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2CookieManager) AddOrUpdateCookie(cookie *ICoreWebView2Cookie) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cookie)))
	return com.Error(ret)
}

func (this *ICoreWebView2CookieManager) DeleteCookie(cookie *ICoreWebView2Cookie) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cookie)))
	return com.Error(ret)
}

func (this *ICoreWebView2CookieManager) DeleteCookies(name string, uri string) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(win32.StrToPointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2CookieManager) DeleteCookiesWithDomainAndPath(name string, domain string, path string) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(win32.StrToPointer(domain)), uintptr(win32.StrToPointer(path)))
	return com.Error(ret)
}

func (this *ICoreWebView2CookieManager) DeleteAllCookies() com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

