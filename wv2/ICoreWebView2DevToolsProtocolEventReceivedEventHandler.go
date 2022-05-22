package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E2FDA4BE-5456-406C-A261-3D452138362C
var IID_ICoreWebView2DevToolsProtocolEventReceivedEventHandler = syscall.GUID{0xE2FDA4BE, 0x5456, 0x406C, 
	[8]byte{0xA2, 0x61, 0x3D, 0x45, 0x21, 0x38, 0x36, 0x2C}}

type ICoreWebView2DevToolsProtocolEventReceivedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) com.Error
}

type ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInterface
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInterface)
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2DevToolsProtocolEventReceivedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj) impl() ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInterface {
	return this.Impl().(ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInterface)
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl != nil {
		return _pICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl
	}
	_pICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl = &ICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2DevToolsProtocolEventReceivedEventHandlerVtbl
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj) ICoreWebView2DevToolsProtocolEventReceivedEventHandler() *ICoreWebView2DevToolsProtocolEventReceivedEventHandler{
	return (*ICoreWebView2DevToolsProtocolEventReceivedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj(impl ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInterface, scoped bool) *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2DevToolsProtocolEventReceivedEventHandler(impl ICoreWebView2DevToolsProtocolEventReceivedEventHandlerInterface) *ICoreWebView2DevToolsProtocolEventReceivedEventHandler {
	return NewICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj(impl, true).ICoreWebView2DevToolsProtocolEventReceivedEventHandler()}

//
type ICoreWebView2DevToolsProtocolEventReceivedEventHandlerByFuncImpl struct {
	ICoreWebView2DevToolsProtocolEventReceivedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) com.Error
}
func (this *ICoreWebView2DevToolsProtocolEventReceivedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2DevToolsProtocolEventReceivedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2DevToolsProtocolEventReceivedEventArgs) com.Error, scoped bool) *ICoreWebView2DevToolsProtocolEventReceivedEventHandler {
	impl := &ICoreWebView2DevToolsProtocolEventReceivedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2DevToolsProtocolEventReceivedEventHandlerComObj(impl, scoped).ICoreWebView2DevToolsProtocolEventReceivedEventHandler()
}

