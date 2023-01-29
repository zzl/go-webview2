package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 4D00C0D1-9434-4EB6-8078-8697A560334F
var IID_ICoreWebView2Controller = syscall.GUID{0x4D00C0D1, 0x9434, 0x4EB6,
	[8]byte{0x80, 0x78, 0x86, 0x97, 0xA5, 0x60, 0x33, 0x4F}}

type ICoreWebView2Controller struct {
	win32.IUnknown
}

func NewICoreWebView2Controller(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Controller {
	if pUnk == nil {
		return nil
	}
	p := (*ICoreWebView2Controller)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Controller) IID() *syscall.GUID {
	return &IID_ICoreWebView2Controller
}

func (this *ICoreWebView2Controller) GetIsVisible(isVisible *int32) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(isVisible)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) SetIsVisible(isVisible int32) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(isVisible))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) GetBounds(bounds *TagRECT) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(bounds)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) SetBounds(bounds TagRECT) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&bounds)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) GetZoomFactor(zoomFactor *float64) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(zoomFactor)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) SetZoomFactor(zoomFactor float64) com.Error {
	addr := (*this.LpVtbl)[8]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(zoomFactor))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Add_ZoomFactorChanged(eventHandler *ICoreWebView2ZoomFactorChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[9]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Remove_ZoomFactorChanged(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[10]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) SetBoundsAndZoomFactor(bounds TagRECT, zoomFactor float64) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), (uintptr)(unsafe.Pointer(&bounds)), uintptr(zoomFactor))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) MoveFocus(reason int32) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(reason))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Add_MoveFocusRequested(eventHandler *ICoreWebView2MoveFocusRequestedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Remove_MoveFocusRequested(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Add_GotFocus(eventHandler *ICoreWebView2FocusChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Remove_GotFocus(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Add_LostFocus(eventHandler *ICoreWebView2FocusChangedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Remove_LostFocus(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Add_AcceleratorKeyPressed(eventHandler *ICoreWebView2AcceleratorKeyPressedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Remove_AcceleratorKeyPressed(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) GetParentWindow(parentWindow *win32.HWND) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(parentWindow)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) SetParentWindow(parentWindow win32.HWND) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(parentWindow))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) NotifyParentWindowPositionChanged() com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) Close() com.Error {
	addr := (*this.LpVtbl)[24]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)))
	return com.Error(ret)
}

func (this *ICoreWebView2Controller) GetCoreWebView2(coreWebView2 **ICoreWebView2) com.Error {
	addr := (*this.LpVtbl)[25]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(coreWebView2)))
	com.AddToScope(coreWebView2)
	return com.Error(ret)
}

