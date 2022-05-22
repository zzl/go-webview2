package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 364471E7-F2BE-4910-BDBA-D72077D51C4B
var IID_ICoreWebView2ContentLoadingEventHandler = syscall.GUID{0x364471E7, 0xF2BE, 0x4910, 
	[8]byte{0xBD, 0xBA, 0xD7, 0x20, 0x77, 0xD5, 0x1C, 0x4B}}

type ICoreWebView2ContentLoadingEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2ContentLoadingEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) com.Error
}

type ICoreWebView2ContentLoadingEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ContentLoadingEventHandlerInterface
}

func (this *ICoreWebView2ContentLoadingEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ContentLoadingEventHandlerInterface)
}

func (this *ICoreWebView2ContentLoadingEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ContentLoadingEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ContentLoadingEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2ContentLoadingEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ContentLoadingEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ContentLoadingEventHandlerComObj) impl() ICoreWebView2ContentLoadingEventHandlerInterface {
	return this.Impl().(ICoreWebView2ContentLoadingEventHandlerInterface)
}

func (this *ICoreWebView2ContentLoadingEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2ContentLoadingEventHandlerVtbl *ICoreWebView2ContentLoadingEventHandlerVtbl

func (this *ICoreWebView2ContentLoadingEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ContentLoadingEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2ContentLoadingEventHandlerVtbl != nil {
		return _pICoreWebView2ContentLoadingEventHandlerVtbl
	}
	_pICoreWebView2ContentLoadingEventHandlerVtbl = &ICoreWebView2ContentLoadingEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2ContentLoadingEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2ContentLoadingEventHandlerVtbl
}

func (this *ICoreWebView2ContentLoadingEventHandlerComObj) ICoreWebView2ContentLoadingEventHandler() *ICoreWebView2ContentLoadingEventHandler{
	return (*ICoreWebView2ContentLoadingEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ContentLoadingEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ContentLoadingEventHandlerComObj(impl ICoreWebView2ContentLoadingEventHandlerInterface, scoped bool) *ICoreWebView2ContentLoadingEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ContentLoadingEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ContentLoadingEventHandler(impl ICoreWebView2ContentLoadingEventHandlerInterface) *ICoreWebView2ContentLoadingEventHandler {
	return NewICoreWebView2ContentLoadingEventHandlerComObj(impl, true).ICoreWebView2ContentLoadingEventHandler()}

//
type ICoreWebView2ContentLoadingEventHandlerByFuncImpl struct {
	ICoreWebView2ContentLoadingEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) com.Error
}
func (this *ICoreWebView2ContentLoadingEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2ContentLoadingEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2ContentLoadingEventArgs) com.Error, scoped bool) *ICoreWebView2ContentLoadingEventHandler {
	impl := &ICoreWebView2ContentLoadingEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ContentLoadingEventHandlerComObj(impl, scoped).ICoreWebView2ContentLoadingEventHandler()
}

