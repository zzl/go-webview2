package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 3DF9B733-B9AE-4A15-86B4-EB9EE9826469
var IID_ICoreWebView2CompositionController = syscall.GUID{0x3DF9B733, 0xB9AE, 0x4A15, 
	[8]byte{0x86, 0xB4, 0xEB, 0x9E, 0xE9, 0x82, 0x64, 0x69}}

type ICoreWebView2CompositionController struct {
	win32.IUnknown
}

func NewICoreWebView2CompositionController(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2CompositionController {
	 if pUnk == nil {
		return nil;
	}
	p := (*ICoreWebView2CompositionController)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2CompositionController) IID() *syscall.GUID {
	return &IID_ICoreWebView2CompositionController
}

func (this *ICoreWebView2CompositionController) GetRootVisualTarget(target **win32.IUnknown) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)))
		com.AddToScope(target)
	return com.Error(ret)
}

func (this *ICoreWebView2CompositionController) SetRootVisualTarget(target *win32.IUnknown) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(target)))
	return com.Error(ret)
}

func (this *ICoreWebView2CompositionController) SendMouseInput(eventKind int32, virtualKeys int32, mouseData uint32, point TagPOINT) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(eventKind), uintptr(virtualKeys), uintptr(mouseData), *(*uintptr)(unsafe.Pointer(&point)))
	return com.Error(ret)
}

func (this *ICoreWebView2CompositionController) SendPointerInput(eventKind int32, pointerInfo *ICoreWebView2PointerInfo) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(eventKind), uintptr(unsafe.Pointer(pointerInfo)))
	return com.Error(ret)
}

func (this *ICoreWebView2CompositionController) GetCursor(cursor *win32.HICON) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(cursor)))
	return com.Error(ret)
}

func (this *ICoreWebView2CompositionController) GetSystemCursorId(systemCursorId *uint32) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(systemCursorId)))
	return com.Error(ret)
}

func (this *ICoreWebView2CompositionController) Add_CursorChanged(eventHandler *ICoreWebView2CursorChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2CompositionController) Remove_CursorChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

