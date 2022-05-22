package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// B99369F3-9B11-47B5-BC6F-8E7895FCEA17
var IID_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler = syscall.GUID{0xB99369F3, 0x9B11, 0x47B5, 
	[8]byte{0xBC, 0x6F, 0x8E, 0x78, 0x95, 0xFC, 0xEA, 0x17}}

type ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, id string) com.Error
}

type ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInterface
}

func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInterface)
}

func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerImpl) Invoke(errorCode com.Error, id string) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj) impl() ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInterface)
}

func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj) Invoke(errorCode com.Error, id win32.PWSTR) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, win32.PwstrToStr(id)))
}

var _pICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl

func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl != nil {
		return _pICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl
	}
	_pICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl = &ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerVtbl
}

func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj) ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler() *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler{
	return (*ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj(impl ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInterface, scoped bool) *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler(impl ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerInterface) *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler {
	return NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj(impl, true).ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler()}

//
type ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerByFuncImpl struct {
	ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerImpl
	handlerFunc func (errorCode com.Error, id string) com.Error
}
func (this *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, id string) com.Error{
	return this.handlerFunc(errorCode, id)
}

func NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerByFunc(handlerFunc func (errorCode com.Error, id string) com.Error, scoped bool) *ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler {
	impl := &ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandlerComObj(impl, scoped).ICoreWebView2AddScriptToExecuteOnDocumentCreatedCompletedHandler()
}

