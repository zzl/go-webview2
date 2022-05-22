package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 05EA24BD-6452-4926-9014-4B82B498135D
var IID_ICoreWebView2FocusChangedEventHandler = syscall.GUID{0x05EA24BD, 0x6452, 0x4926, 
	[8]byte{0x90, 0x14, 0x4B, 0x82, 0xB4, 0x98, 0x13, 0x5D}}

type ICoreWebView2FocusChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2FocusChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error
}

type ICoreWebView2FocusChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2FocusChangedEventHandlerInterface
}

func (this *ICoreWebView2FocusChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2FocusChangedEventHandlerInterface)
}

func (this *ICoreWebView2FocusChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2FocusChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2FocusChangedEventHandlerImpl) Invoke(sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2FocusChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2FocusChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2FocusChangedEventHandlerComObj) impl() ICoreWebView2FocusChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2FocusChangedEventHandlerInterface)
}

func (this *ICoreWebView2FocusChangedEventHandlerComObj) Invoke(sender *ICoreWebView2Controller, args *com.UnknownClass) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2FocusChangedEventHandlerVtbl *ICoreWebView2FocusChangedEventHandlerVtbl

func (this *ICoreWebView2FocusChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2FocusChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2FocusChangedEventHandlerVtbl != nil {
		return _pICoreWebView2FocusChangedEventHandlerVtbl
	}
	_pICoreWebView2FocusChangedEventHandlerVtbl = &ICoreWebView2FocusChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2FocusChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2FocusChangedEventHandlerVtbl
}

func (this *ICoreWebView2FocusChangedEventHandlerComObj) ICoreWebView2FocusChangedEventHandler() *ICoreWebView2FocusChangedEventHandler{
	return (*ICoreWebView2FocusChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2FocusChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2FocusChangedEventHandlerComObj(impl ICoreWebView2FocusChangedEventHandlerInterface, scoped bool) *ICoreWebView2FocusChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2FocusChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2FocusChangedEventHandler(impl ICoreWebView2FocusChangedEventHandlerInterface) *ICoreWebView2FocusChangedEventHandler {
	return NewICoreWebView2FocusChangedEventHandlerComObj(impl, true).ICoreWebView2FocusChangedEventHandler()}

//
type ICoreWebView2FocusChangedEventHandlerByFuncImpl struct {
	ICoreWebView2FocusChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error
}
func (this *ICoreWebView2FocusChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2FocusChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Controller, args *com.UnknownClass) com.Error, scoped bool) *ICoreWebView2FocusChangedEventHandler {
	impl := &ICoreWebView2FocusChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2FocusChangedEventHandlerComObj(impl, scoped).ICoreWebView2FocusChangedEventHandler()
}

