package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// CCF1EF04-FD8E-4D5F-B2DE-0983E41B8C36
var IID_ICoreWebView2PrintToPdfCompletedHandler = syscall.GUID{0xCCF1EF04, 0xFD8E, 0x4D5F, 
	[8]byte{0xB2, 0xDE, 0x09, 0x83, 0xE4, 0x1B, 0x8C, 0x36}}

type ICoreWebView2PrintToPdfCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2PrintToPdfCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, isSuccessful int32) com.Error
}

type ICoreWebView2PrintToPdfCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2PrintToPdfCompletedHandlerInterface
}

func (this *ICoreWebView2PrintToPdfCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2PrintToPdfCompletedHandlerInterface)
}

func (this *ICoreWebView2PrintToPdfCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2PrintToPdfCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2PrintToPdfCompletedHandlerImpl) Invoke(errorCode com.Error, isSuccessful int32) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2PrintToPdfCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2PrintToPdfCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2PrintToPdfCompletedHandlerComObj) impl() ICoreWebView2PrintToPdfCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2PrintToPdfCompletedHandlerInterface)
}

func (this *ICoreWebView2PrintToPdfCompletedHandlerComObj) Invoke(errorCode com.Error, isSuccessful int32) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, isSuccessful))
}

var _pICoreWebView2PrintToPdfCompletedHandlerVtbl *ICoreWebView2PrintToPdfCompletedHandlerVtbl

func (this *ICoreWebView2PrintToPdfCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2PrintToPdfCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2PrintToPdfCompletedHandlerVtbl != nil {
		return _pICoreWebView2PrintToPdfCompletedHandlerVtbl
	}
	_pICoreWebView2PrintToPdfCompletedHandlerVtbl = &ICoreWebView2PrintToPdfCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2PrintToPdfCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2PrintToPdfCompletedHandlerVtbl
}

func (this *ICoreWebView2PrintToPdfCompletedHandlerComObj) ICoreWebView2PrintToPdfCompletedHandler() *ICoreWebView2PrintToPdfCompletedHandler{
	return (*ICoreWebView2PrintToPdfCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2PrintToPdfCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2PrintToPdfCompletedHandlerComObj(impl ICoreWebView2PrintToPdfCompletedHandlerInterface, scoped bool) *ICoreWebView2PrintToPdfCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2PrintToPdfCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2PrintToPdfCompletedHandler(impl ICoreWebView2PrintToPdfCompletedHandlerInterface) *ICoreWebView2PrintToPdfCompletedHandler {
	return NewICoreWebView2PrintToPdfCompletedHandlerComObj(impl, true).ICoreWebView2PrintToPdfCompletedHandler()}

//
type ICoreWebView2PrintToPdfCompletedHandlerByFuncImpl struct {
	ICoreWebView2PrintToPdfCompletedHandlerImpl
	handlerFunc func (errorCode com.Error, isSuccessful int32) com.Error
}
func (this *ICoreWebView2PrintToPdfCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, isSuccessful int32) com.Error{
	return this.handlerFunc(errorCode, isSuccessful)
}

func NewICoreWebView2PrintToPdfCompletedHandlerByFunc(handlerFunc func (errorCode com.Error, isSuccessful int32) com.Error, scoped bool) *ICoreWebView2PrintToPdfCompletedHandler {
	impl := &ICoreWebView2PrintToPdfCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2PrintToPdfCompletedHandlerComObj(impl, scoped).ICoreWebView2PrintToPdfCompletedHandler()
}

