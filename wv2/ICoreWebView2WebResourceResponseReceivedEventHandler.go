package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 7DE9898A-24F5-40C3-A2DE-D4F458E69828
var IID_ICoreWebView2WebResourceResponseReceivedEventHandler = syscall.GUID{0x7DE9898A, 0x24F5, 0x40C3,
	[8]byte{0xA2, 0xDE, 0xD4, 0xF4, 0x58, 0xE6, 0x98, 0x28}}

type ICoreWebView2WebResourceResponseReceivedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2WebResourceResponseReceivedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) com.Error
}

type ICoreWebView2WebResourceResponseReceivedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2WebResourceResponseReceivedEventHandlerInterface
}

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2WebResourceResponseReceivedEventHandlerInterface)
}

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2WebResourceResponseReceivedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2WebResourceResponseReceivedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerComObj) impl() ICoreWebView2WebResourceResponseReceivedEventHandlerInterface {
	return this.Impl().(ICoreWebView2WebResourceResponseReceivedEventHandlerInterface)
}

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2WebResourceResponseReceivedEventHandlerVtbl *ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2WebResourceResponseReceivedEventHandlerVtbl != nil {
		return _pICoreWebView2WebResourceResponseReceivedEventHandlerVtbl
	}
	_pICoreWebView2WebResourceResponseReceivedEventHandlerVtbl = &ICoreWebView2WebResourceResponseReceivedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2WebResourceResponseReceivedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2WebResourceResponseReceivedEventHandlerVtbl
}

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerComObj) ICoreWebView2WebResourceResponseReceivedEventHandler() *ICoreWebView2WebResourceResponseReceivedEventHandler {
	return (*ICoreWebView2WebResourceResponseReceivedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2WebResourceResponseReceivedEventHandlerComObj(impl ICoreWebView2WebResourceResponseReceivedEventHandlerInterface, scoped bool) *ICoreWebView2WebResourceResponseReceivedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2WebResourceResponseReceivedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2WebResourceResponseReceivedEventHandler(impl ICoreWebView2WebResourceResponseReceivedEventHandlerInterface) *ICoreWebView2WebResourceResponseReceivedEventHandler {
	return NewICoreWebView2WebResourceResponseReceivedEventHandlerComObj(impl, true).ICoreWebView2WebResourceResponseReceivedEventHandler()
}

//
type ICoreWebView2WebResourceResponseReceivedEventHandlerByFuncImpl struct {
	ICoreWebView2WebResourceResponseReceivedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) com.Error
}

func (this *ICoreWebView2WebResourceResponseReceivedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2WebResourceResponseReceivedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *ICoreWebView2WebResourceResponseReceivedEventArgs) com.Error, scoped bool) *ICoreWebView2WebResourceResponseReceivedEventHandler {
	impl := &ICoreWebView2WebResourceResponseReceivedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2WebResourceResponseReceivedEventHandlerComObj(impl, scoped).ICoreWebView2WebResourceResponseReceivedEventHandler()
}

