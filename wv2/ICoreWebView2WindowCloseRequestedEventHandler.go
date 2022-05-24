package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 5C19E9E0-092F-486B-AFFA-CA8231913039
var IID_ICoreWebView2WindowCloseRequestedEventHandler = syscall.GUID{0x5C19E9E0, 0x092F, 0x486B, 
	[8]byte{0xAF, 0xFA, 0xCA, 0x82, 0x31, 0x91, 0x30, 0x39}}

type ICoreWebView2WindowCloseRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2WindowCloseRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

type ICoreWebView2WindowCloseRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2WindowCloseRequestedEventHandlerInterface
}

func (this *ICoreWebView2WindowCloseRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2WindowCloseRequestedEventHandlerInterface)
}

func (this *ICoreWebView2WindowCloseRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2WindowCloseRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2WindowCloseRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2WindowCloseRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2WindowCloseRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2WindowCloseRequestedEventHandlerComObj) impl() ICoreWebView2WindowCloseRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2WindowCloseRequestedEventHandlerInterface)
}

func (this *ICoreWebView2WindowCloseRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2WindowCloseRequestedEventHandlerVtbl *ICoreWebView2WindowCloseRequestedEventHandlerVtbl

func (this *ICoreWebView2WindowCloseRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2WindowCloseRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2WindowCloseRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2WindowCloseRequestedEventHandlerVtbl
	}
	_pICoreWebView2WindowCloseRequestedEventHandlerVtbl = &ICoreWebView2WindowCloseRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2WindowCloseRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2WindowCloseRequestedEventHandlerVtbl
}

func (this *ICoreWebView2WindowCloseRequestedEventHandlerComObj) ICoreWebView2WindowCloseRequestedEventHandler() *ICoreWebView2WindowCloseRequestedEventHandler{
	return (*ICoreWebView2WindowCloseRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2WindowCloseRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2WindowCloseRequestedEventHandlerComObj(impl ICoreWebView2WindowCloseRequestedEventHandlerInterface, scoped bool) *ICoreWebView2WindowCloseRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2WindowCloseRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2WindowCloseRequestedEventHandler(impl ICoreWebView2WindowCloseRequestedEventHandlerInterface) *ICoreWebView2WindowCloseRequestedEventHandler {
	return NewICoreWebView2WindowCloseRequestedEventHandlerComObj(impl, true).ICoreWebView2WindowCloseRequestedEventHandler()}

//
type ICoreWebView2WindowCloseRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2WindowCloseRequestedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error
}
func (this *ICoreWebView2WindowCloseRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2WindowCloseRequestedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2WindowCloseRequestedEventHandler {
	impl := &ICoreWebView2WindowCloseRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2WindowCloseRequestedEventHandlerComObj(impl, scoped).ICoreWebView2WindowCloseRequestedEventHandler()
}

