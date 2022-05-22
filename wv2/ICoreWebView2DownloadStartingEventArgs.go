package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E99BBE21-43E9-4544-A732-282764EAFA60
var IID_ICoreWebView2DownloadStartingEventArgs = syscall.GUID{0xE99BBE21, 0x43E9, 0x4544, 
	[8]byte{0xA7, 0x32, 0x28, 0x27, 0x64, 0xEA, 0xFA, 0x60}}

type ICoreWebView2DownloadStartingEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2DownloadStartingEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2DownloadStartingEventArgs {
	p := (*ICoreWebView2DownloadStartingEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2DownloadStartingEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2DownloadStartingEventArgs
}

func (this *ICoreWebView2DownloadStartingEventArgs) GetDownloadOperation(downloadOperation **ICoreWebView2DownloadOperation) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(downloadOperation)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*downloadOperation).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadStartingEventArgs) GetCancel(cancel *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cancel)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadStartingEventArgs) SetCancel(cancel int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(cancel))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadStartingEventArgs) GetResultFilePath(resultFilePath *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(resultFilePath)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadStartingEventArgs) SetResultFilePath(resultFilePath string) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(resultFilePath)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadStartingEventArgs) GetHandled(handled *int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handled)))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadStartingEventArgs) SetHandled(handled int32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(handled))
	return com.Error(ret)
}

func (this *ICoreWebView2DownloadStartingEventArgs) GetDeferral(deferral **ICoreWebView2Deferral) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deferral)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*deferral).IUnknown))
	}
	return com.Error(ret)
}

