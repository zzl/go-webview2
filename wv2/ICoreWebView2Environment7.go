package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 43C22296-3BBD-43A4-9C00-5C0DF6DD29A2
var IID_ICoreWebView2Environment7 = syscall.GUID{0x43C22296, 0x3BBD, 0x43A4, 
	[8]byte{0x9C, 0x00, 0x5C, 0x0D, 0xF6, 0xDD, 0x29, 0xA2}}

type ICoreWebView2Environment7 struct {
	ICoreWebView2Environment6
}

func NewICoreWebView2Environment7(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment7 {
	p := (*ICoreWebView2Environment7)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment7) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment7
}

func (this *ICoreWebView2Environment7) GetUserDataFolder(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

