package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 79110AD3-CD5D-4373-8BC3-C60658F17A5F
var IID_ICoreWebView2Profile = syscall.GUID{0x79110AD3, 0xCD5D, 0x4373,
	[8]byte{0x8B, 0xC3, 0xC6, 0x06, 0x58, 0xF1, 0x7A, 0x5F}}

type ICoreWebView2Profile struct {
	win32.IUnknown
}

func NewICoreWebView2Profile(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Profile {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Profile)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Profile) IID() *syscall.GUID {
	return &IID_ICoreWebView2Profile
}

func (this *ICoreWebView2Profile) GetProfileName(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Profile) GetIsInPrivateModeEnabled(value *int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Profile) GetProfilePath(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Profile) GetDefaultDownloadFolderPath(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Profile) SetDefaultDownloadFolderPath(value string) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Profile) GetPreferredColorScheme(value *int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2Profile) SetPreferredColorScheme(value int32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

