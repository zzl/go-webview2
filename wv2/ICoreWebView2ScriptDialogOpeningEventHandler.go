package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// EF381BF9-AFA8-4E37-91C4-8AC48524BDFB
var IID_ICoreWebView2ScriptDialogOpeningEventHandler = syscall.GUID{0xEF381BF9, 0xAFA8, 0x4E37,
	[8]byte{0x91, 0xC4, 0x8A, 0xC4, 0x85, 0x24, 0xBD, 0xFB}}

type ICoreWebView2ScriptDialogOpeningEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2ScriptDialogOpeningEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) com.Error
}

type ICoreWebView2ScriptDialogOpeningEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ScriptDialogOpeningEventHandlerInterface
}

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ScriptDialogOpeningEventHandlerInterface)
}

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ScriptDialogOpeningEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2ScriptDialogOpeningEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ScriptDialogOpeningEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerComObj) impl() ICoreWebView2ScriptDialogOpeningEventHandlerInterface {
	return this.Impl().(ICoreWebView2ScriptDialogOpeningEventHandlerInterface)
}

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2ScriptDialogOpeningEventHandlerVtbl *ICoreWebView2ScriptDialogOpeningEventHandlerVtbl

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ScriptDialogOpeningEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2ScriptDialogOpeningEventHandlerVtbl != nil {
		return _pICoreWebView2ScriptDialogOpeningEventHandlerVtbl
	}
	_pICoreWebView2ScriptDialogOpeningEventHandlerVtbl = &ICoreWebView2ScriptDialogOpeningEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2ScriptDialogOpeningEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2ScriptDialogOpeningEventHandlerVtbl
}

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerComObj) ICoreWebView2ScriptDialogOpeningEventHandler() *ICoreWebView2ScriptDialogOpeningEventHandler {
	return (*ICoreWebView2ScriptDialogOpeningEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ScriptDialogOpeningEventHandlerComObj(impl ICoreWebView2ScriptDialogOpeningEventHandlerInterface, scoped bool) *ICoreWebView2ScriptDialogOpeningEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ScriptDialogOpeningEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ScriptDialogOpeningEventHandler(impl ICoreWebView2ScriptDialogOpeningEventHandlerInterface) *ICoreWebView2ScriptDialogOpeningEventHandler {
	return NewICoreWebView2ScriptDialogOpeningEventHandlerComObj(impl, true).ICoreWebView2ScriptDialogOpeningEventHandler()
}

type ICoreWebView2ScriptDialogOpeningEventHandlerByFuncImpl struct {
	ICoreWebView2ScriptDialogOpeningEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) com.Error
}

func (this *ICoreWebView2ScriptDialogOpeningEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2ScriptDialogOpeningEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2ScriptDialogOpeningEventArgs) com.Error, scoped bool) *ICoreWebView2ScriptDialogOpeningEventHandler {
	impl := &ICoreWebView2ScriptDialogOpeningEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ScriptDialogOpeningEventHandlerComObj(impl, scoped).ICoreWebView2ScriptDialogOpeningEventHandler()
}
