package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// B96D755E-0319-4E92-A296-23436F46A1FC
var IID_ICoreWebView2Environment = syscall.GUID{0xB96D755E, 0x0319, 0x4E92, 
	[8]byte{0xA2, 0x96, 0x23, 0x43, 0x6F, 0x46, 0xA1, 0xFC}}

type ICoreWebView2Environment struct {
	win32.IUnknown
}

func NewICoreWebView2Environment(pUnk *win32.IUnknown, addRef bool, scoped bool) *ICoreWebView2Environment {
	p := (*ICoreWebView2Environment)(unsafe.Pointer(pUnk))
	if addRef {
		pUnk.AddRef()
	}
	if scoped {
		com.AddToScope(p)
	}
	return p
}

func (this *ICoreWebView2Environment) IID() *syscall.GUID {
	return &IID_ICoreWebView2Environment
}

func (this *ICoreWebView2Environment) CreateCoreWebView2Controller(parentWindow win32.HWND, handler *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler) com.Error {
	addr := (*this.LpVtbl)[3]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(parentWindow), uintptr(unsafe.Pointer(handler)))
	return com.Error(ret)
}

func (this *ICoreWebView2Environment) CreateWebResourceResponse(content *win32.IStream, statusCode int32, reasonPhrase string, headers string, response **ICoreWebView2WebResourceResponse) com.Error {
	addr := (*this.LpVtbl)[4]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(content)), uintptr(statusCode), uintptr(win32.StrToPointer(reasonPhrase)), uintptr(win32.StrToPointer(headers)), uintptr(unsafe.Pointer(response)))
	if com.CurrentScope != nil {
		com.CurrentScope.Add(unsafe.Pointer(&(*response).IUnknown))
	}
	return com.Error(ret)
}

func (this *ICoreWebView2Environment) GetBrowserVersionString(versionInfo *win32.PWSTR) com.Error {
	addr := (*this.LpVtbl)[5]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(versionInfo)))
	return com.Error(ret)
}

func (this *ICoreWebView2Environment) Add_NewBrowserVersionAvailable(eventHandler *ICoreWebView2NewBrowserVersionAvailableEventHandler, token *EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[6]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), uintptr(unsafe.Pointer(eventHandler)), uintptr(unsafe.Pointer(token)))
	return com.Error(ret)
}

func (this *ICoreWebView2Environment) Remove_NewBrowserVersionAvailable(token EventRegistrationToken) com.Error {
	addr := (*this.LpVtbl)[7]
	ret, _, _ := syscall.SyscallN(addr, uintptr(unsafe.Pointer(this)), *(*uintptr)(unsafe.Pointer(&token)))
	return com.Error(ret)
}

