package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 28F0D425-93FE-4E63-9F8D-2AEEC6D3BA1E
var IID_ICoreWebView2EstimatedEndTimeChangedEventHandler = syscall.GUID{0x28F0D425, 0x93FE, 0x4E63, 
	[8]byte{0x9F, 0x8D, 0x2A, 0xEE, 0xC6, 0xD3, 0xBA, 0x1E}}

type ICoreWebView2EstimatedEndTimeChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2EstimatedEndTimeChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error
}

type ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2EstimatedEndTimeChangedEventHandlerInterface
}

func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2EstimatedEndTimeChangedEventHandlerInterface)
}

func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2EstimatedEndTimeChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl) Invoke(sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj) impl() ICoreWebView2EstimatedEndTimeChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2EstimatedEndTimeChangedEventHandlerInterface)
}

func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj) Invoke(sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl *ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl

func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl != nil {
		return _pICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl
	}
	_pICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl = &ICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2EstimatedEndTimeChangedEventHandlerVtbl
}

func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj) ICoreWebView2EstimatedEndTimeChangedEventHandler() *ICoreWebView2EstimatedEndTimeChangedEventHandler{
	return (*ICoreWebView2EstimatedEndTimeChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2EstimatedEndTimeChangedEventHandlerComObj(impl ICoreWebView2EstimatedEndTimeChangedEventHandlerInterface, scoped bool) *ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2EstimatedEndTimeChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2EstimatedEndTimeChangedEventHandler(impl ICoreWebView2EstimatedEndTimeChangedEventHandlerInterface) *ICoreWebView2EstimatedEndTimeChangedEventHandler {
	return NewICoreWebView2EstimatedEndTimeChangedEventHandlerComObj(impl, true).ICoreWebView2EstimatedEndTimeChangedEventHandler()}

//
type ICoreWebView2EstimatedEndTimeChangedEventHandlerByFuncImpl struct {
	ICoreWebView2EstimatedEndTimeChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error
}
func (this *ICoreWebView2EstimatedEndTimeChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2EstimatedEndTimeChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error, scoped bool) *ICoreWebView2EstimatedEndTimeChangedEventHandler {
	impl := &ICoreWebView2EstimatedEndTimeChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2EstimatedEndTimeChangedEventHandlerComObj(impl, scoped).ICoreWebView2EstimatedEndTimeChangedEventHandler()
}

