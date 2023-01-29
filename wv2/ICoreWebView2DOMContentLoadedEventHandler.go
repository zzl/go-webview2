package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 4BAC7E9C-199E-49ED-87ED-249303ACF019
var IID_ICoreWebView2DOMContentLoadedEventHandler = syscall.GUID{0x4BAC7E9C, 0x199E, 0x49ED,
	[8]byte{0x87, 0xED, 0x24, 0x93, 0x03, 0xAC, 0xF0, 0x19}}

type ICoreWebView2DOMContentLoadedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2DOMContentLoadedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error
}

type ICoreWebView2DOMContentLoadedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2DOMContentLoadedEventHandlerInterface
}

func (this *ICoreWebView2DOMContentLoadedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2DOMContentLoadedEventHandlerInterface)
}

func (this *ICoreWebView2DOMContentLoadedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2DOMContentLoadedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2DOMContentLoadedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2DOMContentLoadedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2DOMContentLoadedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2DOMContentLoadedEventHandlerComObj) impl() ICoreWebView2DOMContentLoadedEventHandlerInterface {
	return this.Impl().(ICoreWebView2DOMContentLoadedEventHandlerInterface)
}

func (this *ICoreWebView2DOMContentLoadedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2DOMContentLoadedEventHandlerVtbl *ICoreWebView2DOMContentLoadedEventHandlerVtbl

func (this *ICoreWebView2DOMContentLoadedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2DOMContentLoadedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2DOMContentLoadedEventHandlerVtbl != nil {
		return _pICoreWebView2DOMContentLoadedEventHandlerVtbl
	}
	_pICoreWebView2DOMContentLoadedEventHandlerVtbl = &ICoreWebView2DOMContentLoadedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2DOMContentLoadedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2DOMContentLoadedEventHandlerVtbl
}

func (this *ICoreWebView2DOMContentLoadedEventHandlerComObj) ICoreWebView2DOMContentLoadedEventHandler() *ICoreWebView2DOMContentLoadedEventHandler {
	return (*ICoreWebView2DOMContentLoadedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2DOMContentLoadedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2DOMContentLoadedEventHandlerComObj(impl ICoreWebView2DOMContentLoadedEventHandlerInterface, scoped bool) *ICoreWebView2DOMContentLoadedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2DOMContentLoadedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2DOMContentLoadedEventHandler(impl ICoreWebView2DOMContentLoadedEventHandlerInterface) *ICoreWebView2DOMContentLoadedEventHandler {
	return NewICoreWebView2DOMContentLoadedEventHandlerComObj(impl, true).ICoreWebView2DOMContentLoadedEventHandler()
}

//
type ICoreWebView2DOMContentLoadedEventHandlerByFuncImpl struct {
	ICoreWebView2DOMContentLoadedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error
}

func (this *ICoreWebView2DOMContentLoadedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2DOMContentLoadedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error, scoped bool) *ICoreWebView2DOMContentLoadedEventHandler {
	impl := &ICoreWebView2DOMContentLoadedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2DOMContentLoadedEventHandlerComObj(impl, scoped).ICoreWebView2DOMContentLoadedEventHandler()
}

