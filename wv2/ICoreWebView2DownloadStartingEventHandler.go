package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// EFEDC989-C396-41CA-83F7-07F845A55724
var IID_ICoreWebView2DownloadStartingEventHandler = syscall.GUID{0xEFEDC989, 0xC396, 0x41CA, 
	[8]byte{0x83, 0xF7, 0x07, 0xF8, 0x45, 0xA5, 0x57, 0x24}}

type ICoreWebView2DownloadStartingEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2DownloadStartingEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) com.Error
}

type ICoreWebView2DownloadStartingEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2DownloadStartingEventHandlerInterface
}

func (this *ICoreWebView2DownloadStartingEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2DownloadStartingEventHandlerInterface)
}

func (this *ICoreWebView2DownloadStartingEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2DownloadStartingEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2DownloadStartingEventHandlerImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2DownloadStartingEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2DownloadStartingEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2DownloadStartingEventHandlerComObj) impl() ICoreWebView2DownloadStartingEventHandlerInterface {
	return this.Impl().(ICoreWebView2DownloadStartingEventHandlerInterface)
}

func (this *ICoreWebView2DownloadStartingEventHandlerComObj) Invoke(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2DownloadStartingEventHandlerVtbl *ICoreWebView2DownloadStartingEventHandlerVtbl

func (this *ICoreWebView2DownloadStartingEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2DownloadStartingEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2DownloadStartingEventHandlerVtbl != nil {
		return _pICoreWebView2DownloadStartingEventHandlerVtbl
	}
	_pICoreWebView2DownloadStartingEventHandlerVtbl = &ICoreWebView2DownloadStartingEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2DownloadStartingEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2DownloadStartingEventHandlerVtbl
}

func (this *ICoreWebView2DownloadStartingEventHandlerComObj) ICoreWebView2DownloadStartingEventHandler() *ICoreWebView2DownloadStartingEventHandler{
	return (*ICoreWebView2DownloadStartingEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2DownloadStartingEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2DownloadStartingEventHandlerComObj(impl ICoreWebView2DownloadStartingEventHandlerInterface, scoped bool) *ICoreWebView2DownloadStartingEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2DownloadStartingEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2DownloadStartingEventHandler(impl ICoreWebView2DownloadStartingEventHandlerInterface) *ICoreWebView2DownloadStartingEventHandler {
	return NewICoreWebView2DownloadStartingEventHandlerComObj(impl, true).ICoreWebView2DownloadStartingEventHandler()}

//
type ICoreWebView2DownloadStartingEventHandlerByFuncImpl struct {
	ICoreWebView2DownloadStartingEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) com.Error
}
func (this *ICoreWebView2DownloadStartingEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2DownloadStartingEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2, args *ICoreWebView2DownloadStartingEventArgs) com.Error, scoped bool) *ICoreWebView2DownloadStartingEventHandler {
	impl := &ICoreWebView2DownloadStartingEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2DownloadStartingEventHandlerComObj(impl, scoped).ICoreWebView2DownloadStartingEventHandler()
}

