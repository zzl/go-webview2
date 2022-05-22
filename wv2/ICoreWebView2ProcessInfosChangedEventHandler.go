package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// F4AF0C39-44B9-40E9-8B11-0484CFB9E0A1
var IID_ICoreWebView2ProcessInfosChangedEventHandler = syscall.GUID{0xF4AF0C39, 0x44B9, 0x40E9, 
	[8]byte{0x8B, 0x11, 0x04, 0x84, 0xCF, 0xB9, 0xE0, 0xA1}}

type ICoreWebView2ProcessInfosChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2ProcessInfosChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Environment, args *com.UnknownClass) com.Error
}

type ICoreWebView2ProcessInfosChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ProcessInfosChangedEventHandlerInterface
}

func (this *ICoreWebView2ProcessInfosChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ProcessInfosChangedEventHandlerInterface)
}

func (this *ICoreWebView2ProcessInfosChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ProcessInfosChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ProcessInfosChangedEventHandlerImpl) Invoke(sender *ICoreWebView2Environment, args *com.UnknownClass) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2ProcessInfosChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ProcessInfosChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ProcessInfosChangedEventHandlerComObj) impl() ICoreWebView2ProcessInfosChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2ProcessInfosChangedEventHandlerInterface)
}

func (this *ICoreWebView2ProcessInfosChangedEventHandlerComObj) Invoke(sender *ICoreWebView2Environment, args *com.UnknownClass) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2ProcessInfosChangedEventHandlerVtbl *ICoreWebView2ProcessInfosChangedEventHandlerVtbl

func (this *ICoreWebView2ProcessInfosChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ProcessInfosChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2ProcessInfosChangedEventHandlerVtbl != nil {
		return _pICoreWebView2ProcessInfosChangedEventHandlerVtbl
	}
	_pICoreWebView2ProcessInfosChangedEventHandlerVtbl = &ICoreWebView2ProcessInfosChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2ProcessInfosChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2ProcessInfosChangedEventHandlerVtbl
}

func (this *ICoreWebView2ProcessInfosChangedEventHandlerComObj) ICoreWebView2ProcessInfosChangedEventHandler() *ICoreWebView2ProcessInfosChangedEventHandler{
	return (*ICoreWebView2ProcessInfosChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ProcessInfosChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ProcessInfosChangedEventHandlerComObj(impl ICoreWebView2ProcessInfosChangedEventHandlerInterface, scoped bool) *ICoreWebView2ProcessInfosChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ProcessInfosChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ProcessInfosChangedEventHandler(impl ICoreWebView2ProcessInfosChangedEventHandlerInterface) *ICoreWebView2ProcessInfosChangedEventHandler {
	return NewICoreWebView2ProcessInfosChangedEventHandlerComObj(impl, true).ICoreWebView2ProcessInfosChangedEventHandler()}

//
type ICoreWebView2ProcessInfosChangedEventHandlerByFuncImpl struct {
	ICoreWebView2ProcessInfosChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Environment, args *com.UnknownClass) com.Error
}
func (this *ICoreWebView2ProcessInfosChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Environment, args *com.UnknownClass) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2ProcessInfosChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Environment, args *com.UnknownClass) com.Error, scoped bool) *ICoreWebView2ProcessInfosChangedEventHandler {
	impl := &ICoreWebView2ProcessInfosChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ProcessInfosChangedEventHandlerComObj(impl, scoped).ICoreWebView2ProcessInfosChangedEventHandler()
}

