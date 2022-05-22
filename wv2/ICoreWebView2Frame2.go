package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 7A6A5834-D185-4DBF-B63F-4A9BC43107D4
var IID_ICoreWebView2Frame2 = syscall.GUID{0x7A6A5834, 0xD185, 0x4DBF, 
	[8]byte{0xB6, 0x3F, 0x4A, 0x9B, 0xC4, 0x31, 0x07, 0xD4}}

type ICoreWebView2Frame2 struct {
	ICoreWebView2Frame
}

func NewICoreWebView2Frame2(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Frame2 {
	p := (*ICoreWebView2Frame2)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Frame2) IID() *syscall.GUID {
	return &IID_ICoreWebView2Frame2
}

func (this *ICoreWebView2Frame2) Add_NavigationStarting(eventHandler *ICoreWebView2FrameNavigationStartingEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[11]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Remove_NavigationStarting(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[12]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Add_ContentLoading(eventHandler *ICoreWebView2FrameContentLoadingEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[13]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Remove_ContentLoading(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[14]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Add_NavigationCompleted(eventHandler *ICoreWebView2FrameNavigationCompletedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[15]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Remove_NavigationCompleted(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[16]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Add_DOMContentLoaded(eventHandler *ICoreWebView2FrameDOMContentLoadedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[17]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Remove_DOMContentLoaded(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[18]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) ExecuteScript(javaScript string, handler *ICoreWebView2ExecuteScriptCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[19]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(javaScript)), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) PostWebMessageAsJson(webMessageAsJson string) com.Error {
	addr := (*this.LpVtbl)[20]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(webMessageAsJson)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) PostWebMessageAsString(webMessageAsString string) com.Error {
	addr := (*this.LpVtbl)[21]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(win32.StrToPointer(webMessageAsString)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Add_WebMessageReceived(handler *ICoreWebView2FrameWebMessageReceivedEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[22]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(handler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Frame2) Remove_WebMessageReceived(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[23]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

