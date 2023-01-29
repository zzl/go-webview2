package wv2

import (
	"github.com/zzl/go-com/com"
	"github.com/zzl/go-win32api/v2/win32"
	"syscall"
	"unsafe"
)

// 57D90347-CD0E-4952-A4A2-7483A2756F08
var IID_ICoreWebView2IsMutedChangedEventHandler = syscall.GUID{0x57D90347, 0xCD0E, 0x4952,
	[8]byte{0xA4, 0xA2, 0x74, 0x83, 0xA2, 0x75, 0x6F, 0x08}}

type ICoreWebView2IsMutedChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2IsMutedChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

type ICoreWebView2IsMutedChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2IsMutedChangedEventHandlerInterface
}

func (this *ICoreWebView2IsMutedChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2IsMutedChangedEventHandlerInterface)
}

func (this *ICoreWebView2IsMutedChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2IsMutedChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2IsMutedChangedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2IsMutedChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2IsMutedChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2IsMutedChangedEventHandlerComObj) impl() ICoreWebView2IsMutedChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2IsMutedChangedEventHandlerInterface)
}

func (this *ICoreWebView2IsMutedChangedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2IsMutedChangedEventHandlerVtbl *ICoreWebView2IsMutedChangedEventHandlerVtbl

func (this *ICoreWebView2IsMutedChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2IsMutedChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2IsMutedChangedEventHandlerVtbl != nil {
		return _pICoreWebView2IsMutedChangedEventHandlerVtbl
	}
	_pICoreWebView2IsMutedChangedEventHandlerVtbl = &ICoreWebView2IsMutedChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2IsMutedChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2IsMutedChangedEventHandlerVtbl
}

func (this *ICoreWebView2IsMutedChangedEventHandlerComObj) ICoreWebView2IsMutedChangedEventHandler() *ICoreWebView2IsMutedChangedEventHandler {
	return (*ICoreWebView2IsMutedChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2IsMutedChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2IsMutedChangedEventHandlerComObj(impl ICoreWebView2IsMutedChangedEventHandlerInterface, scoped bool) *ICoreWebView2IsMutedChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2IsMutedChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2IsMutedChangedEventHandler(impl ICoreWebView2IsMutedChangedEventHandlerInterface) *ICoreWebView2IsMutedChangedEventHandler {
	return NewICoreWebView2IsMutedChangedEventHandlerComObj(impl, true).ICoreWebView2IsMutedChangedEventHandler()
}

type ICoreWebView2IsMutedChangedEventHandlerByFuncImpl struct {
	ICoreWebView2IsMutedChangedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

func (this *ICoreWebView2IsMutedChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2IsMutedChangedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2IsMutedChangedEventHandler {
	impl := &ICoreWebView2IsMutedChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2IsMutedChangedEventHandlerComObj(impl, scoped).ICoreWebView2IsMutedChangedEventHandler()
}
