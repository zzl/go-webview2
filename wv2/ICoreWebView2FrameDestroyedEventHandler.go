package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 59DD7B4C-9BAA-11EB-A8B3-0242AC130003
var IID_ICoreWebView2FrameDestroyedEventHandler = syscall.GUID{0x59DD7B4C, 0x9BAA, 0x11EB, 
	[8]byte{0xA8, 0xB3, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2FrameDestroyedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FrameDestroyedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error
}

type ICoreWebView2FrameDestroyedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FrameDestroyedEventHandlerInterface
}

func (this *ICoreWebView2FrameDestroyedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FrameDestroyedEventHandlerInterface)
}

func (this *ICoreWebView2FrameDestroyedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FrameDestroyedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FrameDestroyedEventHandlerImpl) Invoke(sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2FrameDestroyedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FrameDestroyedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FrameDestroyedEventHandlerComObj) impl() ICoreWebView2FrameDestroyedEventHandlerInterface {
	return this.Impl().(ICoreWebView2FrameDestroyedEventHandlerInterface)
}

func (this *ICoreWebView2FrameDestroyedEventHandlerComObj) Invoke(sender *ICoreWebView2Frame, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FrameDestroyedEventHandlerVtbl *ICoreWebView2FrameDestroyedEventHandlerVtbl

func (this *ICoreWebView2FrameDestroyedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FrameDestroyedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2FrameDestroyedEventHandlerVtbl != nil {
		return _pICoreWebView2FrameDestroyedEventHandlerVtbl
	}
	_pICoreWebView2FrameDestroyedEventHandlerVtbl = &ICoreWebView2FrameDestroyedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2FrameDestroyedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FrameDestroyedEventHandlerVtbl
}

func (this *ICoreWebView2FrameDestroyedEventHandlerComObj) ICoreWebView2FrameDestroyedEventHandler() *ICoreWebView2FrameDestroyedEventHandler{
	return (*ICoreWebView2FrameDestroyedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FrameDestroyedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FrameDestroyedEventHandlerComObj(impl ICoreWebView2FrameDestroyedEventHandlerInterface, scoped bool) *ICoreWebView2FrameDestroyedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FrameDestroyedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FrameDestroyedEventHandler(impl ICoreWebView2FrameDestroyedEventHandlerInterface) *ICoreWebView2FrameDestroyedEventHandler {
	return NewICoreWebView2FrameDestroyedEventHandlerComObj(impl, true).ICoreWebView2FrameDestroyedEventHandler()}

//
type ICoreWebView2FrameDestroyedEventHandlerByFuncImpl struct {
	ICoreWebView2FrameDestroyedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error
}
func (this *ICoreWebView2FrameDestroyedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FrameDestroyedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2FrameDestroyedEventHandler {
	impl := &ICoreWebView2FrameDestroyedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FrameDestroyedEventHandlerComObj(impl, scoped).ICoreWebView2FrameDestroyedEventHandler()
}

