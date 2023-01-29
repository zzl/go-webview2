package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// B29C7E28-FA79-41A8-8E44-65811C76DCB2
var IID_ICoreWebView2AcceleratorKeyPressedEventHandler = syscall.GUID{0xB29C7E28, 0xFA79, 0x41A8,
	[8]byte{0x8E, 0x44, 0x65, 0x81, 0x1C, 0x76, 0xDC, 0xB2}}

type ICoreWebView2AcceleratorKeyPressedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2AcceleratorKeyPressedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) com.Error
}

type ICoreWebView2AcceleratorKeyPressedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2AcceleratorKeyPressedEventHandlerInterface
}

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2AcceleratorKeyPressedEventHandlerInterface)
}

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2AcceleratorKeyPressedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerImpl) Invoke(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2AcceleratorKeyPressedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2AcceleratorKeyPressedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerComObj) impl() ICoreWebView2AcceleratorKeyPressedEventHandlerInterface {
	return this.Impl().(ICoreWebView2AcceleratorKeyPressedEventHandlerInterface)
}

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerComObj) Invoke(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2AcceleratorKeyPressedEventHandlerVtbl *ICoreWebView2AcceleratorKeyPressedEventHandlerVtbl

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2AcceleratorKeyPressedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2AcceleratorKeyPressedEventHandlerVtbl != nil {
		return _pICoreWebView2AcceleratorKeyPressedEventHandlerVtbl
	}
	_pICoreWebView2AcceleratorKeyPressedEventHandlerVtbl = &ICoreWebView2AcceleratorKeyPressedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2AcceleratorKeyPressedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2AcceleratorKeyPressedEventHandlerVtbl
}

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerComObj) ICoreWebView2AcceleratorKeyPressedEventHandler() *ICoreWebView2AcceleratorKeyPressedEventHandler {
	return (*ICoreWebView2AcceleratorKeyPressedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2AcceleratorKeyPressedEventHandlerComObj(impl ICoreWebView2AcceleratorKeyPressedEventHandlerInterface, scoped bool) *ICoreWebView2AcceleratorKeyPressedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2AcceleratorKeyPressedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2AcceleratorKeyPressedEventHandler(impl ICoreWebView2AcceleratorKeyPressedEventHandlerInterface) *ICoreWebView2AcceleratorKeyPressedEventHandler {
	return NewICoreWebView2AcceleratorKeyPressedEventHandlerComObj(impl, true).ICoreWebView2AcceleratorKeyPressedEventHandler()
}

//
type ICoreWebView2AcceleratorKeyPressedEventHandlerByFuncImpl struct {
	ICoreWebView2AcceleratorKeyPressedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) com.Error
}

func (this *ICoreWebView2AcceleratorKeyPressedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2AcceleratorKeyPressedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2Controller, args *ICoreWebView2AcceleratorKeyPressedEventArgs) com.Error, scoped bool) *ICoreWebView2AcceleratorKeyPressedEventHandler {
	impl := &ICoreWebView2AcceleratorKeyPressedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2AcceleratorKeyPressedEventHandlerComObj(impl, scoped).ICoreWebView2AcceleratorKeyPressedEventHandler()
}

