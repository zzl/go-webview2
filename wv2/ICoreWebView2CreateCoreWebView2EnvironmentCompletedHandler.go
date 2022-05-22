package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 4E8A3389-C9D8-4BD2-B6B5-124FEE6CC14D
var IID_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler = syscall.GUID{0x4E8A3389, 0xC9D8, 0x4BD2, 
	[8]byte{0xB6, 0xB5, 0x12, 0x4F, 0xEE, 0x6C, 0xC1, 0x4D}}

type ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, createdEnvironment *ICoreWebView2Environment) com.Error
}

type ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInterface
}

func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInterface)
}

func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl) Invoke(errorCode com.Error, createdEnvironment *ICoreWebView2Environment) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj) impl() ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInterface)
}

func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj) Invoke(errorCode com.Error, createdEnvironment *ICoreWebView2Environment) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, createdEnvironment))
}

var _pICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl

func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl != nil {
		return _pICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl
	}
	_pICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl = &ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerVtbl
}

func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj) ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler() *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler{
	return (*ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj(impl ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInterface, scoped bool) *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler(impl ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerInterface) *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler {
	return NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj(impl, true).ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler()}

//
type ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerByFuncImpl struct {
	ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerImpl
	handlerFunc func (errorCode com.Error, createdEnvironment *ICoreWebView2Environment) com.Error
}
func (this *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, createdEnvironment *ICoreWebView2Environment) com.Error{
	return this.handlerFunc(errorCode, createdEnvironment)
}

func NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerByFunc(handlerFunc func (errorCode com.Error, createdEnvironment *ICoreWebView2Environment) com.Error, scoped bool) *ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler {
	impl := &ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2CreateCoreWebView2EnvironmentCompletedHandlerComObj(impl, scoped).ICoreWebView2CreateCoreWebView2EnvironmentCompletedHandler()
}

