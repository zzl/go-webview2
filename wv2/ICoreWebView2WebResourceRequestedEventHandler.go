package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// AB00B74C-15F1-4646-80E8-E76341D25D71
var IID_ICoreWebView2WebResourceRequestedEventHandler = syscall.GUID{0xAB00B74C, 0x15F1, 0x4646, 
	[8]byte{0x80, 0xE8, 0xE7, 0x63, 0x41, 0xD2, 0x5D, 0x71}}

type ICoreWebView2WebResourceRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2WebResourceRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) com.Error
}

type ICoreWebView2WebResourceRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2WebResourceRequestedEventHandlerInterface
}

func (this *ICoreWebView2WebResourceRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2WebResourceRequestedEventHandlerInterface)
}

func (this *ICoreWebView2WebResourceRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2WebResourceRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2WebResourceRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2WebResourceRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2WebResourceRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2WebResourceRequestedEventHandlerComObj) impl() ICoreWebView2WebResourceRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2WebResourceRequestedEventHandlerInterface)
}

func (this *ICoreWebView2WebResourceRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2WebResourceRequestedEventHandlerVtbl *ICoreWebView2WebResourceRequestedEventHandlerVtbl

func (this *ICoreWebView2WebResourceRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2WebResourceRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2WebResourceRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2WebResourceRequestedEventHandlerVtbl
	}
	_pICoreWebView2WebResourceRequestedEventHandlerVtbl = &ICoreWebView2WebResourceRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2WebResourceRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2WebResourceRequestedEventHandlerVtbl
}

func (this *ICoreWebView2WebResourceRequestedEventHandlerComObj) ICoreWebView2WebResourceRequestedEventHandler() *ICoreWebView2WebResourceRequestedEventHandler{
	return (*ICoreWebView2WebResourceRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2WebResourceRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2WebResourceRequestedEventHandlerComObj(impl ICoreWebView2WebResourceRequestedEventHandlerInterface, scoped bool) *ICoreWebView2WebResourceRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2WebResourceRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2WebResourceRequestedEventHandler(impl ICoreWebView2WebResourceRequestedEventHandlerInterface) *ICoreWebView2WebResourceRequestedEventHandler {
	return NewICoreWebView2WebResourceRequestedEventHandlerComObj(impl, true).ICoreWebView2WebResourceRequestedEventHandler()}

//
type ICoreWebView2WebResourceRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2WebResourceRequestedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) com.Error
}
func (this *ICoreWebView2WebResourceRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2WebResourceRequestedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2WebResourceRequestedEventArgs) com.Error, scoped bool) *ICoreWebView2WebResourceRequestedEventHandler {
	impl := &ICoreWebView2WebResourceRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2WebResourceRequestedEventHandlerComObj(impl, scoped).ICoreWebView2WebResourceRequestedEventHandler()
}

