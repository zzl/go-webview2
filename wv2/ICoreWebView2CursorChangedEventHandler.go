package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 9DA43CCC-26E1-4DAD-B56C-D8961C94C571
var IID_ICoreWebView2CursorChangedEventHandler = syscall.GUID{0x9DA43CCC, 0x26E1, 0x4DAD,
	[8]byte{0xB5, 0x6C, 0xD8, 0x96, 0x1C, 0x94, 0xC5, 0x71}}

type ICoreWebView2CursorChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2CursorChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2CompositionController, args *win32.IUnknown) com.Error
}

type ICoreWebView2CursorChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2CursorChangedEventHandlerInterface
}

func (this *ICoreWebView2CursorChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2CursorChangedEventHandlerInterface)
}

func (this *ICoreWebView2CursorChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2CursorChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2CursorChangedEventHandlerImpl) Invoke(sender *ICoreWebView2CompositionController, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2CursorChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2CursorChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2CursorChangedEventHandlerComObj) impl() ICoreWebView2CursorChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2CursorChangedEventHandlerInterface)
}

func (this *ICoreWebView2CursorChangedEventHandlerComObj) Invoke(sender *ICoreWebView2CompositionController, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2CursorChangedEventHandlerVtbl *ICoreWebView2CursorChangedEventHandlerVtbl

func (this *ICoreWebView2CursorChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2CursorChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2CursorChangedEventHandlerVtbl != nil {
		return _pICoreWebView2CursorChangedEventHandlerVtbl
	}
	_pICoreWebView2CursorChangedEventHandlerVtbl = &ICoreWebView2CursorChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2CursorChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2CursorChangedEventHandlerVtbl
}

func (this *ICoreWebView2CursorChangedEventHandlerComObj) ICoreWebView2CursorChangedEventHandler() *ICoreWebView2CursorChangedEventHandler {
	return (*ICoreWebView2CursorChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2CursorChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2CursorChangedEventHandlerComObj(impl ICoreWebView2CursorChangedEventHandlerInterface, scoped bool) *ICoreWebView2CursorChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2CursorChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2CursorChangedEventHandler(impl ICoreWebView2CursorChangedEventHandlerInterface) *ICoreWebView2CursorChangedEventHandler {
	return NewICoreWebView2CursorChangedEventHandlerComObj(impl, true).ICoreWebView2CursorChangedEventHandler()
}

//
type ICoreWebView2CursorChangedEventHandlerByFuncImpl struct {
	ICoreWebView2CursorChangedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2CompositionController, args *win32.IUnknown) com.Error
}

func (this *ICoreWebView2CursorChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2CompositionController, args *win32.IUnknown) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2CursorChangedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2CompositionController, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2CursorChangedEventHandler {
	impl := &ICoreWebView2CursorChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2CursorChangedEventHandlerComObj(impl, scoped).ICoreWebView2CursorChangedEventHandler()
}

