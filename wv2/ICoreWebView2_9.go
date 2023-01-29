package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 4D7B2EAB-9FDC-468D-B998-A9260B5ED651
var IID_ICoreWebView2_9 = syscall.GUID{0x4D7B2EAB, 0x9FDC, 0x468D,
	[8]byte{0xB9, 0x98, 0xA9, 0x26, 0x0B, 0x5E, 0xD6, 0x51}}

type ICoreWebView2_9 struct {
	ICoreWebView2_8
}

func NewICoreWebView2_9(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2_9 {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2_9)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2_9) IID() *syscall.GUID {
	return &IID_ICoreWebView2_9
}

func (this *ICoreWebView2_9) Add_IsDefaultDownloadDialogOpenChanged(handler *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[88]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_9) Remove_IsDefaultDownloadDialogOpenChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[89]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2_9) GetIsDefaultDownloadDialogOpen(value *int32) com.Error {
	addr := (*this.LpVtbl)[90]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2_9) OpenDefaultDownloadDialog() com.Error {
	addr := (*this.LpVtbl)[91]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2_9) CloseDefaultDownloadDialog() com.Error {
	addr := (*this.LpVtbl)[92]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2_9) GetDefaultDownloadDialogCornerAlignment(value *int32) com.Error {
	addr := (*this.LpVtbl)[93]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2_9) SetDefaultDownloadDialogCornerAlignment(value int32) com.Error {
	addr := (*this.LpVtbl)[94]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2_9) GetDefaultDownloadDialogMargin(value *TagPOINT) com.Error {
	addr := (*this.LpVtbl)[95]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2_9) SetDefaultDownloadDialogMargin(value TagPOINT) com.Error {
	addr := (*this.LpVtbl)[96]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&value)))
	return com.Error(ret)
}
