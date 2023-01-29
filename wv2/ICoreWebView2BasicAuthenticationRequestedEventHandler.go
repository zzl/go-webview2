package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 58B4D6C2-18D4-497E-B39B-9A96533FA278
var IID_ICoreWebView2BasicAuthenticationRequestedEventHandler = syscall.GUID{0x58B4D6C2, 0x18D4, 0x497E,
	[8]byte{0xB3, 0x9B, 0x9A, 0x96, 0x53, 0x3F, 0xA2, 0x78}}

type ICoreWebView2BasicAuthenticationRequestedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2BasicAuthenticationRequestedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) com.Error
}

type ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2BasicAuthenticationRequestedEventHandlerInterface
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2BasicAuthenticationRequestedEventHandlerInterface)
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2BasicAuthenticationRequestedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj) impl() ICoreWebView2BasicAuthenticationRequestedEventHandlerInterface {
	return this.Impl().(ICoreWebView2BasicAuthenticationRequestedEventHandlerInterface)
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl *ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl != nil {
		return _pICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl
	}
	_pICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl = &ICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2BasicAuthenticationRequestedEventHandlerVtbl
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj) ICoreWebView2BasicAuthenticationRequestedEventHandler() *ICoreWebView2BasicAuthenticationRequestedEventHandler {
	return (*ICoreWebView2BasicAuthenticationRequestedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2BasicAuthenticationRequestedEventHandlerComObj(impl ICoreWebView2BasicAuthenticationRequestedEventHandlerInterface, scoped bool) *ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2BasicAuthenticationRequestedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2BasicAuthenticationRequestedEventHandler(impl ICoreWebView2BasicAuthenticationRequestedEventHandlerInterface) *ICoreWebView2BasicAuthenticationRequestedEventHandler {
	return NewICoreWebView2BasicAuthenticationRequestedEventHandlerComObj(impl, true).ICoreWebView2BasicAuthenticationRequestedEventHandler()
}

type ICoreWebView2BasicAuthenticationRequestedEventHandlerByFuncImpl struct {
	ICoreWebView2BasicAuthenticationRequestedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) com.Error
}

func (this *ICoreWebView2BasicAuthenticationRequestedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2BasicAuthenticationRequestedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2BasicAuthenticationRequestedEventArgs) com.Error, scoped bool) *ICoreWebView2BasicAuthenticationRequestedEventHandler {
	impl := &ICoreWebView2BasicAuthenticationRequestedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2BasicAuthenticationRequestedEventHandlerComObj(impl, scoped).ICoreWebView2BasicAuthenticationRequestedEventHandler()
}
