package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 04D3FE1D-AB87-42FB-A898-DA241D35B63C
var IID_ICoreWebView2ContextMenuRequestedEventHandler = syscall.GUID{0x04D3FE1D, 0xAB87, 0x42FB, 
	[8]byte{0xA8, 0x98, 0xDA, 0x24, 0x1D, 0x35, 0xB6, 0x3C}}

type ICoreWebView2ContextMenuRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2ContextMenuRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) com.Error
}

type ICoreWebView2ContextMenuRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ContextMenuRequestedEventHandlerInterface
}

func (this *ICoreWebView2ContextMenuRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ContextMenuRequestedEventHandlerInterface)
}

func (this *ICoreWebView2ContextMenuRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ContextMenuRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ContextMenuRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2ContextMenuRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ContextMenuRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ContextMenuRequestedEventHandlerComObj) impl() ICoreWebView2ContextMenuRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2ContextMenuRequestedEventHandlerInterface)
}

func (this *ICoreWebView2ContextMenuRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2ContextMenuRequestedEventHandlerVtbl *ICoreWebView2ContextMenuRequestedEventHandlerVtbl

func (this *ICoreWebView2ContextMenuRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ContextMenuRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2ContextMenuRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2ContextMenuRequestedEventHandlerVtbl
	}
	_pICoreWebView2ContextMenuRequestedEventHandlerVtbl = &ICoreWebView2ContextMenuRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2ContextMenuRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2ContextMenuRequestedEventHandlerVtbl
}

func (this *ICoreWebView2ContextMenuRequestedEventHandlerComObj) ICoreWebView2ContextMenuRequestedEventHandler() *ICoreWebView2ContextMenuRequestedEventHandler{
	return (*ICoreWebView2ContextMenuRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ContextMenuRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ContextMenuRequestedEventHandlerComObj(impl ICoreWebView2ContextMenuRequestedEventHandlerInterface, scoped bool) *ICoreWebView2ContextMenuRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ContextMenuRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ContextMenuRequestedEventHandler(impl ICoreWebView2ContextMenuRequestedEventHandlerInterface) *ICoreWebView2ContextMenuRequestedEventHandler {
	return NewICoreWebView2ContextMenuRequestedEventHandlerComObj(impl, true).ICoreWebView2ContextMenuRequestedEventHandler()}

//
type ICoreWebView2ContextMenuRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2ContextMenuRequestedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) com.Error
}
func (this *ICoreWebView2ContextMenuRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2ContextMenuRequestedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2ContextMenuRequestedEventArgs) com.Error, scoped bool) *ICoreWebView2ContextMenuRequestedEventHandler {
	impl := &ICoreWebView2ContextMenuRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ContextMenuRequestedEventHandlerComObj(impl, scoped).ICoreWebView2ContextMenuRequestedEventHandler()
}

