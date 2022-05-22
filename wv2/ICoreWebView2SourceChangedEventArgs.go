package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 31E0E545-1DBA-4266-8914-F63848A1F7D7
var IID_ICoreWebView2SourceChangedEventArgs = syscall.GUID{0x31E0E545, 0x1DBA, 0x4266, 
	[8]byte{0x89, 0x14, 0xF6, 0x38, 0x48, 0xA1, 0xF7, 0xD7}}

type ICoreWebView2SourceChangedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2SourceChangedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2SourceChangedEventArgs {
	p := (*ICoreWebView2SourceChangedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2SourceChangedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2SourceChangedEventArgs
}

func (this *ICoreWebView2SourceChangedEventArgs) GetIsNewDocument(isNewDocument *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isNewDocument)))
	return com.Error(ret)
}

