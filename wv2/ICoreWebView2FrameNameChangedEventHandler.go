package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 435C7DC8-9BAA-11EB-A8B3-0242AC130003
var IID_ICoreWebView2FrameNameChangedEventHandler = syscall.GUID{0x435C7DC8, 0x9BAA, 0x11EB, 
	[8]byte{0xA8, 0xB3, 0x02, 0x42, 0xAC, 0x13, 0x00, 0x03}}

type ICoreWebView2FrameNameChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FrameNameChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error
}

type ICoreWebView2FrameNameChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FrameNameChangedEventHandlerInterface
}

func (this *ICoreWebView2FrameNameChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FrameNameChangedEventHandlerInterface)
}

func (this *ICoreWebView2FrameNameChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FrameNameChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FrameNameChangedEventHandlerImpl) Invoke(sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2FrameNameChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FrameNameChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FrameNameChangedEventHandlerComObj) impl() ICoreWebView2FrameNameChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2FrameNameChangedEventHandlerInterface)
}

func (this *ICoreWebView2FrameNameChangedEventHandlerComObj) Invoke(sender *ICoreWebView2Frame, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FrameNameChangedEventHandlerVtbl *ICoreWebView2FrameNameChangedEventHandlerVtbl

func (this *ICoreWebView2FrameNameChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FrameNameChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2FrameNameChangedEventHandlerVtbl != nil {
		return _pICoreWebView2FrameNameChangedEventHandlerVtbl
	}
	_pICoreWebView2FrameNameChangedEventHandlerVtbl = &ICoreWebView2FrameNameChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2FrameNameChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FrameNameChangedEventHandlerVtbl
}

func (this *ICoreWebView2FrameNameChangedEventHandlerComObj) ICoreWebView2FrameNameChangedEventHandler() *ICoreWebView2FrameNameChangedEventHandler{
	return (*ICoreWebView2FrameNameChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FrameNameChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FrameNameChangedEventHandlerComObj(impl ICoreWebView2FrameNameChangedEventHandlerInterface, scoped bool) *ICoreWebView2FrameNameChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FrameNameChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FrameNameChangedEventHandler(impl ICoreWebView2FrameNameChangedEventHandlerInterface) *ICoreWebView2FrameNameChangedEventHandler {
	return NewICoreWebView2FrameNameChangedEventHandlerComObj(impl, true).ICoreWebView2FrameNameChangedEventHandler()}

//
type ICoreWebView2FrameNameChangedEventHandlerByFuncImpl struct {
	ICoreWebView2FrameNameChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error
}
func (this *ICoreWebView2FrameNameChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FrameNameChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Frame, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2FrameNameChangedEventHandler {
	impl := &ICoreWebView2FrameNameChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FrameNameChangedEventHandlerComObj(impl, scoped).ICoreWebView2FrameNameChangedEventHandler()
}

