package wv2

import (
	"github.com/zzl/go-win32api/v2/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 3117DA26-AE13-438D-BD46-EDBEB2C4CE81
var IID_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler = syscall.GUID{0x3117DA26, 0xAE13, 0x438D,
	[8]byte{0xBD, 0x46, 0xED, 0xBE, 0xB2, 0xC4, 0xCE, 0x81}}

type ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

type ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInterface
}

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInterface)
}

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	var ret com.Error
	return ret
}

type ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj) impl() ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInterface)
}

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj) Invoke(sender *ICoreWebView2, args *win32.IUnknown) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
	}
	if _pICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl != nil {
		return _pICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl
	}
	_pICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl = &ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:       syscall.NewCallback((*ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerVtbl
}

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj) ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler() *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler {
	return (*ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj(impl ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInterface, scoped bool) *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler(impl ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerInterface) *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler {
	return NewICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj(impl, true).ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler()
}

//
type ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerByFuncImpl struct {
	ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerImpl
	handlerFunc func(sender *ICoreWebView2, args *win32.IUnknown) com.Error
}

func (this *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2, args *win32.IUnknown) com.Error {
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerByFunc(handlerFunc func(sender *ICoreWebView2, args *win32.IUnknown) com.Error, scoped bool) *ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler {
	impl := &ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandlerComObj(impl, scoped).ICoreWebView2IsDefaultDownloadDialogOpenChangedEventHandler()
}

