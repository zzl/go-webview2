package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 973AE2EF-FF18-4894-8FB2-3C758F046810
var IID_ICoreWebView2PermissionRequestedEventArgs = syscall.GUID{0x973AE2EF, 0xFF18, 0x4894,
	[8]byte{0x8F, 0xB2, 0x3C, 0x75, 0x8F, 0x04, 0x68, 0x10}}

type ICoreWebView2PermissionRequestedEventArgs struct {
	win32.IUnknown
}

func NewICoreWebView2PermissionRequestedEventArgs(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2PermissionRequestedEventArgs {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2PermissionRequestedEventArgs)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2PermissionRequestedEventArgs) IID() *syscall.GUID {
	return &IID_ICoreWebView2PermissionRequestedEventArgs
}

func (this *ICoreWebView2PermissionRequestedEventArgs) GetUri(uri *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(uri)))
	return com.Error(ret)
}

func (this *ICoreWebView2PermissionRequestedEventArgs) GetPermissionKind(permissionKind *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(permissionKind)))
	return com.Error(ret)
}

func (this *ICoreWebView2PermissionRequestedEventArgs) GetIsUserInitiated(isUserInitiated *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isUserInitiated)))
	return com.Error(ret)
}

func (this *ICoreWebView2PermissionRequestedEventArgs) GetState(state *int32) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(state)))
	return com.Error(ret)
}

func (this *ICoreWebView2PermissionRequestedEventArgs) SetState(state int32) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(state))
	return com.Error(ret)
}

func (this *ICoreWebView2PermissionRequestedEventArgs) GetDeferral(deferral **ICoreWebView2Deferral) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(deferral)))
	com.AddToScope(deferral)
	return com.Error(ret)
}
