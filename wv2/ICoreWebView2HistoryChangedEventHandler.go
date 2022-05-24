package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// C79A420C-EFD9-4058-9295-3E8B4BCAB645
var IID_ICoreWebView2HistoryChangedEventHandler = syscall.GUID{0xC79A420C, 0xEFD9, 0x4058, 
	[8]byte{0x92, 0x95, 0x3E, 0x8B, 0x4B, 0xCA, 0xB6, 0x45}}

type ICoreWebView2HistoryChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2HistoryChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

type ICoreWebView2HistoryChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2HistoryChangedEventHandlerInterface
}

func (this *ICoreWebView2HistoryChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2HistoryChangedEventHandlerInterface)
}

func (this *ICoreWebView2HistoryChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2HistoryChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2HistoryChangedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2HistoryChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2HistoryChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2HistoryChangedEventHandlerComObj) impl() ICoreWebView2HistoryChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2HistoryChangedEventHandlerInterface)
}

func (this *ICoreWebView2HistoryChangedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2HistoryChangedEventHandlerVtbl *ICoreWebView2HistoryChangedEventHandlerVtbl

func (this *ICoreWebView2HistoryChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2HistoryChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2HistoryChangedEventHandlerVtbl != nil {
		return _pICoreWebView2HistoryChangedEventHandlerVtbl
	}
	_pICoreWebView2HistoryChangedEventHandlerVtbl = &ICoreWebView2HistoryChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2HistoryChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2HistoryChangedEventHandlerVtbl
}

func (this *ICoreWebView2HistoryChangedEventHandlerComObj) ICoreWebView2HistoryChangedEventHandler() *ICoreWebView2HistoryChangedEventHandler{
	return (*ICoreWebView2HistoryChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2HistoryChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2HistoryChangedEventHandlerComObj(impl ICoreWebView2HistoryChangedEventHandlerInterface, scoped bool) *ICoreWebView2HistoryChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2HistoryChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2HistoryChangedEventHandler(impl ICoreWebView2HistoryChangedEventHandlerInterface) *ICoreWebView2HistoryChangedEventHandler {
	return NewICoreWebView2HistoryChangedEventHandlerComObj(impl, true).ICoreWebView2HistoryChangedEventHandler()}

//
type ICoreWebView2HistoryChangedEventHandlerByFuncImpl struct {
	ICoreWebView2HistoryChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error
}
func (this *ICoreWebView2HistoryChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2HistoryChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2HistoryChangedEventHandler {
	impl := &ICoreWebView2HistoryChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2HistoryChangedEventHandlerComObj(impl, scoped).ICoreWebView2HistoryChangedEventHandler()
}

