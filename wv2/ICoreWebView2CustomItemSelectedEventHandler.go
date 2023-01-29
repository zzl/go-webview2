package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 49E1D0BC-FE9E-4481-B7C2-32324AA21998
var IID_ICoreWebView2CustomItemSelectedEventHandler = syscall.GUID{0x49E1D0BC, 0xFE9E, 0x4481,
	[8]byte{0xB7, 0xC2, 0x32, 0x32, 0x4A, 0xA2, 0x19, 0x98}}

type ICoreWebView2CustomItemSelectedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2CustomItemSelectedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2ContextMenuItem, args *win32.IUnknown) com.Error
}

type ICoreWebView2CustomItemSelectedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2CustomItemSelectedEventHandlerInterface
}

func (this *ICoreWebView2CustomItemSelectedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2CustomItemSelectedEventHandlerInterface)
}

func (this *ICoreWebView2CustomItemSelectedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2CustomItemSelectedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2CustomItemSelectedEventHandlerImpl) Invoke(sender *ICoreWebView2ContextMenuItem, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2CustomItemSelectedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2CustomItemSelectedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2CustomItemSelectedEventHandlerComObj) impl() ICoreWebView2CustomItemSelectedEventHandlerInterface {
	return this.Impl().(ICoreWebView2CustomItemSelectedEventHandlerInterface)
}

func (this *ICoreWebView2CustomItemSelectedEventHandlerComObj) Invoke(sender *ICoreWebView2ContextMenuItem, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2CustomItemSelectedEventHandlerVtbl *ICoreWebView2CustomItemSelectedEventHandlerVtbl

func (this *ICoreWebView2CustomItemSelectedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2CustomItemSelectedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2CustomItemSelectedEventHandlerVtbl != nil {
		return _pICoreWebView2CustomItemSelectedEventHandlerVtbl
	}
	_pICoreWebView2CustomItemSelectedEventHandlerVtbl = &ICoreWebView2CustomItemSelectedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2CustomItemSelectedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2CustomItemSelectedEventHandlerVtbl
}

func (this *ICoreWebView2CustomItemSelectedEventHandlerComObj) ICoreWebView2CustomItemSelectedEventHandler() *ICoreWebView2CustomItemSelectedEventHandler {
	return (*ICoreWebView2CustomItemSelectedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2CustomItemSelectedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2CustomItemSelectedEventHandlerComObj(impl ICoreWebView2CustomItemSelectedEventHandlerInterface, scoped bool) *ICoreWebView2CustomItemSelectedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2CustomItemSelectedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2CustomItemSelectedEventHandler(impl ICoreWebView2CustomItemSelectedEventHandlerInterface) *ICoreWebView2CustomItemSelectedEventHandler {
	return NewICoreWebView2CustomItemSelectedEventHandlerComObj(impl, true).ICoreWebView2CustomItemSelectedEventHandler()
}

//
type ICoreWebView2CustomItemSelectedEventHandlerByFuncImpl struct {
	ICoreWebView2CustomItemSelectedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2ContextMenuItem, args *win32.IUnknown) com.Error
}

func (this *ICoreWebView2CustomItemSelectedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2ContextMenuItem, args *win32.IUnknown) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2CustomItemSelectedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2ContextMenuItem, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2CustomItemSelectedEventHandler {
	impl := &ICoreWebView2CustomItemSelectedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2CustomItemSelectedEventHandlerComObj(impl, scoped).ICoreWebView2CustomItemSelectedEventHandler()
}

