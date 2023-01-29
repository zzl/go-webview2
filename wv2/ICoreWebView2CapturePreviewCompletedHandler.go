package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 697E05E9-3D8F-45FA-96F4-8FFE1EDEDAF5
var IID_ICoreWebView2CapturePreviewCompletedHandler = syscall.GUID{0x697E05E9, 0x3D8F, 0x45FA,
	[8]byte{0x96, 0xF4, 0x8F, 0xFE, 0x1E, 0xDE, 0xDA, 0xF5}}

type ICoreWebView2CapturePreviewCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2CapturePreviewCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error) com.Error
}

type ICoreWebView2CapturePreviewCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2CapturePreviewCompletedHandlerInterface
}

func (this *ICoreWebView2CapturePreviewCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2CapturePreviewCompletedHandlerInterface)
}

func (this *ICoreWebView2CapturePreviewCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2CapturePreviewCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2CapturePreviewCompletedHandlerImpl) Invoke(errorCode com.Error) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2CapturePreviewCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2CapturePreviewCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2CapturePreviewCompletedHandlerComObj) impl() ICoreWebView2CapturePreviewCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2CapturePreviewCompletedHandlerInterface)
}

func (this *ICoreWebView2CapturePreviewCompletedHandlerComObj) Invoke(errorCode com.Error) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode))
}

var _pICoreWebView2CapturePreviewCompletedHandlerVtbl *ICoreWebView2CapturePreviewCompletedHandlerVtbl

func (this *ICoreWebView2CapturePreviewCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2CapturePreviewCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2CapturePreviewCompletedHandlerVtbl != nil {
		return _pICoreWebView2CapturePreviewCompletedHandlerVtbl
	}
	_pICoreWebView2CapturePreviewCompletedHandlerVtbl = &ICoreWebView2CapturePreviewCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2CapturePreviewCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2CapturePreviewCompletedHandlerVtbl
}

func (this *ICoreWebView2CapturePreviewCompletedHandlerComObj) ICoreWebView2CapturePreviewCompletedHandler() *ICoreWebView2CapturePreviewCompletedHandler {
	return (*ICoreWebView2CapturePreviewCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2CapturePreviewCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2CapturePreviewCompletedHandlerComObj(impl ICoreWebView2CapturePreviewCompletedHandlerInterface, scoped bool) *ICoreWebView2CapturePreviewCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2CapturePreviewCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2CapturePreviewCompletedHandler(impl ICoreWebView2CapturePreviewCompletedHandlerInterface) *ICoreWebView2CapturePreviewCompletedHandler {
	return NewICoreWebView2CapturePreviewCompletedHandlerComObj(impl, true).ICoreWebView2CapturePreviewCompletedHandler()
}

type ICoreWebView2CapturePreviewCompletedHandlerByFuncImpl struct {
	ICoreWebView2CapturePreviewCompletedHandlerImpl
	handlerFunc func(errorCode com.Error) com.Error
}

func (this *ICoreWebView2CapturePreviewCompletedHandlerByFuncImpl) Invoke(errorCode com.Error) com.Error {
	return this.handlerFunc(errorCode)
}

func NewICoreWebView2CapturePreviewCompletedHandlerByFunc(handlerFunc func(errorCode com.Error) com.Error, scoped bool) *ICoreWebView2CapturePreviewCompletedHandler {
	impl := &ICoreWebView2CapturePreviewCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2CapturePreviewCompletedHandlerComObj(impl, scoped).ICoreWebView2CapturePreviewCompletedHandler()
}
