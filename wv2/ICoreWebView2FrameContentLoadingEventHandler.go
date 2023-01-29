package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 0D6156F2-D332-49A7-9E03-7D8F2FEEEE54
var IID_ICoreWebView2FrameContentLoadingEventHandler = syscall.GUID{0x0D6156F2, 0xD332, 0x49A7,
	[8]byte{0x9E, 0x03, 0x7D, 0x8F, 0x2F, 0xEE, 0xEE, 0x54}}

type ICoreWebView2FrameContentLoadingEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FrameContentLoadingEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) com.Error
}

type ICoreWebView2FrameContentLoadingEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FrameContentLoadingEventHandlerInterface
}

func (this *ICoreWebView2FrameContentLoadingEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FrameContentLoadingEventHandlerInterface)
}

func (this *ICoreWebView2FrameContentLoadingEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FrameContentLoadingEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FrameContentLoadingEventHandlerImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2FrameContentLoadingEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FrameContentLoadingEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FrameContentLoadingEventHandlerComObj) impl() ICoreWebView2FrameContentLoadingEventHandlerInterface {
	return this.Impl().(ICoreWebView2FrameContentLoadingEventHandlerInterface)
}

func (this *ICoreWebView2FrameContentLoadingEventHandlerComObj) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FrameContentLoadingEventHandlerVtbl *ICoreWebView2FrameContentLoadingEventHandlerVtbl

func (this *ICoreWebView2FrameContentLoadingEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FrameContentLoadingEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2FrameContentLoadingEventHandlerVtbl != nil {
		return _pICoreWebView2FrameContentLoadingEventHandlerVtbl
	}
	_pICoreWebView2FrameContentLoadingEventHandlerVtbl = &ICoreWebView2FrameContentLoadingEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2FrameContentLoadingEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FrameContentLoadingEventHandlerVtbl
}

func (this *ICoreWebView2FrameContentLoadingEventHandlerComObj) ICoreWebView2FrameContentLoadingEventHandler() *ICoreWebView2FrameContentLoadingEventHandler {
	return (*ICoreWebView2FrameContentLoadingEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FrameContentLoadingEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FrameContentLoadingEventHandlerComObj(impl ICoreWebView2FrameContentLoadingEventHandlerInterface, scoped bool) *ICoreWebView2FrameContentLoadingEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FrameContentLoadingEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FrameContentLoadingEventHandler(impl ICoreWebView2FrameContentLoadingEventHandlerInterface) *ICoreWebView2FrameContentLoadingEventHandler {
	return NewICoreWebView2FrameContentLoadingEventHandlerComObj(impl, true).ICoreWebView2FrameContentLoadingEventHandler()
}

//
type ICoreWebView2FrameContentLoadingEventHandlerByFuncImpl struct {
	ICoreWebView2FrameContentLoadingEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) com.Error
}

func (this *ICoreWebView2FrameContentLoadingEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FrameContentLoadingEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2Frame, args *ICoreWebView2ContentLoadingEventArgs) com.Error, scoped bool) *ICoreWebView2FrameContentLoadingEventHandler {
	impl := &ICoreWebView2FrameContentLoadingEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FrameContentLoadingEventHandlerComObj(impl, scoped).ICoreWebView2FrameContentLoadingEventHandler()
}

