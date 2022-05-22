package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// F5F2B923-953E-4042-9F95-F3A118E1AFD4
var IID_ICoreWebView2DocumentTitleChangedEventHandler = syscall.GUID{0xF5F2B923, 0x953E, 0x4042, 
	[8]byte{0x9F, 0x95, 0xF3, 0xA1, 0x18, 0xE1, 0xAF, 0xD4}}

type ICoreWebView2DocumentTitleChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2DocumentTitleChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *com.UnknownClass) com.Error
}

type ICoreWebView2DocumentTitleChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2DocumentTitleChangedEventHandlerInterface
}

func (this *ICoreWebView2DocumentTitleChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2DocumentTitleChangedEventHandlerInterface)
}

func (this *ICoreWebView2DocumentTitleChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2DocumentTitleChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2DocumentTitleChangedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *com.UnknownClass) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2DocumentTitleChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2DocumentTitleChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2DocumentTitleChangedEventHandlerComObj) impl() ICoreWebView2DocumentTitleChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2DocumentTitleChangedEventHandlerInterface)
}

func (this *ICoreWebView2DocumentTitleChangedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *com.UnknownClass) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2DocumentTitleChangedEventHandlerVtbl *ICoreWebView2DocumentTitleChangedEventHandlerVtbl

func (this *ICoreWebView2DocumentTitleChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2DocumentTitleChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2DocumentTitleChangedEventHandlerVtbl != nil {
		return _pICoreWebView2DocumentTitleChangedEventHandlerVtbl
	}
	_pICoreWebView2DocumentTitleChangedEventHandlerVtbl = &ICoreWebView2DocumentTitleChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2DocumentTitleChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2DocumentTitleChangedEventHandlerVtbl
}

func (this *ICoreWebView2DocumentTitleChangedEventHandlerComObj) ICoreWebView2DocumentTitleChangedEventHandler() *ICoreWebView2DocumentTitleChangedEventHandler{
	return (*ICoreWebView2DocumentTitleChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2DocumentTitleChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2DocumentTitleChangedEventHandlerComObj(impl ICoreWebView2DocumentTitleChangedEventHandlerInterface, scoped bool) *ICoreWebView2DocumentTitleChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2DocumentTitleChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2DocumentTitleChangedEventHandler(impl ICoreWebView2DocumentTitleChangedEventHandlerInterface) *ICoreWebView2DocumentTitleChangedEventHandler {
	return NewICoreWebView2DocumentTitleChangedEventHandlerComObj(impl, true).ICoreWebView2DocumentTitleChangedEventHandler()}

//
type ICoreWebView2DocumentTitleChangedEventHandlerByFuncImpl struct {
	ICoreWebView2DocumentTitleChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *com.UnknownClass) com.Error
}
func (this *ICoreWebView2DocumentTitleChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *com.UnknownClass) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2DocumentTitleChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *com.UnknownClass) com.Error, scoped bool) *ICoreWebView2DocumentTitleChangedEventHandler {
	impl := &ICoreWebView2DocumentTitleChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2DocumentTitleChangedEventHandlerComObj(impl, scoped).ICoreWebView2DocumentTitleChangedEventHandler()
}

