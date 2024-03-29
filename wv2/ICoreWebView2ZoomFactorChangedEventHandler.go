package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// B52D71D6-C4DF-4543-A90C-64A3E60F38CB
var IID_ICoreWebView2ZoomFactorChangedEventHandler = syscall.GUID{0xB52D71D6, 0xC4DF, 0x4543,
	[8]byte{0xA9, 0x0C, 0x64, 0xA3, 0xE6, 0x0F, 0x38, 0xCB}}

type ICoreWebView2ZoomFactorChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2ZoomFactorChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Controller, args *win32.IUnknown) com.Error
}

type ICoreWebView2ZoomFactorChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ZoomFactorChangedEventHandlerInterface
}

func (this *ICoreWebView2ZoomFactorChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ZoomFactorChangedEventHandlerInterface)
}

func (this *ICoreWebView2ZoomFactorChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ZoomFactorChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ZoomFactorChangedEventHandlerImpl) Invoke(sender *ICoreWebView2Controller, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2ZoomFactorChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ZoomFactorChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ZoomFactorChangedEventHandlerComObj) impl() ICoreWebView2ZoomFactorChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2ZoomFactorChangedEventHandlerInterface)
}

func (this *ICoreWebView2ZoomFactorChangedEventHandlerComObj) Invoke(sender *ICoreWebView2Controller, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2ZoomFactorChangedEventHandlerVtbl *ICoreWebView2ZoomFactorChangedEventHandlerVtbl

func (this *ICoreWebView2ZoomFactorChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ZoomFactorChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2ZoomFactorChangedEventHandlerVtbl != nil {
		return _pICoreWebView2ZoomFactorChangedEventHandlerVtbl
	}
	_pICoreWebView2ZoomFactorChangedEventHandlerVtbl = &ICoreWebView2ZoomFactorChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2ZoomFactorChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2ZoomFactorChangedEventHandlerVtbl
}

func (this *ICoreWebView2ZoomFactorChangedEventHandlerComObj) ICoreWebView2ZoomFactorChangedEventHandler() *ICoreWebView2ZoomFactorChangedEventHandler {
	return (*ICoreWebView2ZoomFactorChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ZoomFactorChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ZoomFactorChangedEventHandlerComObj(impl ICoreWebView2ZoomFactorChangedEventHandlerInterface, scoped bool) *ICoreWebView2ZoomFactorChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ZoomFactorChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ZoomFactorChangedEventHandler(impl ICoreWebView2ZoomFactorChangedEventHandlerInterface) *ICoreWebView2ZoomFactorChangedEventHandler {
	return NewICoreWebView2ZoomFactorChangedEventHandlerComObj(impl, true).ICoreWebView2ZoomFactorChangedEventHandler()
}

//
type ICoreWebView2ZoomFactorChangedEventHandlerByFuncImpl struct {
	ICoreWebView2ZoomFactorChangedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2Controller, args *win32.IUnknown) com.Error
}

func (this *ICoreWebView2ZoomFactorChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Controller, args *win32.IUnknown) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2ZoomFactorChangedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2Controller, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2ZoomFactorChangedEventHandler {
	impl := &ICoreWebView2ZoomFactorChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ZoomFactorChangedEventHandlerComObj(impl, scoped).ICoreWebView2ZoomFactorChangedEventHandler()
}

