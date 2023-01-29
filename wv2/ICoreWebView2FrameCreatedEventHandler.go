package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 38059770-9BAA-11EB-A8B3-0242AC130003
var IID_ICoreWebView2FrameCreatedEventHandler = syscall.GUID{0x38059770, 0x9BAA, 0x11EB,
	[8]byte{0xA8, 0xB3, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2FrameCreatedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FrameCreatedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) com.Error
}

type ICoreWebView2FrameCreatedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FrameCreatedEventHandlerInterface
}

func (this *ICoreWebView2FrameCreatedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FrameCreatedEventHandlerInterface)
}

func (this *ICoreWebView2FrameCreatedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FrameCreatedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FrameCreatedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2FrameCreatedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FrameCreatedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FrameCreatedEventHandlerComObj) impl() ICoreWebView2FrameCreatedEventHandlerInterface {
	return this.Impl().(ICoreWebView2FrameCreatedEventHandlerInterface)
}

func (this *ICoreWebView2FrameCreatedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FrameCreatedEventHandlerVtbl *ICoreWebView2FrameCreatedEventHandlerVtbl

func (this *ICoreWebView2FrameCreatedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FrameCreatedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2FrameCreatedEventHandlerVtbl != nil {
		return _pICoreWebView2FrameCreatedEventHandlerVtbl
	}
	_pICoreWebView2FrameCreatedEventHandlerVtbl = &ICoreWebView2FrameCreatedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2FrameCreatedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FrameCreatedEventHandlerVtbl
}

func (this *ICoreWebView2FrameCreatedEventHandlerComObj) ICoreWebView2FrameCreatedEventHandler() *ICoreWebView2FrameCreatedEventHandler {
	return (*ICoreWebView2FrameCreatedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FrameCreatedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FrameCreatedEventHandlerComObj(impl ICoreWebView2FrameCreatedEventHandlerInterface, scoped bool) *ICoreWebView2FrameCreatedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FrameCreatedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FrameCreatedEventHandler(impl ICoreWebView2FrameCreatedEventHandlerInterface) *ICoreWebView2FrameCreatedEventHandler {
	return NewICoreWebView2FrameCreatedEventHandlerComObj(impl, true).ICoreWebView2FrameCreatedEventHandler()
}

type ICoreWebView2FrameCreatedEventHandlerByFuncImpl struct {
	ICoreWebView2FrameCreatedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) com.Error
}

func (this *ICoreWebView2FrameCreatedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FrameCreatedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2FrameCreatedEventArgs) com.Error, scoped bool) *ICoreWebView2FrameCreatedEventHandler {
	impl := &ICoreWebView2FrameCreatedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FrameCreatedEventHandlerComObj(impl, scoped).ICoreWebView2FrameCreatedEventHandler()
}
