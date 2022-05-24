package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-com/ole"
	"syscall"
	"unsafe"
)

// F1131A5E-9BA9-11EB-A8B3-0242AC130003
var IID_ICoreWebView2Frame = syscall.GUID{0xF1131A5E, 0x9BA9, 0x11EB, 
	[8]byte{0xA8, 0xB3, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2Frame struct {
	win32.IUnknown
}

func NewICoreWebView2Frame(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Frame {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2Frame)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Frame) IID() *syscall.GUID {
	return &IID_ICoreWebView2Frame
}

func (this *ICoreWebView2Frame) GetName(name *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(name)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame) Add_NameChanged(eventHandler *ICoreWebView2FrameNameChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame) Remove_NameChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame) AddHostObjectToScriptWithOrigins(name string, object *ole.Variant, originsCount uint32, origins *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)), uintptr(unsafe.Pointer(object)), uintptr(originsCount), uintptr(unsafe.Pointer(origins)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame) RemoveHostObjectFromScript(name string) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(name)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame) Add_Destroyed(eventHandler *ICoreWebView2FrameDestroyedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame) Remove_Destroyed(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame) IsDestroyed(destroyed *int32) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(destroyed)))
	return com.Error(ret)
}

