package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 5A4F5069-5C15-47C3-8646-F4DE1C116670
var IID_ICoreWebView2GetCookiesCompletedHandler = syscall.GUID{0x5A4F5069, 0x5C15, 0x47C3,
	[8]byte{0x86, 0x46, 0xF4, 0xDE, 0x1C, 0x11, 0x66, 0x70}}

type ICoreWebView2GetCookiesCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2GetCookiesCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(result com.Error, cookieList *ICoreWebView2CookieList) com.Error
}

type ICoreWebView2GetCookiesCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2GetCookiesCompletedHandlerInterface
}

func (this *ICoreWebView2GetCookiesCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2GetCookiesCompletedHandlerInterface)
}

func (this *ICoreWebView2GetCookiesCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2GetCookiesCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2GetCookiesCompletedHandlerImpl) Invoke(result com.Error, cookieList *ICoreWebView2CookieList) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2GetCookiesCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2GetCookiesCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2GetCookiesCompletedHandlerComObj) impl() ICoreWebView2GetCookiesCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2GetCookiesCompletedHandlerInterface)
}

func (this *ICoreWebView2GetCookiesCompletedHandlerComObj) Invoke(result com.Error, cookieList *ICoreWebView2CookieList) uintptr {
	return (uintptr)(this.impl().Invoke(result, cookieList))
}

var _pICoreWebView2GetCookiesCompletedHandlerVtbl *ICoreWebView2GetCookiesCompletedHandlerVtbl

func (this *ICoreWebView2GetCookiesCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2GetCookiesCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2GetCookiesCompletedHandlerVtbl != nil {
		return _pICoreWebView2GetCookiesCompletedHandlerVtbl
	}
	_pICoreWebView2GetCookiesCompletedHandlerVtbl = &ICoreWebView2GetCookiesCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2GetCookiesCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2GetCookiesCompletedHandlerVtbl
}

func (this *ICoreWebView2GetCookiesCompletedHandlerComObj) ICoreWebView2GetCookiesCompletedHandler() *ICoreWebView2GetCookiesCompletedHandler {
	return (*ICoreWebView2GetCookiesCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2GetCookiesCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2GetCookiesCompletedHandlerComObj(impl ICoreWebView2GetCookiesCompletedHandlerInterface, scoped bool) *ICoreWebView2GetCookiesCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2GetCookiesCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2GetCookiesCompletedHandler(impl ICoreWebView2GetCookiesCompletedHandlerInterface) *ICoreWebView2GetCookiesCompletedHandler {
	return NewICoreWebView2GetCookiesCompletedHandlerComObj(impl, true).ICoreWebView2GetCookiesCompletedHandler()
}

//
type ICoreWebView2GetCookiesCompletedHandlerByFuncImpl struct {
	ICoreWebView2GetCookiesCompletedHandlerImpl
	handlerFunc func(result com.Error, cookieList *ICoreWebView2CookieList) com.Error
}

func (this *ICoreWebView2GetCookiesCompletedHandlerByFuncImpl) Invoke(result com.Error, cookieList *ICoreWebView2CookieList) com.Error {
	return this.handlerFunc(result, cookieList)
}

func NewICoreWebView2GetCookiesCompletedHandlerByFunc(handlerFunc func(result com.Error, cookieList *ICoreWebView2CookieList) com.Error, scoped bool) *ICoreWebView2GetCookiesCompletedHandler {
	impl := &ICoreWebView2GetCookiesCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2GetCookiesCompletedHandlerComObj(impl, scoped).ICoreWebView2GetCookiesCompletedHandler()
}

