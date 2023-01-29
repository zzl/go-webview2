package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 9ADBE429-F36D-432B-9DDC-F8881FBD76E3
var IID_ICoreWebView2NavigationStartingEventHandler = syscall.GUID{0x9ADBE429, 0xF36D, 0x432B,
	[8]byte{0x9D, 0xDC, 0xF8, 0x88, 0x1F, 0xBD, 0x76, 0xE3}}

type ICoreWebView2NavigationStartingEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2NavigationStartingEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) com.Error
}

type ICoreWebView2NavigationStartingEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2NavigationStartingEventHandlerInterface
}

func (this *ICoreWebView2NavigationStartingEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2NavigationStartingEventHandlerInterface)
}

func (this *ICoreWebView2NavigationStartingEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2NavigationStartingEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2NavigationStartingEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2NavigationStartingEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2NavigationStartingEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2NavigationStartingEventHandlerComObj) impl() ICoreWebView2NavigationStartingEventHandlerInterface {
	return this.Impl().(ICoreWebView2NavigationStartingEventHandlerInterface)
}

func (this *ICoreWebView2NavigationStartingEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2NavigationStartingEventHandlerVtbl *ICoreWebView2NavigationStartingEventHandlerVtbl

func (this *ICoreWebView2NavigationStartingEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2NavigationStartingEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2NavigationStartingEventHandlerVtbl != nil {
		return _pICoreWebView2NavigationStartingEventHandlerVtbl
	}
	_pICoreWebView2NavigationStartingEventHandlerVtbl = &ICoreWebView2NavigationStartingEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2NavigationStartingEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2NavigationStartingEventHandlerVtbl
}

func (this *ICoreWebView2NavigationStartingEventHandlerComObj) ICoreWebView2NavigationStartingEventHandler() *ICoreWebView2NavigationStartingEventHandler {
	return (*ICoreWebView2NavigationStartingEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2NavigationStartingEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2NavigationStartingEventHandlerComObj(impl ICoreWebView2NavigationStartingEventHandlerInterface, scoped bool) *ICoreWebView2NavigationStartingEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2NavigationStartingEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2NavigationStartingEventHandler(impl ICoreWebView2NavigationStartingEventHandlerInterface) *ICoreWebView2NavigationStartingEventHandler {
	return NewICoreWebView2NavigationStartingEventHandlerComObj(impl, true).ICoreWebView2NavigationStartingEventHandler()
}

//
type ICoreWebView2NavigationStartingEventHandlerByFuncImpl struct {
	ICoreWebView2NavigationStartingEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) com.Error
}

func (this *ICoreWebView2NavigationStartingEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2NavigationStartingEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2NavigationStartingEventArgs) com.Error, scoped bool) *ICoreWebView2NavigationStartingEventHandler {
	impl := &ICoreWebView2NavigationStartingEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2NavigationStartingEventHandlerComObj(impl, scoped).ICoreWebView2NavigationStartingEventHandler()
}

