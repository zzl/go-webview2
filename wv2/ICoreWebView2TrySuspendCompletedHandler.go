package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 00F206A7-9D17-4605-91F6-4E8E4DE192E3
var IID_ICoreWebView2TrySuspendCompletedHandler = syscall.GUID{0x00F206A7, 0x9D17, 0x4605, 
	[8]byte{0x91, 0xF6, 0x4E, 0x8E, 0x4D, 0xE1, 0x92, 0xE3}}

type ICoreWebView2TrySuspendCompletedHandler struct {
	win32.IUnknown
}

type ICoreWebView2TrySuspendCompletedHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(errorCode com.Error, isSuccessful int32) com.Error
}

type ICoreWebView2TrySuspendCompletedHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2TrySuspendCompletedHandlerInterface
}

func (this *ICoreWebView2TrySuspendCompletedHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2TrySuspendCompletedHandlerInterface)
}

func (this *ICoreWebView2TrySuspendCompletedHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2TrySuspendCompletedHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2TrySuspendCompletedHandlerImpl) Invoke(errorCode com.Error, isSuccessful int32) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2TrySuspendCompletedHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2TrySuspendCompletedHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2TrySuspendCompletedHandlerComObj) impl() ICoreWebView2TrySuspendCompletedHandlerInterface {
	return this.Impl().(ICoreWebView2TrySuspendCompletedHandlerInterface)
}

func (this *ICoreWebView2TrySuspendCompletedHandlerComObj) Invoke(errorCode com.Error, isSuccessful int32) uintptr {
	return (uintptr)(this.impl().Invoke(errorCode, isSuccessful))
}

var _pICoreWebView2TrySuspendCompletedHandlerVtbl *ICoreWebView2TrySuspendCompletedHandlerVtbl

func (this *ICoreWebView2TrySuspendCompletedHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2TrySuspendCompletedHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2TrySuspendCompletedHandlerVtbl != nil {
		return _pICoreWebView2TrySuspendCompletedHandlerVtbl
	}
	_pICoreWebView2TrySuspendCompletedHandlerVtbl = &ICoreWebView2TrySuspendCompletedHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2TrySuspendCompletedHandlerComObj).Invoke),
	}
	return _pICoreWebView2TrySuspendCompletedHandlerVtbl
}

func (this *ICoreWebView2TrySuspendCompletedHandlerComObj) ICoreWebView2TrySuspendCompletedHandler() *ICoreWebView2TrySuspendCompletedHandler{
	return (*ICoreWebView2TrySuspendCompletedHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2TrySuspendCompletedHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2TrySuspendCompletedHandlerComObj(impl ICoreWebView2TrySuspendCompletedHandlerInterface, scoped bool) *ICoreWebView2TrySuspendCompletedHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2TrySuspendCompletedHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2TrySuspendCompletedHandler(impl ICoreWebView2TrySuspendCompletedHandlerInterface) *ICoreWebView2TrySuspendCompletedHandler {
	return NewICoreWebView2TrySuspendCompletedHandlerComObj(impl, true).ICoreWebView2TrySuspendCompletedHandler()}

//
type ICoreWebView2TrySuspendCompletedHandlerByFuncImpl struct {
	ICoreWebView2TrySuspendCompletedHandlerImpl
	handlerFunc func (errorCode com.Error, isSuccessful int32) com.Error
}
func (this *ICoreWebView2TrySuspendCompletedHandlerByFuncImpl) Invoke(errorCode com.Error, isSuccessful int32) com.Error{
	return this.handlerFunc(errorCode, isSuccessful)
}

func NewICoreWebView2TrySuspendCompletedHandlerByFunc(handlerFunc func (errorCode com.Error, isSuccessful int32) com.Error, scoped bool) *ICoreWebView2TrySuspendCompletedHandler {
	impl := &ICoreWebView2TrySuspendCompletedHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2TrySuspendCompletedHandlerComObj(impl, scoped).ICoreWebView2TrySuspendCompletedHandler()
}

