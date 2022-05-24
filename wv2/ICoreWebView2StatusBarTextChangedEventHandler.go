package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// A5E3B0D0-10DF-4156-BFAD-3B43867ACAC6
var IID_ICoreWebView2StatusBarTextChangedEventHandler = syscall.GUID{0xA5E3B0D0, 0x10DF, 0x4156, 
	[8]byte{0xBF, 0xAD, 0x3B, 0x43, 0x86, 0x7A, 0xCA, 0xC6}}

type ICoreWebView2StatusBarTextChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2StatusBarTextChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

type ICoreWebView2StatusBarTextChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2StatusBarTextChangedEventHandlerInterface
}

func (this *ICoreWebView2StatusBarTextChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2StatusBarTextChangedEventHandlerInterface)
}

func (this *ICoreWebView2StatusBarTextChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2StatusBarTextChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2StatusBarTextChangedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2StatusBarTextChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2StatusBarTextChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2StatusBarTextChangedEventHandlerComObj) impl() ICoreWebView2StatusBarTextChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2StatusBarTextChangedEventHandlerInterface)
}

func (this *ICoreWebView2StatusBarTextChangedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2StatusBarTextChangedEventHandlerVtbl *ICoreWebView2StatusBarTextChangedEventHandlerVtbl

func (this *ICoreWebView2StatusBarTextChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2StatusBarTextChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2StatusBarTextChangedEventHandlerVtbl != nil {
		return _pICoreWebView2StatusBarTextChangedEventHandlerVtbl
	}
	_pICoreWebView2StatusBarTextChangedEventHandlerVtbl = &ICoreWebView2StatusBarTextChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2StatusBarTextChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2StatusBarTextChangedEventHandlerVtbl
}

func (this *ICoreWebView2StatusBarTextChangedEventHandlerComObj) ICoreWebView2StatusBarTextChangedEventHandler() *ICoreWebView2StatusBarTextChangedEventHandler{
	return (*ICoreWebView2StatusBarTextChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2StatusBarTextChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2StatusBarTextChangedEventHandlerComObj(impl ICoreWebView2StatusBarTextChangedEventHandlerInterface, scoped bool) *ICoreWebView2StatusBarTextChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2StatusBarTextChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2StatusBarTextChangedEventHandler(impl ICoreWebView2StatusBarTextChangedEventHandlerInterface) *ICoreWebView2StatusBarTextChangedEventHandler {
	return NewICoreWebView2StatusBarTextChangedEventHandlerComObj(impl, true).ICoreWebView2StatusBarTextChangedEventHandler()}

//
type ICoreWebView2StatusBarTextChangedEventHandlerByFuncImpl struct {
	ICoreWebView2StatusBarTextChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error
}
func (this *ICoreWebView2StatusBarTextChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2StatusBarTextChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2StatusBarTextChangedEventHandler {
	impl := &ICoreWebView2StatusBarTextChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2StatusBarTextChangedEventHandlerComObj(impl, scoped).ICoreWebView2StatusBarTextChangedEventHandler()
}

