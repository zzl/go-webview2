package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 81336594-7EDE-4BA9-BF71-ACF0A95B58DD
var IID_ICoreWebView2StateChangedEventHandler = syscall.GUID{0x81336594, 0x7EDE, 0x4BA9, 
	[8]byte{0xBF, 0x71, 0xAC, 0xF0, 0xA9, 0x5B, 0x58, 0xDD}}

type ICoreWebView2StateChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2StateChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2DownloadOperation, args *win32.IUnknown) com.Error
}

type ICoreWebView2StateChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2StateChangedEventHandlerInterface
}

func (this *ICoreWebView2StateChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2StateChangedEventHandlerInterface)
}

func (this *ICoreWebView2StateChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2StateChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2StateChangedEventHandlerImpl) Invoke(sender *ICoreWebView2DownloadOperation, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2StateChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2StateChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2StateChangedEventHandlerComObj) impl() ICoreWebView2StateChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2StateChangedEventHandlerInterface)
}

func (this *ICoreWebView2StateChangedEventHandlerComObj) Invoke(sender *ICoreWebView2DownloadOperation, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2StateChangedEventHandlerVtbl *ICoreWebView2StateChangedEventHandlerVtbl

func (this *ICoreWebView2StateChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2StateChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2StateChangedEventHandlerVtbl != nil {
		return _pICoreWebView2StateChangedEventHandlerVtbl
	}
	_pICoreWebView2StateChangedEventHandlerVtbl = &ICoreWebView2StateChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2StateChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2StateChangedEventHandlerVtbl
}

func (this *ICoreWebView2StateChangedEventHandlerComObj) ICoreWebView2StateChangedEventHandler() *ICoreWebView2StateChangedEventHandler{
	return (*ICoreWebView2StateChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2StateChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2StateChangedEventHandlerComObj(impl ICoreWebView2StateChangedEventHandlerInterface, scoped bool) *ICoreWebView2StateChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2StateChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2StateChangedEventHandler(impl ICoreWebView2StateChangedEventHandlerInterface) *ICoreWebView2StateChangedEventHandler {
	return NewICoreWebView2StateChangedEventHandlerComObj(impl, true).ICoreWebView2StateChangedEventHandler()}

//
type ICoreWebView2StateChangedEventHandlerByFuncImpl struct {
	ICoreWebView2StateChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2DownloadOperation, args *win32.IUnknown) com.Error
}
func (this *ICoreWebView2StateChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2DownloadOperation, args *win32.IUnknown) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2StateChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2DownloadOperation, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2StateChangedEventHandler {
	impl := &ICoreWebView2StateChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2StateChangedEventHandlerComObj(impl, scoped).ICoreWebView2StateChangedEventHandler()
}

