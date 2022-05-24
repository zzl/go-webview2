package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 79C24D83-09A3-45AE-9418-487F32A58740
var IID_ICoreWebView2_7 = syscall.GUID{0x79C24D83, 0x09A3, 0x45AE, 
	[8]byte{0x94, 0x18, 0x48, 0x7F, 0x32, 0xA5, 0x87, 0x40}}

type ICoreWebView2_7 struct {
	ICoreWebView2_6
}

func NewICoreWebView2_7(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_7 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2_7)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_7) IID() *syscall.GUID {
	return &IID_ICoreWebView2_7
}

func (this *ICoreWebView2_7) PrintToPdf(resultFilePath string, printSettings *ICoreWebView2PrintSettings, handler *ICoreWebView2PrintToPdfCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[80]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(resultFilePath)), uintptr(unsafe.Pointer(printSettings)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

