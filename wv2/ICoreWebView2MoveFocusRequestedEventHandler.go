package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 69035451-6DC7-4CB8-9BCE-B2BD70AD289F
var IID_ICoreWebView2MoveFocusRequestedEventHandler = syscall.GUID{0x69035451, 0x6DC7, 0x4CB8, 
	[8]byte{0x9B, 0xCE, 0xB2, 0xBD, 0x70, 0xAD, 0x28, 0x9F}}

type ICoreWebView2MoveFocusRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2MoveFocusRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) com.Error
}

type ICoreWebView2MoveFocusRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2MoveFocusRequestedEventHandlerInterface
}

func (this *ICoreWebView2MoveFocusRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2MoveFocusRequestedEventHandlerInterface)
}

func (this *ICoreWebView2MoveFocusRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2MoveFocusRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2MoveFocusRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2MoveFocusRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2MoveFocusRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2MoveFocusRequestedEventHandlerComObj) impl() ICoreWebView2MoveFocusRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2MoveFocusRequestedEventHandlerInterface)
}

func (this *ICoreWebView2MoveFocusRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2MoveFocusRequestedEventHandlerVtbl *ICoreWebView2MoveFocusRequestedEventHandlerVtbl

func (this *ICoreWebView2MoveFocusRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2MoveFocusRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2MoveFocusRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2MoveFocusRequestedEventHandlerVtbl
	}
	_pICoreWebView2MoveFocusRequestedEventHandlerVtbl = &ICoreWebView2MoveFocusRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2MoveFocusRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2MoveFocusRequestedEventHandlerVtbl
}

func (this *ICoreWebView2MoveFocusRequestedEventHandlerComObj) ICoreWebView2MoveFocusRequestedEventHandler() *ICoreWebView2MoveFocusRequestedEventHandler{
	return (*ICoreWebView2MoveFocusRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2MoveFocusRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2MoveFocusRequestedEventHandlerComObj(impl ICoreWebView2MoveFocusRequestedEventHandlerInterface, scoped bool) *ICoreWebView2MoveFocusRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2MoveFocusRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2MoveFocusRequestedEventHandler(impl ICoreWebView2MoveFocusRequestedEventHandlerInterface) *ICoreWebView2MoveFocusRequestedEventHandler {
	return NewICoreWebView2MoveFocusRequestedEventHandlerComObj(impl, true).ICoreWebView2MoveFocusRequestedEventHandler()}

//
type ICoreWebView2MoveFocusRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2MoveFocusRequestedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) com.Error
}
func (this *ICoreWebView2MoveFocusRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2MoveFocusRequestedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2Controller, args *ICoreWebView2MoveFocusRequestedEventArgs) com.Error, scoped bool) *ICoreWebView2MoveFocusRequestedEventHandler {
	impl := &ICoreWebView2MoveFocusRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2MoveFocusRequestedEventHandlerComObj(impl, scoped).ICoreWebView2MoveFocusRequestedEventHandler()
}

