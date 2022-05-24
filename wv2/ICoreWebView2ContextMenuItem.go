package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 7AED49E3-A93F-497A-811C-749C6B6B6C65
var IID_ICoreWebView2ContextMenuItem = syscall.GUID{0x7AED49E3, 0xA93F, 0x497A, 
	[8]byte{0x81, 0x1C, 0x74, 0x9C, 0x6B, 0x6B, 0x6C, 0x65}}

type ICoreWebView2ContextMenuItem struct {
	win32.IUnknown
}

func NewICoreWebView2ContextMenuItem(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2ContextMenuItem {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2ContextMenuItem)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2ContextMenuItem) IID() *syscall.GUID {
	return &IID_ICoreWebView2ContextMenuItem
}

func (this *ICoreWebView2ContextMenuItem) GetName(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) GetLabel(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) GetCommandId(value *int32) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) GetShortcutKeyDescription(value *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) GetIcon(value **win32.IUnknown) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
		com.AddToScope(value)
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) GetKind(value *int32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) SetIsEnabled(value int32) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) GetIsEnabled(value *int32) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) SetIsChecked(value int32) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(value))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) GetIsChecked(value *int32) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) GetChildren(value **ICoreWebView2ContextMenuItemCollection) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(value)))
		com.AddToScope(value)
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) Add_CustomItemSelected(eventHandler *ICoreWebView2CustomItemSelectedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2ContextMenuItem) Remove_CustomItemSelected(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

