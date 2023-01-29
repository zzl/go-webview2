package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 79E0AEA4-990B-42D9-AA1D-0FCC2E5BC7F1
var IID_ICoreWebView2ProcessFailedEventHandler = syscall.GUID{0x79E0AEA4, 0x990B, 0x42D9,
	[8]byte{0xAA, 0x1D, 0x0F, 0xCC, 0x2E, 0x5B, 0xC7, 0xF1}}

type ICoreWebView2ProcessFailedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2ProcessFailedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) com.Error
}

type ICoreWebView2ProcessFailedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ProcessFailedEventHandlerInterface
}

func (this *ICoreWebView2ProcessFailedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ProcessFailedEventHandlerInterface)
}

func (this *ICoreWebView2ProcessFailedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ProcessFailedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ProcessFailedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2ProcessFailedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ProcessFailedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ProcessFailedEventHandlerComObj) impl() ICoreWebView2ProcessFailedEventHandlerInterface {
	return this.Impl().(ICoreWebView2ProcessFailedEventHandlerInterface)
}

func (this *ICoreWebView2ProcessFailedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2ProcessFailedEventHandlerVtbl *ICoreWebView2ProcessFailedEventHandlerVtbl

func (this *ICoreWebView2ProcessFailedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ProcessFailedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2ProcessFailedEventHandlerVtbl != nil {
		return _pICoreWebView2ProcessFailedEventHandlerVtbl
	}
	_pICoreWebView2ProcessFailedEventHandlerVtbl = &ICoreWebView2ProcessFailedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2ProcessFailedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2ProcessFailedEventHandlerVtbl
}

func (this *ICoreWebView2ProcessFailedEventHandlerComObj) ICoreWebView2ProcessFailedEventHandler() *ICoreWebView2ProcessFailedEventHandler {
	return (*ICoreWebView2ProcessFailedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ProcessFailedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ProcessFailedEventHandlerComObj(impl ICoreWebView2ProcessFailedEventHandlerInterface, scoped bool) *ICoreWebView2ProcessFailedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ProcessFailedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ProcessFailedEventHandler(impl ICoreWebView2ProcessFailedEventHandlerInterface) *ICoreWebView2ProcessFailedEventHandler {
	return NewICoreWebView2ProcessFailedEventHandlerComObj(impl, true).ICoreWebView2ProcessFailedEventHandler()
}

type ICoreWebView2ProcessFailedEventHandlerByFuncImpl struct {
	ICoreWebView2ProcessFailedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) com.Error
}

func (this *ICoreWebView2ProcessFailedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2ProcessFailedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2ProcessFailedEventArgs) com.Error, scoped bool) *ICoreWebView2ProcessFailedEventHandler {
	impl := &ICoreWebView2ProcessFailedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ProcessFailedEventHandlerComObj(impl, scoped).ICoreWebView2ProcessFailedEventHandler()
}
