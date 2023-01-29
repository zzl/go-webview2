package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 49511172-CC67-4BCA-9923-137112F4C4CC
var IID_ICoreWebView2ExecuteScriptCompletedHandler = syscall.GUID{0x49511172, 0xCC67, 0x4BCA,
	[8]byte{0x99, 0x23, 0x13, 0x71, 0x12, 0xF4, 0xC4, 0xCC}}

type ICoreWebView2ExecuteScriptCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2ExecuteScriptCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, resultObjectAsJson string) com.Error
}

type ICoreWebView2ExecuteScriptCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ExecuteScriptCompletedHandlerInterface
}

func (this *ICoreWebView2ExecuteScriptCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ExecuteScriptCompletedHandlerInterface)
}

func (this *ICoreWebView2ExecuteScriptCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ExecuteScriptCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ExecuteScriptCompletedHandlerImpl) Invoke(errorCode com.Error, resultObjectAsJson string) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2ExecuteScriptCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ExecuteScriptCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ExecuteScriptCompletedHandlerComObj) impl() ICoreWebView2ExecuteScriptCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2ExecuteScriptCompletedHandlerInterface)
}

func (this *ICoreWebView2ExecuteScriptCompletedHandlerComObj) Invoke(errorCode com.Error, resultObjectAsJson win32.PWSTR) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, win32.PwstrToStr(resultObjectAsJson)))
}

var _pICoreWebView2ExecuteScriptCompletedHandlerVtbl *ICoreWebView2ExecuteScriptCompletedHandlerVtbl

func (this *ICoreWebView2ExecuteScriptCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ExecuteScriptCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2ExecuteScriptCompletedHandlerVtbl != nil {
		return _pICoreWebView2ExecuteScriptCompletedHandlerVtbl
	}
	_pICoreWebView2ExecuteScriptCompletedHandlerVtbl = &ICoreWebView2ExecuteScriptCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2ExecuteScriptCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2ExecuteScriptCompletedHandlerVtbl
}

func (this *ICoreWebView2ExecuteScriptCompletedHandlerComObj) ICoreWebView2ExecuteScriptCompletedHandler() *ICoreWebView2ExecuteScriptCompletedHandler {
	return (*ICoreWebView2ExecuteScriptCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ExecuteScriptCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ExecuteScriptCompletedHandlerComObj(impl ICoreWebView2ExecuteScriptCompletedHandlerInterface, scoped bool) *ICoreWebView2ExecuteScriptCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ExecuteScriptCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ExecuteScriptCompletedHandler(impl ICoreWebView2ExecuteScriptCompletedHandlerInterface) *ICoreWebView2ExecuteScriptCompletedHandler {
	return NewICoreWebView2ExecuteScriptCompletedHandlerComObj(impl, true).ICoreWebView2ExecuteScriptCompletedHandler()
}

//
type ICoreWebView2ExecuteScriptCompletedHandlerByFuncImpl struct {
	ICoreWebView2ExecuteScriptCompletedHandlerImpl
	handlerFunc func(errorCode com.Error, resultObjectAsJson string) com.Error
}

func (this *ICoreWebView2ExecuteScriptCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, resultObjectAsJson string) com.Error {
	return this.handlerFunc(errorCode, resultObjectAsJson)
}

func NewICoreWebView2ExecuteScriptCompletedHandlerByFunc(handlerFunc func(errorCode com.Error, resultObjectAsJson string) com.Error, scoped bool) *ICoreWebView2ExecuteScriptCompletedHandler {
	impl := &ICoreWebView2ExecuteScriptCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ExecuteScriptCompletedHandlerComObj(impl, scoped).ICoreWebView2ExecuteScriptCompletedHandler()
}

