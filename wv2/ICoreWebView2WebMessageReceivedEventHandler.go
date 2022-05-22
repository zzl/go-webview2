package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 57213F19-00E6-49FA-8E07-898EA01ECBD2
var IID_ICoreWebView2WebMessageReceivedEventHandler = syscall.GUID{0x57213F19, 0x00E6, 0x49FA, 
	[8]byte{0x8E, 0x07, 0x89, 0x8E, 0xA0, 0x1E, 0xCB, 0xD2}}

type ICoreWebView2WebMessageReceivedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2WebMessageReceivedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error
}

type ICoreWebView2WebMessageReceivedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2WebMessageReceivedEventHandlerInterface
}

func (this *ICoreWebView2WebMessageReceivedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2WebMessageReceivedEventHandlerInterface)
}

func (this *ICoreWebView2WebMessageReceivedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2WebMessageReceivedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2WebMessageReceivedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2WebMessageReceivedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2WebMessageReceivedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2WebMessageReceivedEventHandlerComObj) impl() ICoreWebView2WebMessageReceivedEventHandlerInterface {
	return this.Impl().(ICoreWebView2WebMessageReceivedEventHandlerInterface)
}

func (this *ICoreWebView2WebMessageReceivedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2WebMessageReceivedEventHandlerVtbl *ICoreWebView2WebMessageReceivedEventHandlerVtbl

func (this *ICoreWebView2WebMessageReceivedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2WebMessageReceivedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2WebMessageReceivedEventHandlerVtbl != nil {
		return _pICoreWebView2WebMessageReceivedEventHandlerVtbl
	}
	_pICoreWebView2WebMessageReceivedEventHandlerVtbl = &ICoreWebView2WebMessageReceivedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2WebMessageReceivedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2WebMessageReceivedEventHandlerVtbl
}

func (this *ICoreWebView2WebMessageReceivedEventHandlerComObj) ICoreWebView2WebMessageReceivedEventHandler() *ICoreWebView2WebMessageReceivedEventHandler{
	return (*ICoreWebView2WebMessageReceivedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2WebMessageReceivedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2WebMessageReceivedEventHandlerComObj(impl ICoreWebView2WebMessageReceivedEventHandlerInterface, scoped bool) *ICoreWebView2WebMessageReceivedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2WebMessageReceivedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2WebMessageReceivedEventHandler(impl ICoreWebView2WebMessageReceivedEventHandlerInterface) *ICoreWebView2WebMessageReceivedEventHandler {
	return NewICoreWebView2WebMessageReceivedEventHandlerComObj(impl, true).ICoreWebView2WebMessageReceivedEventHandler()}

//
type ICoreWebView2WebMessageReceivedEventHandlerByFuncImpl struct {
	ICoreWebView2WebMessageReceivedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error
}
func (this *ICoreWebView2WebMessageReceivedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2WebMessageReceivedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2WebMessageReceivedEventArgs) com.Error, scoped bool) *ICoreWebView2WebMessageReceivedEventHandler {
	impl := &ICoreWebView2WebMessageReceivedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2WebMessageReceivedEventHandlerComObj(impl, scoped).ICoreWebView2WebMessageReceivedEventHandler()
}

