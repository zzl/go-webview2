package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// F9A2976E-D34E-44FC-ADEE-81B6B57CA914
var IID_ICoreWebView2NewBrowserVersionAvailableEventHandler = syscall.GUID{0xF9A2976E, 0xD34E, 0x44FC,
	[8]byte{0xAD, 0xEE, 0x81, 0xB6, 0xB5, 0x7C, 0xA9, 0x14}}

type ICoreWebView2NewBrowserVersionAvailableEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2NewBrowserVersionAvailableEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2Environment, args *win32.IUnknown) com.Error
}

type ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2NewBrowserVersionAvailableEventHandlerInterface
}

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2NewBrowserVersionAvailableEventHandlerInterface)
}

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2NewBrowserVersionAvailableEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl) Invoke(sender *ICoreWebView2Environment, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj) impl() ICoreWebView2NewBrowserVersionAvailableEventHandlerInterface {
	return this.Impl().(ICoreWebView2NewBrowserVersionAvailableEventHandlerInterface)
}

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj) Invoke(sender *ICoreWebView2Environment, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl *ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl != nil {
		return _pICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl
	}
	_pICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl = &ICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2NewBrowserVersionAvailableEventHandlerVtbl
}

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj) ICoreWebView2NewBrowserVersionAvailableEventHandler() *ICoreWebView2NewBrowserVersionAvailableEventHandler {
	return (*ICoreWebView2NewBrowserVersionAvailableEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2NewBrowserVersionAvailableEventHandlerComObj(impl ICoreWebView2NewBrowserVersionAvailableEventHandlerInterface, scoped bool) *ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2NewBrowserVersionAvailableEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2NewBrowserVersionAvailableEventHandler(impl ICoreWebView2NewBrowserVersionAvailableEventHandlerInterface) *ICoreWebView2NewBrowserVersionAvailableEventHandler {
	return NewICoreWebView2NewBrowserVersionAvailableEventHandlerComObj(impl, true).ICoreWebView2NewBrowserVersionAvailableEventHandler()
}

type ICoreWebView2NewBrowserVersionAvailableEventHandlerByFuncImpl struct {
	ICoreWebView2NewBrowserVersionAvailableEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2Environment, args *win32.IUnknown) com.Error
}

func (this *ICoreWebView2NewBrowserVersionAvailableEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2Environment, args *win32.IUnknown) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2NewBrowserVersionAvailableEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2Environment, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2NewBrowserVersionAvailableEventHandler {
	impl := &ICoreWebView2NewBrowserVersionAvailableEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2NewBrowserVersionAvailableEventHandlerComObj(impl, scoped).ICoreWebView2NewBrowserVersionAvailableEventHandler()
}
