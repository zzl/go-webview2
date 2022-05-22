package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 4DAB9422-46FA-4C3E-A5D2-41D2071D3680
var IID_ICoreWebView2ProcessFailedEventArgs2 = syscall.GUID{0x4DAB9422, 0x46FA, 0x4C3E, 
	[8]byte{0xA5, 0xD2, 0x41, 0xD2, 0x07, 0x1D, 0x36, 0x80}}

type ICoreWebView2ProcessFailedEventArgs2 struct {
	ICoreWebView2ProcessFailedEventArgs
}

func NewICoreWebView2ProcessFailedEventArgs2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ProcessFailedEventArgs2 {
	p := (*ICoreWebView2ProcessFailedEventArgs2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ProcessFailedEventArgs2) IID() *syscall.GUID {
	return &IID_ICoreWebView2ProcessFailedEventArgs2
}

func (this *ICoreWebView2ProcessFailedEventArgs2) GetReason(reason *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(reason)))
	return com.Error(ret)
}

func (this *ICoreWebView2ProcessFailedEventArgs2) GetExitCode(exitCode *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(exitCode)))
	return com.Error(ret)
}

func (this *ICoreWebView2ProcessFailedEventArgs2) GetProcessDescription(processDescription *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(processDescription)))
	return com.Error(ret)
}

func (this *ICoreWebView2ProcessFailedEventArgs2) GetFrameInfosForFailedProcess(frames **ICoreWebView2FrameInfoCollection) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(frames)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*frames).IUnknown))
	}
	return com.Error(ret)
}

