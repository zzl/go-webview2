package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 7390BB70-ABE0-4843-9529-F143B31B03D6
var IID_ICoreWebView2ScriptDialogOpeningEventArgs = syscall.GUID{0x7390BB70, 0xABE0, 0x4843, 
	[8]byte{0x95, 0x29, 0xF1, 0x43, 0xB3, 0x1B, 0x03, 0xD6}}

type ICoreWebView2ScriptDialogOpeningEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2ScriptDialogOpeningEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ScriptDialogOpeningEventArgs {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2ScriptDialogOpeningEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2ScriptDialogOpeningEventArgs
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) GetUri(uri *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) GetKind(kind *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(kind)))
	return com.Error(ret)
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) GetMessage(message *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(message)))
	return com.Error(ret)
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) Accept() com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) GetDefaultText(defaultText *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(defaultText)))
	return com.Error(ret)
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) GetResultText(resultText *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resultText)))
	return com.Error(ret)
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) SetResultText(resultText string) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(resultText)))
	return com.Error(ret)
}

func (this *ICoreWebView2ScriptDialogOpeningEventArgs) GetDeferral(deferral **ICoreWebView2Deferral) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deferral)))
		com.AddToScope(deferral)
	return com.Error(ret)
}

