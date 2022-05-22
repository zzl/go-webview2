package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 0F99A40C-E962-4207-9E92-E3D542EFF849
var IID_ICoreWebView2WebMessageReceivedEventArgs = syscall.GUID{0x0F99A40C, 0xE962, 0x4207, 
	[8]byte{0x9E, 0x92, 0xE3, 0xD5, 0x42, 0xEF, 0xF8, 0x49}}

type ICoreWebView2WebMessageReceivedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2WebMessageReceivedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2WebMessageReceivedEventArgs {
	p := (*ICoreWebView2WebMessageReceivedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2WebMessageReceivedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2WebMessageReceivedEventArgs
}

func (this *ICoreWebView2WebMessageReceivedEventArgs) GetSource(source *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(source)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebMessageReceivedEventArgs) GetWebMessageAsJson(webMessageAsJson *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webMessageAsJson)))
	return com.Error(ret)
}

func (this *ICoreWebView2WebMessageReceivedEventArgs) TryGetWebMessageAsString(webMessageAsString *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(webMessageAsString)))
	return com.Error(ret)
}

