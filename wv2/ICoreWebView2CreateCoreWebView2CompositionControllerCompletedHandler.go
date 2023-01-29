package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 02FAB84B-1428-4FB7-AD45-1B2E64736184
var IID_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler = syscall.GUID{0x02FAB84B, 0x1428, 0x4FB7,
	[8]byte{0xAD, 0x45, 0x1B, 0x2E, 0x64, 0x73, 0x61, 0x84}}

type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, webView *ICoreWebView2CompositionController) com.Error
}

type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInterface
}

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInterface)
}

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl) Invoke(errorCode com.Error, webView *ICoreWebView2CompositionController) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj) impl() ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInterface)
}

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj) Invoke(errorCode com.Error, webView *ICoreWebView2CompositionController) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, webView))
}

var _pICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl != nil {
		return _pICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl
	}
	_pICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl = &ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerVtbl
}

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj) ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler() *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler {
	return (*ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj(impl ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInterface, scoped bool) *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler(impl ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerInterface) *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler {
	return NewICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj(impl, true).ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler()
}

//
type ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerByFuncImpl struct {
	ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerImpl
	handlerFunc func(errorCode com.Error, webView *ICoreWebView2CompositionController) com.Error
}

func (this *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, webView *ICoreWebView2CompositionController) com.Error {
	return this.handlerFunc(errorCode, webView)
}

func NewICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerByFunc(handlerFunc func(errorCode com.Error, webView *ICoreWebView2CompositionController) com.Error, scoped bool) *ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler {
	impl := &ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandlerComObj(impl, scoped).ICoreWebView2CreateCoreWebView2CompositionControllerCompletedHandler()
}

