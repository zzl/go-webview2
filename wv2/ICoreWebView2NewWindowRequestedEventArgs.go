package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 34ACB11C-FC37-4418-9132-F9C21D1EAFB9
var IID_ICoreWebView2NewWindowRequestedEventArgs = syscall.GUID{0x34ACB11C, 0xFC37, 0x4418, 
	[8]byte{0x91, 0x32, 0xF9, 0xC2, 0x1D, 0x1E, 0xAF, 0xB9}}

type ICoreWebView2NewWindowRequestedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2NewWindowRequestedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2NewWindowRequestedEventArgs {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2NewWindowRequestedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2NewWindowRequestedEventArgs
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) GetUri(uri *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) SetNewWindow(newWindow *ICoreWebView2) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newWindow)))
	return com.Error(ret)
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) GetNewWindow(newWindow **ICoreWebView2) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(newWindow)))
		com.AddToScope(newWindow)
	return com.Error(ret)
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) SetHandled(handled int32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(handled))
	return com.Error(ret)
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) GetHandled(handled *int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handled)))
	return com.Error(ret)
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) GetIsUserInitiated(isUserInitiated *int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isUserInitiated)))
	return com.Error(ret)
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) GetDeferral(deferral **ICoreWebView2Deferral) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deferral)))
		com.AddToScope(deferral)
	return com.Error(ret)
}

func (this *ICoreWebView2NewWindowRequestedEventArgs) GetWindowFeatures(value **ICoreWebView2WindowFeatures) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
		com.AddToScope(value)
	return com.Error(ret)
}

