package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// B50D82CC-CC28-481D-9614-CB048895E6A0
var IID_ICoreWebView2Frame3 = syscall.GUID{0xB50D82CC, 0xCC28, 0x481D,
	[8]byte{0x96, 0x14, 0xCB, 0x04, 0x88, 0x95, 0xE6, 0xA0}}

type ICoreWebView2Frame3 struct {
	ICoreWebView2Frame2
}

func NewICoreWebView2Frame3(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Frame3 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Frame3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Frame3) IID() *syscall.GUID {
	return &IID_ICoreWebView2Frame3
}

func (this *ICoreWebView2Frame3) Add_PermissionRequested(handler *ICoreWebView2FramePermissionRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame3) Remove_PermissionRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}
