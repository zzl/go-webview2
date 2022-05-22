package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// D33A35BF-1C49-4F98-93AB-006E0533FE1C
var IID_ICoreWebView2NavigationCompletedEventHandler = syscall.GUID{0xD33A35BF, 0x1C49, 0x4F98, 
	[8]byte{0x93, 0xAB, 0x00, 0x6E, 0x05, 0x33, 0xFE, 0x1C}}

type ICoreWebView2NavigationCompletedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2NavigationCompletedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) com.Error
}

type ICoreWebView2NavigationCompletedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2NavigationCompletedEventHandlerInterface
}

func (this *ICoreWebView2NavigationCompletedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2NavigationCompletedEventHandlerInterface)
}

func (this *ICoreWebView2NavigationCompletedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2NavigationCompletedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2NavigationCompletedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2NavigationCompletedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2NavigationCompletedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2NavigationCompletedEventHandlerComObj) impl() ICoreWebView2NavigationCompletedEventHandlerInterface {
	return this.Impl().(ICoreWebView2NavigationCompletedEventHandlerInterface)
}

func (this *ICoreWebView2NavigationCompletedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2NavigationCompletedEventHandlerVtbl *ICoreWebView2NavigationCompletedEventHandlerVtbl

func (this *ICoreWebView2NavigationCompletedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2NavigationCompletedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2NavigationCompletedEventHandlerVtbl != nil {
		return _pICoreWebView2NavigationCompletedEventHandlerVtbl
	}
	_pICoreWebView2NavigationCompletedEventHandlerVtbl = &ICoreWebView2NavigationCompletedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2NavigationCompletedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2NavigationCompletedEventHandlerVtbl
}

func (this *ICoreWebView2NavigationCompletedEventHandlerComObj) ICoreWebView2NavigationCompletedEventHandler() *ICoreWebView2NavigationCompletedEventHandler{
	return (*ICoreWebView2NavigationCompletedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2NavigationCompletedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2NavigationCompletedEventHandlerComObj(impl ICoreWebView2NavigationCompletedEventHandlerInterface, scoped bool) *ICoreWebView2NavigationCompletedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2NavigationCompletedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2NavigationCompletedEventHandler(impl ICoreWebView2NavigationCompletedEventHandlerInterface) *ICoreWebView2NavigationCompletedEventHandler {
	return NewICoreWebView2NavigationCompletedEventHandlerComObj(impl, true).ICoreWebView2NavigationCompletedEventHandler()}

//
type ICoreWebView2NavigationCompletedEventHandlerByFuncImpl struct {
	ICoreWebView2NavigationCompletedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) com.Error
}
func (this *ICoreWebView2NavigationCompletedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2NavigationCompletedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2NavigationCompletedEventArgs) com.Error, scoped bool) *ICoreWebView2NavigationCompletedEventHandler {
	impl := &ICoreWebView2NavigationCompletedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2NavigationCompletedEventHandlerComObj(impl, scoped).ICoreWebView2NavigationCompletedEventHandler()
}

