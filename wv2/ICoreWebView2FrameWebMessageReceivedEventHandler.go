package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E371E005-6D1D-4517-934B-A8F1629C62A5
var IID_ICoreWebView2FrameWebMessageReceivedEventHandler = syscall.GUID{0xE371E005, 0x6D1D, 0x4517, 
	[8]byte{0x93, 0x4B, 0xA8, 0xF1, 0x62, 0x9C, 0x62, 0xA5}}

type ICoreWebView2FrameWebMessageReceivedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FrameWebMessageReceivedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error
}

type ICoreWebView2FrameWebMessageReceivedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FrameWebMessageReceivedEventHandlerInterface
}

func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FrameWebMessageReceivedEventHandlerInterface)
}

func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FrameWebMessageReceivedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FrameWebMessageReceivedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerComObj) impl() ICoreWebView2FrameWebMessageReceivedEventHandlerInterface {
	return this.Impl().(ICoreWebView2FrameWebMessageReceivedEventHandlerInterface)
}

func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerComObj) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FrameWebMessageReceivedEventHandlerVtbl *ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl

func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2FrameWebMessageReceivedEventHandlerVtbl != nil {
		return _pICoreWebView2FrameWebMessageReceivedEventHandlerVtbl
	}
	_pICoreWebView2FrameWebMessageReceivedEventHandlerVtbl = &ICoreWebView2FrameWebMessageReceivedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2FrameWebMessageReceivedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FrameWebMessageReceivedEventHandlerVtbl
}

func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerComObj) ICoreWebView2FrameWebMessageReceivedEventHandler() *ICoreWebView2FrameWebMessageReceivedEventHandler{
	return (*ICoreWebView2FrameWebMessageReceivedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FrameWebMessageReceivedEventHandlerComObj(impl ICoreWebView2FrameWebMessageReceivedEventHandlerInterface, scoped bool) *ICoreWebView2FrameWebMessageReceivedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FrameWebMessageReceivedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FrameWebMessageReceivedEventHandler(impl ICoreWebView2FrameWebMessageReceivedEventHandlerInterface) *ICoreWebView2FrameWebMessageReceivedEventHandler {
	return NewICoreWebView2FrameWebMessageReceivedEventHandlerComObj(impl, true).ICoreWebView2FrameWebMessageReceivedEventHandler()}

//
type ICoreWebView2FrameWebMessageReceivedEventHandlerByFuncImpl struct {
	ICoreWebView2FrameWebMessageReceivedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error
}
func (this *ICoreWebView2FrameWebMessageReceivedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FrameWebMessageReceivedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Frame, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error, scoped bool) *ICoreWebView2FrameWebMessageReceivedEventHandler {
	impl := &ICoreWebView2FrameWebMessageReceivedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FrameWebMessageReceivedEventHandlerComObj(impl, scoped).ICoreWebView2FrameWebMessageReceivedEventHandler()
}

