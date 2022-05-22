package wv2

import (
	"github.com/zzl/go-win32api/win32"
	"github.com/zzl/go-com/com"
	"syscall"
	"unsafe"
)

// 828E8AB6-D94C-4264-9CEF-5217170D6251
var IID_ICoreWebView2BytesReceivedChangedEventHandler = syscall.GUID{0x828E8AB6, 0xD94C, 0x4264, 
	[8]byte{0x9C, 0xEF, 0x52, 0x17, 0x17, 0x0D, 0x62, 0x51}}

type ICoreWebView2BytesReceivedChangedEventHandler struct {
	win32.IUnknown
}

type ICoreWebView2BytesReceivedChangedEventHandlerInterface interface {
	win32.IUnknownInterface
	Invoke(sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error
}

type ICoreWebView2BytesReceivedChangedEventHandlerImpl struct {
	com.IUnknownImpl
	RealObject ICoreWebView2BytesReceivedChangedEventHandlerInterface
}

func (this *ICoreWebView2BytesReceivedChangedEventHandlerImpl) SetRealObject(obj interface{}) {
	this.RealObject = obj.(ICoreWebView2BytesReceivedChangedEventHandlerInterface)
}

func (this *ICoreWebView2BytesReceivedChangedEventHandlerImpl) QueryInterface(riid *syscall.GUID, ppvObject unsafe.Pointer) win32.HRESULT {
	if *riid == IID_ICoreWebView2BytesReceivedChangedEventHandler {
		this.AssignPpvObject(ppvObject)
		this.AddRef()
		return win32.S_OK
	}
	return this.IUnknownImpl.QueryInterface(riid, ppvObject)
}

func (this *ICoreWebView2BytesReceivedChangedEventHandlerImpl) Invoke(sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error {
	var ret com.Error
	return ret
}
type ICoreWebView2BytesReceivedChangedEventHandlerVtbl struct {
	win32.IUnknownVtbl
	Invoke uintptr
}

type ICoreWebView2BytesReceivedChangedEventHandlerComObj struct {
	com.IUnknownComObj
}

func (this *ICoreWebView2BytesReceivedChangedEventHandlerComObj) impl() ICoreWebView2BytesReceivedChangedEventHandlerInterface {
	return this.Impl().(ICoreWebView2BytesReceivedChangedEventHandlerInterface)
}

func (this *ICoreWebView2BytesReceivedChangedEventHandlerComObj) Invoke(sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) uintptr {
	return (uintptr)(this.impl().Invoke(sender, args))
}

var _pICoreWebView2BytesReceivedChangedEventHandlerVtbl *ICoreWebView2BytesReceivedChangedEventHandlerVtbl

func (this *ICoreWebView2BytesReceivedChangedEventHandlerComObj) BuildVtbl(lock bool) *ICoreWebView2BytesReceivedChangedEventHandlerVtbl {
	if lock {
		com.MuVtbl.Lock()
		defer com.MuVtbl.Unlock()
}
	if _pICoreWebView2BytesReceivedChangedEventHandlerVtbl != nil {
		return _pICoreWebView2BytesReceivedChangedEventHandlerVtbl
	}
	_pICoreWebView2BytesReceivedChangedEventHandlerVtbl = &ICoreWebView2BytesReceivedChangedEventHandlerVtbl{
		IUnknownVtbl: *this.IUnknownComObj.BuildVtbl(false),
		Invoke:	syscall.NewCallback((*ICoreWebView2BytesReceivedChangedEventHandlerComObj).Invoke),
	}
	return _pICoreWebView2BytesReceivedChangedEventHandlerVtbl
}

func (this *ICoreWebView2BytesReceivedChangedEventHandlerComObj) ICoreWebView2BytesReceivedChangedEventHandler() *ICoreWebView2BytesReceivedChangedEventHandler{
	return (*ICoreWebView2BytesReceivedChangedEventHandler)(unsafe.Pointer(this))
}

func (this *ICoreWebView2BytesReceivedChangedEventHandlerComObj) GetVtbl() *win32.IUnknownVtbl {
	return &this.BuildVtbl(true).IUnknownVtbl
}

func NewICoreWebView2BytesReceivedChangedEventHandlerComObj(impl ICoreWebView2BytesReceivedChangedEventHandlerInterface, scoped bool) *ICoreWebView2BytesReceivedChangedEventHandlerComObj {
	comObj := com.NewComObj[ICoreWebView2BytesReceivedChangedEventHandlerComObj](impl)
	if scoped {
		com.AddToScope(comObj)
	}
	return comObj
}

func NewICoreWebView2BytesReceivedChangedEventHandler(impl ICoreWebView2BytesReceivedChangedEventHandlerInterface) *ICoreWebView2BytesReceivedChangedEventHandler {
	return NewICoreWebView2BytesReceivedChangedEventHandlerComObj(impl, true).ICoreWebView2BytesReceivedChangedEventHandler()}

//
type ICoreWebView2BytesReceivedChangedEventHandlerByFuncImpl struct {
	ICoreWebView2BytesReceivedChangedEventHandlerImpl
	handlerFunc func (sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error
}
func (this *ICoreWebView2BytesReceivedChangedEventHandlerByFuncImpl) Invoke(sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error{
	return this.handlerFunc(sender, args)
}

func NewICoreWebView2BytesReceivedChangedEventHandlerByFunc(handlerFunc func (sender *ICoreWebView2DownloadOperation, args *com.UnknownClass) com.Error, scoped bool) *ICoreWebView2BytesReceivedChangedEventHandler {
	impl := &ICoreWebView2BytesReceivedChangedEventHandlerByFuncImpl{handlerFunc: handlerFunc}
	return NewICoreWebView2BytesReceivedChangedEventHandlerComObj(impl, scoped).ICoreWebView2BytesReceivedChangedEventHandler()
}

