package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// A0D6DF20-3B92-416D-AA0C-437A9C727857
var IID_ICoreWebView2_3 = syscall.GUID{0xA0D6DF20, 0x3B92, 0x416D, 
	[8]byte{0xAA, 0x0C, 0x43, 0x7A, 0x9C, 0x72, 0x78, 0x57}}

type ICoreWebView2_3 struct {
	ICoreWebView2_2
}

func NewICoreWebView2_3(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_3 {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2_3)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_3) IID() *syscall.GUID {
	return &IID_ICoreWebView2_3
}

func (this *ICoreWebView2_3) TrySuspend(handler *ICoreWebView2TrySuspendCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[68]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2_3) Resume() com.Error {
	addr := (*this.LpVtbl)[69]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2_3) GetIsSuspended(isSuspended *int32) com.Error {
	addr := (*this.LpVtbl)[70]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isSuspended)))
	return com.Error(ret)
}

func (this *ICoreWebView2_3) SetVirtualHostNameToFolderMapping(hostName string, folderPath string, accessKind int32) com.Error {
	addr := (*this.LpVtbl)[71]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(hostName)), uintptr(win32.StrToPointer(folderPath)), uintptr(accessKind))
	return com.Error(ret)
}

func (this *ICoreWebView2_3) ClearVirtualHostNameToFolderMapping(hostName string) com.Error {
	addr := (*this.LpVtbl)[72]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(hostName)))
	return com.Error(ret)
}

