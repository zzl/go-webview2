package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 38D9520D-340F-4D1E-A775-43FCE9753683
var IID_ICoreWebView2FrameDOMContentLoadedEventHandler = syscall.GUID{0x38D9520D, 0x340F, 0x4D1E, 
	[8]byte{0xA7, 0x75, 0x43, 0xFC, 0xE9, 0x75, 0x36, 0x83}}

type ICoreWebView2FrameDOMContentLoadedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FrameDOMContentLoadedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error
}

type ICoreWebView2FrameDOMContentLoadedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FrameDOMContentLoadedEventHandlerInterface
}

func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FrameDOMContentLoadedEventHandlerInterface)
}

func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FrameDOMContentLoadedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FrameDOMContentLoadedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerComObj) impl() ICoreWebView2FrameDOMContentLoadedEventHandlerInterface {
	return this.Impl().(ICoreWebView2FrameDOMContentLoadedEventHandlerInterface)
}

func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerComObj) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FrameDOMContentLoadedEventHandlerVtbl *ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl

func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2FrameDOMContentLoadedEventHandlerVtbl != nil {
		return _pICoreWebView2FrameDOMContentLoadedEventHandlerVtbl
	}
	_pICoreWebView2FrameDOMContentLoadedEventHandlerVtbl = &ICoreWebView2FrameDOMContentLoadedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2FrameDOMContentLoadedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FrameDOMContentLoadedEventHandlerVtbl
}

func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerComObj) ICoreWebView2FrameDOMContentLoadedEventHandler() *ICoreWebView2FrameDOMContentLoadedEventHandler{
	return (*ICoreWebView2FrameDOMContentLoadedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FrameDOMContentLoadedEventHandlerComObj(impl ICoreWebView2FrameDOMContentLoadedEventHandlerInterface, scoped bool) *ICoreWebView2FrameDOMContentLoadedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FrameDOMContentLoadedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FrameDOMContentLoadedEventHandler(impl ICoreWebView2FrameDOMContentLoadedEventHandlerInterface) *ICoreWebView2FrameDOMContentLoadedEventHandler {
	return NewICoreWebView2FrameDOMContentLoadedEventHandlerComObj(impl, true).ICoreWebView2FrameDOMContentLoadedEventHandler()}

//
type ICoreWebView2FrameDOMContentLoadedEventHandlerByFuncImpl struct {
	ICoreWebView2FrameDOMContentLoadedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error
}
func (this *ICoreWebView2FrameDOMContentLoadedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FrameDOMContentLoadedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2DOMContentLoadedEventArgs) com.Error, scoped bool) *ICoreWebView2FrameDOMContentLoadedEventHandler {
	impl := &ICoreWebView2FrameDOMContentLoadedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FrameDOMContentLoadedEventHandlerComObj(impl, scoped).ICoreWebView2FrameDOMContentLoadedEventHandler()
}

