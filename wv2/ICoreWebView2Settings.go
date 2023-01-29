package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E562E4F0-D7FA-43AC-8D71-C05150499F00
var IID_ICoreWebView2Settings = syscall.GUID{0xE562E4F0, 0xD7FA, 0x43AC,
	[8]byte{0x8D, 0x71, 0xC0, 0x51, 0x50, 0x49, 0x9F, 0x00}}

type ICoreWebView2Settings struct {
	win32.IUnknown
}

func NewICoreWebView2Settings(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Settings {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Settings)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Settings) IID() *syscall.GUID {
	return &IID_ICoreWebView2Settings
}

func (this *ICoreWebView2Settings) GetIsScriptEnabled(isScriptEnabled *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isScriptEnabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetIsScriptEnabled(isScriptEnabled int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(isScriptEnabled))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) GetIsWebMessageEnabled(isWebMessageEnabled *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isWebMessageEnabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetIsWebMessageEnabled(isWebMessageEnabled int32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(isWebMessageEnabled))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) GetAreDefaultScriptDialogsEnabled(areDefaultScriptDialogsEnabled *int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(areDefaultScriptDialogsEnabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetAreDefaultScriptDialogsEnabled(areDefaultScriptDialogsEnabled int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(areDefaultScriptDialogsEnabled))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) GetIsStatusBarEnabled(isStatusBarEnabled *int32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isStatusBarEnabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetIsStatusBarEnabled(isStatusBarEnabled int32) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(isStatusBarEnabled))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) GetAreDevToolsEnabled(areDevToolsEnabled *int32) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(areDevToolsEnabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetAreDevToolsEnabled(areDevToolsEnabled int32) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(areDevToolsEnabled))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) GetAreDefaultContextMenusEnabled(enabled *int32) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(enabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetAreDefaultContextMenusEnabled(enabled int32) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(enabled))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) GetAreHostObjectsAllowed(allowed *int32) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(allowed)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetAreHostObjectsAllowed(allowed int32) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(allowed))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) GetIsZoomControlEnabled(enabled *int32) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(enabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetIsZoomControlEnabled(enabled int32) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(enabled))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) GetIsBuiltInErrorPageEnabled(enabled *int32) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(enabled)))
	return com.Error(ret)
}

func (this *ICoreWebView2Settings) SetIsBuiltInErrorPageEnabled(enabled int32) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(enabled))
	return com.Error(ret)
}

