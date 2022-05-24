package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// E45D98B1-AFEF-45BE-8BAF-6C7728867F73
var IID_ICoreWebView2ContainsFullScreenElementChangedEventHandler = syscall.GUID{0xE45D98B1, 0xAFEF, 0x45BE, 
	[8]byte{0x8B, 0xAF, 0x6C, 0x77, 0x28, 0x86, 0x7F, 0x73}}

type ICoreWebView2ContainsFullScreenElementChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2ContainsFullScreenElementChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

type ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2ContainsFullScreenElementChangedEventHandlerInterface
}

func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2ContainsFullScreenElementChangedEventHandlerInterface)
}

func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2ContainsFullScreenElementChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj) impl() ICoreWebView2ContainsFullScreenElementChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2ContainsFullScreenElementChangedEventHandlerInterface)
}

func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl *ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl

func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl != nil {
		return _pICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl
	}
	_pICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl = &ICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2ContainsFullScreenElementChangedEventHandlerVtbl
}

func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj) ICoreWebView2ContainsFullScreenElementChangedEventHandler() *ICoreWebView2ContainsFullScreenElementChangedEventHandler{
	return (*ICoreWebView2ContainsFullScreenElementChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj(impl ICoreWebView2ContainsFullScreenElementChangedEventHandlerInterface, scoped bool) *ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2ContainsFullScreenElementChangedEventHandler(impl ICoreWebView2ContainsFullScreenElementChangedEventHandlerInterface) *ICoreWebView2ContainsFullScreenElementChangedEventHandler {
	return NewICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj(impl, true).ICoreWebView2ContainsFullScreenElementChangedEventHandler()}

//
type ICoreWebView2ContainsFullScreenElementChangedEventHandlerByFuncImpl struct {
	ICoreWebView2ContainsFullScreenElementChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error
}
func (this *ICoreWebView2ContainsFullScreenElementChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2ContainsFullScreenElementChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2ContainsFullScreenElementChangedEventHandler {
	impl := &ICoreWebView2ContainsFullScreenElementChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2ContainsFullScreenElementChangedEventHandlerComObj(impl, scoped).ICoreWebView2ContainsFullScreenElementChangedEventHandler()
}

