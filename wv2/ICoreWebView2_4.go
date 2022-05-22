package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 20D02D59-6DF2-42DC-BD06-F98A694B1302
var IID_ICoreWebView2_4 = syscall.GUID{0x20D02D59, 0x6DF2, 0x42DC, 
	[8]byte{0xBD, 0x06, 0xF9, 0x8A, 0x69, 0x4B, 0x13, 0x02}}

type ICoreWebView2_4 struct {
	ICoreWebView2_3
}

func NewICoreWebView2_4(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_4 {
	p := (*ICoreWebView2_4)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_4) IID() *syscall.GUID {
	return &IID_ICoreWebView2_4
}

func (this *ICoreWebView2_4) Add_FrameCreated(eventHandler *ICoreWebView2FrameCreatedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[73]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_4) Remove_FrameCreated(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[74]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_4) Add_DownloadStarting(eventHandler *ICoreWebView2DownloadStartingEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[75]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_4) Remove_DownloadStarting(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[76]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

