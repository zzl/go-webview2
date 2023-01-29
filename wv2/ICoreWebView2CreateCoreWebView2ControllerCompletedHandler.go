package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 6C4819F3-C9B7-4260-8127-C9F5BDE7F68C
var IID_ICoreWebView2CreateCoreWebView2ControllerCompletedHandler = syscall.GUID{0x6C4819F3, 0xC9B7, 0x4260,
	[8]byte{0x81, 0x27, 0xC9, 0xF5, 0xBD, 0xE7, 0xF6, 0x8C}}

type ICoreWebView2CreateCoreWebView2ControllerCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, createdController *ICoreWebView2Controller) com.Error
}

type ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInterface
}

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInterface)
}

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2CreateCoreWebView2ControllerCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl) Invoke(errorCode com.Error, createdController *ICoreWebView2Controller) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj) impl() ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInterface)
}

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj) Invoke(errorCode com.Error, createdController *ICoreWebView2Controller) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, createdController))
}

var _pICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl != nil {
		return _pICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl
	}
	_pICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl = &ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2CreateCoreWebView2ControllerCompletedHandlerVtbl
}

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj) ICoreWebView2CreateCoreWebView2ControllerCompletedHandler() *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler {
	return (*ICoreWebView2CreateCoreWebView2ControllerCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj(impl ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInterface, scoped bool) *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2CreateCoreWebView2ControllerCompletedHandler(impl ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerInterface) *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler {
	return NewICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj(impl, true).ICoreWebView2CreateCoreWebView2ControllerCompletedHandler()
}

//
type ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerByFuncImpl struct {
	ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerImpl
	handlerFunc func(errorCode com.Error, createdController *ICoreWebView2Controller) com.Error
}

func (this *ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, createdController *ICoreWebView2Controller) com.Error {
	return this.handlerFunc(errorCode, createdController)
}

func NewICoreWebView2CreateCoreWebView2ControllerCompletedHandlerByFunc(handlerFunc func(errorCode com.Error, createdController *ICoreWebView2Controller) com.Error, scoped bool) *ICoreWebView2CreateCoreWebView2ControllerCompletedHandler {
	impl := &ICoreWebView2CreateCoreWebView2ControllerCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2CreateCoreWebView2ControllerCompletedHandlerComObj(impl, scoped).ICoreWebView2CreateCoreWebView2ControllerCompletedHandler()
}

