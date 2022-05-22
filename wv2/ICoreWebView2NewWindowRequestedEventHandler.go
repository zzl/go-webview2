package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// D4C185FE-C81C-4989-97AF-2D3FA7AB5651
var IID_ICoreWebView2NewWindowRequestedEventHandler = syscall.GUID{0xD4C185FE, 0xC81C, 0x4989, 
	[8]byte{0x97, 0xAF, 0x2D, 0x3F, 0xA7, 0xAB, 0x56, 0x51}}

type ICoreWebView2NewWindowRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2NewWindowRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) com.Error
}

type ICoreWebView2NewWindowRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2NewWindowRequestedEventHandlerInterface
}

func (this *ICoreWebView2NewWindowRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2NewWindowRequestedEventHandlerInterface)
}

func (this *ICoreWebView2NewWindowRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2NewWindowRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2NewWindowRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2NewWindowRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2NewWindowRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2NewWindowRequestedEventHandlerComObj) impl() ICoreWebView2NewWindowRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2NewWindowRequestedEventHandlerInterface)
}

func (this *ICoreWebView2NewWindowRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2NewWindowRequestedEventHandlerVtbl *ICoreWebView2NewWindowRequestedEventHandlerVtbl

func (this *ICoreWebView2NewWindowRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2NewWindowRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2NewWindowRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2NewWindowRequestedEventHandlerVtbl
	}
	_pICoreWebView2NewWindowRequestedEventHandlerVtbl = &ICoreWebView2NewWindowRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2NewWindowRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2NewWindowRequestedEventHandlerVtbl
}

func (this *ICoreWebView2NewWindowRequestedEventHandlerComObj) ICoreWebView2NewWindowRequestedEventHandler() *ICoreWebView2NewWindowRequestedEventHandler{
	return (*ICoreWebView2NewWindowRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2NewWindowRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2NewWindowRequestedEventHandlerComObj(impl ICoreWebView2NewWindowRequestedEventHandlerInterface, scoped bool) *ICoreWebView2NewWindowRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2NewWindowRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2NewWindowRequestedEventHandler(impl ICoreWebView2NewWindowRequestedEventHandlerInterface) *ICoreWebView2NewWindowRequestedEventHandler {
	return NewICoreWebView2NewWindowRequestedEventHandlerComObj(impl, true).ICoreWebView2NewWindowRequestedEventHandler()}

//
type ICoreWebView2NewWindowRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2NewWindowRequestedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) com.Error
}
func (this *ICoreWebView2NewWindowRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2NewWindowRequestedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2NewWindowRequestedEventArgs) com.Error, scoped bool) *ICoreWebView2NewWindowRequestedEventHandler {
	impl := &ICoreWebView2NewWindowRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2NewWindowRequestedEventHandlerComObj(impl, scoped).ICoreWebView2NewWindowRequestedEventHandler()
}

