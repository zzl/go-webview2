package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 5C4889F0-5EF6-4C5A-952C-D8F1B92D0574
var IID_ICoreWebView2CallDevToolsProtocolMethodCompletedHandler = syscall.GUID{0x5C4889F0, 0x5EF6, 0x4C5A, 
	[8]byte{0x95, 0x2C, 0xD8, 0xF1, 0xB9, 0x2D, 0x05, 0x74}}

type ICoreWebView2CallDevToolsProtocolMethodCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, returnObjectAsJson string) com.Error
}

type ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerInterface
}

func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerInterface)
}

func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2CallDevToolsProtocolMethodCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerImpl) Invoke(errorCode com.Error, returnObjectAsJson string) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj) impl() ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerInterface)
}

func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj) Invoke(errorCode com.Error, returnObjectAsJson win32.PWSTR) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, win32.PwstrToStr(returnObjectAsJson)))
}

var _pICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl

func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl != nil {
		return _pICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl
	}
	_pICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl = &ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2CallDevToolsProtocolMethodCompletedHandlerVtbl
}

func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj) ICoreWebView2CallDevToolsProtocolMethodCompletedHandler() *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler{
	return (*ICoreWebView2CallDevToolsProtocolMethodCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj(impl ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerInterface, scoped bool) *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2CallDevToolsProtocolMethodCompletedHandler(impl ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerInterface) *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler {
	return NewICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj(impl, true).ICoreWebView2CallDevToolsProtocolMethodCompletedHandler()}

//
type ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerByFuncImpl struct {
	ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerImpl
	handlerFunc func (errorCode com.Error, returnObjectAsJson string) com.Error
}
func (this *ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, returnObjectAsJson string) com.Error{
	return this.handlerFunc(errorCode, returnObjectAsJson)
}

func NewICoreWebView2CallDevToolsProtocolMethodCompletedHandlerByFunc(handlerFunc func (errorCode com.Error, returnObjectAsJson string) com.Error, scoped bool) *ICoreWebView2CallDevToolsProtocolMethodCompletedHandler {
	impl := &ICoreWebView2CallDevToolsProtocolMethodCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2CallDevToolsProtocolMethodCompletedHandlerComObj(impl, scoped).ICoreWebView2CallDevToolsProtocolMethodCompletedHandler()
}

